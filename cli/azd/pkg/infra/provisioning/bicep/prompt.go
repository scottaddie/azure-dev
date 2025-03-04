// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package bicep

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/azure/azure-dev/cli/azd/pkg/account"
	"github.com/azure/azure-dev/cli/azd/pkg/azure"
	"github.com/azure/azure-dev/cli/azd/pkg/input"
	"github.com/azure/azure-dev/cli/azd/pkg/output"
	"github.com/azure/azure-dev/cli/azd/pkg/password"
	"github.com/azure/azure-dev/cli/azd/pkg/prompt"

	"github.com/azure/azure-dev/cli/azd/pkg/infra/provisioning"
)

// promptDialogItemForParameter builds the input.PromptDialogItem for the given required parameter.
func (p *BicepProvider) promptDialogItemForParameter(
	key string,
	param azure.ArmTemplateParameterDefinition,
) input.PromptDialogItem {
	help, _ := param.Description()
	paramType := p.mapBicepTypeToInterfaceType(param.Type)

	var dialogItem input.PromptDialogItem
	dialogItem.ID = key
	dialogItem.DisplayName = key
	dialogItem.Required = true

	if help != "" {
		dialogItem.Description = to.Ptr(help)
	}

	if paramType == provisioning.ParameterTypeBoolean {
		dialogItem.Kind = "select"
		dialogItem.Choices = []input.PromptDialogChoice{{Value: "true"}, {Value: "false"}}
	} else if param.AllowedValues != nil {
		dialogItem.Kind = "select"
		for _, v := range *param.AllowedValues {
			dialogItem.Choices = append(dialogItem.Choices, input.PromptDialogChoice{Value: fmt.Sprintf("%v", v)})
		}
	} else if param.Secure() {
		dialogItem.Kind = "password"
	} else {
		dialogItem.Kind = "string"
	}

	return dialogItem
}

func autoGenerate(parameter string, azdMetadata azure.AzdMetadata) (string, error) {
	if azdMetadata.AutoGenerateConfig == nil {
		return "", fmt.Errorf("auto generation metadata config is missing for parameter '%s'", parameter)
	}
	genValue, err := password.Generate(password.GenerateConfig{
		Length:     azdMetadata.AutoGenerateConfig.Length,
		NoLower:    azdMetadata.AutoGenerateConfig.NoLower,
		NoUpper:    azdMetadata.AutoGenerateConfig.NoUpper,
		NoNumeric:  azdMetadata.AutoGenerateConfig.NoNumeric,
		NoSpecial:  azdMetadata.AutoGenerateConfig.NoSpecial,
		MinLower:   azdMetadata.AutoGenerateConfig.MinLower,
		MinUpper:   azdMetadata.AutoGenerateConfig.MinUpper,
		MinNumeric: azdMetadata.AutoGenerateConfig.MinNumeric,
		MinSpecial: azdMetadata.AutoGenerateConfig.MinSpecial,
	})
	if err != nil {
		return "", err
	}
	return genValue, nil
}

// locationsWithQuotaFor checks which locations have available quota for a specified SKU.
// It concurrently queries the Azure API for usage data in each location and filters the results
// based on the specified quota and capacity requirements.
//
// Parameters:
//   - ctx: The context for controlling cancellation and deadlines.
//   - subId: The subscription ID to query against.
//   - locations: A list of Azure locations to check for quota availability.
//   - quotaFor: The SKU name and optional capacity (comma-separated) to check for.
//
// Returns:
//   - A slice of location strings that have the required quota and capacity available.
//   - An error if any issues occur during the process or if no locations meet the criteria.
//
// The function first queries the Azure API for usage data in each location concurrently.
// It then filters the results based on the specified SKU name and capacity requirements.
// If no locations meet the criteria, it returns an error with details about the maximum available capacity found.
func (a *BicepProvider) locationsWithQuotaFor(
	ctx context.Context, subId string, locations []string, quotaFor string) ([]string, error) {
	var sharedResults sync.Map
	var wg sync.WaitGroup
	for _, location := range locations {
		wg.Add(1)
		go func(location string) {
			defer wg.Done()
			results, err := a.azureClient.GetAiUsages(ctx, subId, location)
			if err != nil {
				// log the error but don't return it
				log.Println("error getting usage for location", location, ":", err)
				return
			}
			sharedResults.Store(location, results)
		}(location)
	}
	wg.Wait()
	var results []string
	skuUsageName := quotaFor
	var skuCapacity float64
	skuCapacity = 1
	if strings.Contains(skuUsageName, ",") {
		parts := strings.Split(skuUsageName, ",")
		skuUsageName = parts[0]
		cap, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("parsing quota '%s': %w", quotaFor, err)
		}
		skuCapacity = cap
	}
	var maxAvailableCap float64
	maxAvailableCapLocation := ""
	sharedResults.Range(func(key, value any) bool {
		usages := value.([]*armcognitiveservices.Usage)
		hasS0SkuQuota := slices.ContainsFunc(usages, func(q *armcognitiveservices.Usage) bool {
			return *q.Name.Value == "OpenAI.S0.AccountCount" && *q.CurrentValue < *q.Limit
		})
		hasQuotaForModel := slices.ContainsFunc(usages, func(q *armcognitiveservices.Usage) bool {
			hasQuota := *q.Name.Value == skuUsageName
			if !hasQuota {
				return false
			}
			remaining := *q.Limit - *q.CurrentValue
			if remaining > maxAvailableCap {
				maxAvailableCap = remaining
				maxAvailableCapLocation = key.(string)
			}
			return *q.Name.Value == skuUsageName && remaining >= skuCapacity
		})
		if hasQuotaForModel && hasS0SkuQuota {
			results = append(results, key.(string))
		}
		return true
	})
	if len(results) == 0 {
		return nil, fmt.Errorf(
			"no locations with quota for model '%s' and capacity '%.0f'. "+
				"Maximum available capacity of %.0f was found in '%s'",
			skuUsageName, skuCapacity, maxAvailableCap, maxAvailableCapLocation)
	}
	return results, nil
}

