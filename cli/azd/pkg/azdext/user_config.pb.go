// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: user_config.proto

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

// Request message for Get
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_user_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Response message for Get
type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Found bool   `protobuf:"varint,2,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_user_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *GetResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

// Request message for GetString
type GetStringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetStringRequest) Reset() {
	*x = GetStringRequest{}
	mi := &file_user_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStringRequest) ProtoMessage() {}

func (x *GetStringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStringRequest.ProtoReflect.Descriptor instead.
func (*GetStringRequest) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{2}
}

func (x *GetStringRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Response message for GetString
type GetStringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Found bool   `protobuf:"varint,2,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *GetStringResponse) Reset() {
	*x = GetStringResponse{}
	mi := &file_user_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStringResponse) ProtoMessage() {}

func (x *GetStringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStringResponse.ProtoReflect.Descriptor instead.
func (*GetStringResponse) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{3}
}

func (x *GetStringResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetStringResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

// Request message for GetSection
type GetSectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetSectionRequest) Reset() {
	*x = GetSectionRequest{}
	mi := &file_user_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSectionRequest) ProtoMessage() {}

func (x *GetSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSectionRequest.ProtoReflect.Descriptor instead.
func (*GetSectionRequest) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{4}
}

func (x *GetSectionRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Response message for GetSection
type GetSectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section []byte `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
	Found   bool   `protobuf:"varint,2,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *GetSectionResponse) Reset() {
	*x = GetSectionResponse{}
	mi := &file_user_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSectionResponse) ProtoMessage() {}

func (x *GetSectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSectionResponse.ProtoReflect.Descriptor instead.
func (*GetSectionResponse) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{5}
}

func (x *GetSectionResponse) GetSection() []byte {
	if x != nil {
		return x.Section
	}
	return nil
}

func (x *GetSectionResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

// Request message for Set
type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	mi := &file_user_config_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{6}
}

func (x *SetRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SetRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Response message for Set
type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	mi := &file_user_config_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{7}
}

func (x *SetResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Request message for Unset
type UnsetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *UnsetRequest) Reset() {
	*x = UnsetRequest{}
	mi := &file_user_config_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsetRequest) ProtoMessage() {}

func (x *UnsetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsetRequest.ProtoReflect.Descriptor instead.
func (*UnsetRequest) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{8}
}

func (x *UnsetRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Response message for Unset
type UnsetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UnsetResponse) Reset() {
	*x = UnsetResponse{}
	mi := &file_user_config_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsetResponse) ProtoMessage() {}

func (x *UnsetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_config_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsetResponse.ProtoReflect.Descriptor instead.
func (*UnsetResponse) Descriptor() ([]byte, []int) {
	return file_user_config_proto_rawDescGZIP(), []int{9}
}

func (x *UnsetResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_user_config_proto protoreflect.FileDescriptor

var file_user_config_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x22, 0x20, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x39, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x44, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x22, 0x36, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x22, 0x0a, 0x0c, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x22, 0x27, 0x0a, 0x0d, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xb0, 0x02, 0x0a,
	0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x7a, 0x64, 0x65,
	0x78, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x18, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x7a, 0x64, 0x65,
	0x78, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x53, 0x65, 0x74,
	0x12, 0x12, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x55, 0x6e, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x55, 0x6e, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x7a, 0x64, 0x65, 0x78,
	0x74, 0x2e, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x7a,
	0x75, 0x72, 0x65, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x6c,
	0x69, 0x2f, 0x61, 0x7a, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74,
	0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x61, 0x7a, 0x64, 0x65, 0x78, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_config_proto_rawDescOnce sync.Once
	file_user_config_proto_rawDescData = file_user_config_proto_rawDesc
)

func file_user_config_proto_rawDescGZIP() []byte {
	file_user_config_proto_rawDescOnce.Do(func() {
		file_user_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_config_proto_rawDescData)
	})
	return file_user_config_proto_rawDescData
}

var file_user_config_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_config_proto_goTypes = []any{
	(*GetRequest)(nil),         // 0: azdext.GetRequest
	(*GetResponse)(nil),        // 1: azdext.GetResponse
	(*GetStringRequest)(nil),   // 2: azdext.GetStringRequest
	(*GetStringResponse)(nil),  // 3: azdext.GetStringResponse
	(*GetSectionRequest)(nil),  // 4: azdext.GetSectionRequest
	(*GetSectionResponse)(nil), // 5: azdext.GetSectionResponse
	(*SetRequest)(nil),         // 6: azdext.SetRequest
	(*SetResponse)(nil),        // 7: azdext.SetResponse
	(*UnsetRequest)(nil),       // 8: azdext.UnsetRequest
	(*UnsetResponse)(nil),      // 9: azdext.UnsetResponse
}
var file_user_config_proto_depIdxs = []int32{
	0, // 0: azdext.UserConfigService.Get:input_type -> azdext.GetRequest
	2, // 1: azdext.UserConfigService.GetString:input_type -> azdext.GetStringRequest
	4, // 2: azdext.UserConfigService.GetSection:input_type -> azdext.GetSectionRequest
	6, // 3: azdext.UserConfigService.Set:input_type -> azdext.SetRequest
	8, // 4: azdext.UserConfigService.Unset:input_type -> azdext.UnsetRequest
	1, // 5: azdext.UserConfigService.Get:output_type -> azdext.GetResponse
	3, // 6: azdext.UserConfigService.GetString:output_type -> azdext.GetStringResponse
	5, // 7: azdext.UserConfigService.GetSection:output_type -> azdext.GetSectionResponse
	7, // 8: azdext.UserConfigService.Set:output_type -> azdext.SetResponse
	9, // 9: azdext.UserConfigService.Unset:output_type -> azdext.UnsetResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_config_proto_init() }
func file_user_config_proto_init() {
	if File_user_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_config_proto_goTypes,
		DependencyIndexes: file_user_config_proto_depIdxs,
		MessageInfos:      file_user_config_proto_msgTypes,
	}.Build()
	File_user_config_proto = out.File
	file_user_config_proto_rawDesc = nil
	file_user_config_proto_goTypes = nil
	file_user_config_proto_depIdxs = nil
}
