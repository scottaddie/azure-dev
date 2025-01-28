// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: deployment.proto

package azdext

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment *Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
}

func (x *GetDeploymentResponse) Reset() {
	*x = GetDeploymentResponse{}
	mi := &file_deployment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentResponse) ProtoMessage() {}

func (x *GetDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentResponse.ProtoReflect.Descriptor instead.
func (*GetDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *GetDeploymentResponse) GetDeployment() *Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

type GetDeploymentContextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AzureContext *AzureContext `protobuf:"bytes,1,opt,name=AzureContext,proto3" json:"AzureContext,omitempty"`
}

func (x *GetDeploymentContextResponse) Reset() {
	*x = GetDeploymentContextResponse{}
	mi := &file_deployment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDeploymentContextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentContextResponse) ProtoMessage() {}

func (x *GetDeploymentContextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentContextResponse.ProtoReflect.Descriptor instead.
func (*GetDeploymentContextResponse) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{1}
}

func (x *GetDeploymentContextResponse) GetAzureContext() *AzureContext {
	if x != nil {
		return x.AzureContext
	}
	return nil
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location     string            `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	DeploymentId string            `protobuf:"bytes,3,opt,name=deploymentId,proto3" json:"deploymentId,omitempty"`
	Name         string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Type         string            `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Tags         map[string]string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Outputs      map[string]string `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Resources    []string          `protobuf:"bytes,8,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	mi := &file_deployment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{2}
}

func (x *Deployment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Deployment) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Deployment) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *Deployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deployment) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Deployment) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Deployment) GetOutputs() map[string]string {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *Deployment) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_deployment_proto protoreflect.FileDescriptor

var file_deployment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x58, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x7a,
	0x64, 0x65, 0x78, 0x74, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x52, 0x0c, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22,
	0x84, 0x03, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x7a, 0x64, 0x65,
	0x78, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xad, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e,
	0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x7a, 0x64,
	0x65, 0x78, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x61, 0x7a, 0x64, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x61, 0x7a, 0x64, 0x65,
	0x78, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deployment_proto_rawDescOnce sync.Once
	file_deployment_proto_rawDescData = file_deployment_proto_rawDesc
)

func file_deployment_proto_rawDescGZIP() []byte {
	file_deployment_proto_rawDescOnce.Do(func() {
		file_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_deployment_proto_rawDescData)
	})
	return file_deployment_proto_rawDescData
}

var file_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_deployment_proto_goTypes = []any{
	(*GetDeploymentResponse)(nil),        // 0: azdext.GetDeploymentResponse
	(*GetDeploymentContextResponse)(nil), // 1: azdext.GetDeploymentContextResponse
	(*Deployment)(nil),                   // 2: azdext.Deployment
	nil,                                  // 3: azdext.Deployment.TagsEntry
	nil,                                  // 4: azdext.Deployment.OutputsEntry
	(*AzureContext)(nil),                 // 5: azdext.AzureContext
	(*EmptyRequest)(nil),                 // 6: azdext.EmptyRequest
}
var file_deployment_proto_depIdxs = []int32{
	2, // 0: azdext.GetDeploymentResponse.deployment:type_name -> azdext.Deployment
	5, // 1: azdext.GetDeploymentContextResponse.AzureContext:type_name -> azdext.AzureContext
	3, // 2: azdext.Deployment.tags:type_name -> azdext.Deployment.TagsEntry
	4, // 3: azdext.Deployment.outputs:type_name -> azdext.Deployment.OutputsEntry
	6, // 4: azdext.DeploymentService.GetDeployment:input_type -> azdext.EmptyRequest
	6, // 5: azdext.DeploymentService.GetDeploymentContext:input_type -> azdext.EmptyRequest
	0, // 6: azdext.DeploymentService.GetDeployment:output_type -> azdext.GetDeploymentResponse
	1, // 7: azdext.DeploymentService.GetDeploymentContext:output_type -> azdext.GetDeploymentContextResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_deployment_proto_init() }
func file_deployment_proto_init() {
	if File_deployment_proto != nil {
		return
	}
	file_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deployment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deployment_proto_goTypes,
		DependencyIndexes: file_deployment_proto_depIdxs,
		MessageInfos:      file_deployment_proto_msgTypes,
	}.Build()
	File_deployment_proto = out.File
	file_deployment_proto_rawDesc = nil
	file_deployment_proto_goTypes = nil
	file_deployment_proto_depIdxs = nil
}