func (p *BicepProvider) promptForParameter(
	ctx context.Context,
	key string,
	param azure.ArmTemplateParameterDefinition,
) (any, error) {
	securedParam := "parameter"
	isSecuredParam := param.Secure()
	if isSecuredParam {
		securedParam = "secured parameter"
	}
	msg := fmt.Sprintf("Enter a value for the '%s' infrastructure %s:", key, securedParam)
	help, _ := param.Description()
	azdMetadata, _ := param.AzdMetadata()
	paramType := p.mapBicepTypeToInterfaceType(param.Type)

	var value any

	if paramType == provisioning.ParameterTypeString &&
		azdMetadata.Type != nil &&
		*azdMetadata.Type == azure.AzdMetadataTypeLocation {

		// location can be combined with allowedValues and with usageName metadata
		// allowedValues == nil => all locations are allowed
		// allowedValues != nil => only the locations in the allowedValues are allowed
		// usageName != nil => the usageName is validated for quota for each allowed location (this is for Ai models),
		//                     reducing the allowed locations to only those that have quota available
		// usageName == nil => No quota validation is done
		var allowedLocations []string
		if param.AllowedValues != nil {
			allowedLocations = make([]string, len(*param.AllowedValues))
			for i, option := range *param.AllowedValues {
				allowedLocations[i] = option.(string)
			}
		}
		if azdMetadata.UsageName != nil {
			if allowedLocations == nil {
				allLocations, err := p.subscriptionManager.ListLocations(ctx, p.env.GetSubscriptionId())
				if err != nil {
					return nil, fmt.Errorf("listing locations: %w", err)
				}
				allowedLocations = make([]string, len(allLocations))
				for i, location := range allLocations {
					allowedLocations[i] = location.Name
				}
			}
			withQuotaLocations, err := p.locationsWithQuotaFor(
				ctx, p.env.GetSubscriptionId(), allowedLocations, *azdMetadata.UsageName)
			if err != nil {
				return nil, fmt.Errorf("getting locations with quota: %w", err)
			}
			allowedLocations = withQuotaLocations
		}

		location, err := p.prompters.PromptLocation(ctx, p.env.GetSubscriptionId(), msg, func(loc account.Location) bool {
			return locationParameterFilterImpl(allowedLocations, loc)
		}, defaultPromptValue(param))
		if err != nil {
			return nil, err
		}
		value = location
	} else if paramType == provisioning.ParameterTypeString &&
		azdMetadata.Type != nil &&
		*azdMetadata.Type == azure.AzdMetadataTypeResourceGroup {

		p.console.Message(ctx, fmt.Sprintf(
			"Parameter %s requires an %s resource group.", output.WithUnderline("%s", key), output.WithBold("existing")))
		rgName, err := p.prompters.PromptResourceGroup(ctx, prompt.PromptResourceOptions{
			DisableCreateNew: true,
		})
		if err != nil {
			return nil, err
		}
		value = rgName
	} else if paramType == provisioning.ParameterTypeString &&
		azdMetadata.Type != nil &&
		*azdMetadata.Type == azure.AzdMetadataTypeGenerateOrManual {

		var manualUserInput bool
		defaultOption := "Auto generate"
		options := []string{defaultOption, "Manual input"}
		choice, err := p.console.Select(ctx, input.ConsoleOptions{
			Message: fmt.Sprintf(
				"Parameter %s can be either autogenerated or you can enter its value. What would you like to do?", key),
			Options:      options,
			DefaultValue: defaultOption,
		})
		if err != nil {
			return nil, err
		}
		manualUserInput = options[choice] != defaultOption

		if manualUserInput {
			resultValue, err := promptWithValidation(ctx, p.console, input.ConsoleOptions{
				Message:    msg,
				Help:       help,
				IsPassword: isSecuredParam,
			}, convertString, validateLengthRange(key, param.MinLength, param.MaxLength))
			if err != nil {
				return nil, err
			}
			value = resultValue
		} else {
			genValue, err := autoGenerate(key, azdMetadata)
			if err != nil {
				return nil, err
			}
			value = genValue
		}
	} else if param.AllowedValues != nil {
		options := make([]string, 0, len(*param.AllowedValues))
		for _, option := range *param.AllowedValues {
			options = append(options, fmt.Sprintf("%v", option))
		}

		if len(options) == 0 {
			return nil, fmt.Errorf("parameter '%s' has no allowed values defined", key)
		}

		choice, err := p.console.Select(ctx, input.ConsoleOptions{
			Message: msg,
			Help:    help,
			Options: options,
		})
		if err != nil {
			return nil, err
		}
		value = (*param.AllowedValues)[choice]
	} else {
		switch paramType {
		case provisioning.ParameterTypeBoolean:
			options := []string{"False", "True"}
			choice, err := p.console.Select(ctx, input.ConsoleOptions{
				Message: msg,
				Help:    help,
				Options: options,
			})
			if err != nil {
				return nil, err
			}
			value = (options[choice] == "True")
		case provisioning.ParameterTypeNumber:
			userValue, err := promptWithValidation(ctx, p.console, input.ConsoleOptions{
				Message: msg,
				Help:    help,
			}, convertInt, validateValueRange(key, param.MinValue, param.MaxValue))
			if err != nil {
				return nil, err
			}
			value = userValue
		case provisioning.ParameterTypeString:
			userValue, err := promptWithValidation(ctx, p.console, input.ConsoleOptions{
				Message:    msg,
				Help:       help,
				IsPassword: isSecuredParam,
			}, convertString, validateLengthRange(key, param.MinLength, param.MaxLength))
			if err != nil {
				return nil, err
			}
			value = userValue
		case provisioning.ParameterTypeArray:
			userValue, err := promptWithValidation(ctx, p.console, input.ConsoleOptions{
				Message: msg,
				Help:    help,
			}, convertJson[[]any], validateJsonArray)
			if err != nil {
				return nil, err
			}
			value = userValue
		case provisioning.ParameterTypeObject:
			userValue, err := promptWithValidation(ctx, p.console, input.ConsoleOptions{
				Message: msg,
				Help:    help,
			}, convertJson[map[string]any], validateJsonObject)
			if err != nil {
				return nil, err
			}
			value = userValue
		default:
			panic(fmt.Sprintf("unknown parameter type: %s", p.mapBicepTypeToInterfaceType(param.Type)))
		}
	}

	return value, nil
}

// promptWithValidation prompts for a value using the console and then validates that it satisfies all the validation
// functions. If it does, it is converted from a string to a value using the converter and returned. If any validation
// fails, the prompt is retried after printing the error (prefixed with "Error: ") to the console. If there are is an
// error prompting it is returned as is.
func promptWithValidation[T any](
	ctx context.Context,
	console input.Console,
	options input.ConsoleOptions,
	converter func(string) T,
	validators ...func(string) error,
) (T, error) {
	for {
		userValue, err := console.Prompt(ctx, options)
		if err != nil {
			return *new(T), err
		}

		isValid := true

		for _, validator := range validators {
			if err := validator(userValue); err != nil {
				console.Message(ctx, output.WithErrorFormat("Error: %s.", err))
				isValid = false
				break
			}
		}

		if isValid {
			return converter(userValue), nil
		}
	}
}

func convertString(s string) string {
	return s
}

func convertInt(s string) int {
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		panic(fmt.Sprintf("convertInt: %v", err))
	} else {
		return int(i)
	}
}

func convertJson[T any](s string) T {
	var t T
	if err := json.Unmarshal([]byte(s), &t); err != nil {
		panic(fmt.Sprintf("convertJson: %v", err))
	}
	return t
}
