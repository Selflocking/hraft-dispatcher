// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: command/command.proto

package command

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

type Command_Type int32

const (
	Command_COMMAND_TYPE_ADD             Command_Type = 0
	Command_COMMAND_TYPE_REMOVE          Command_Type = 1
	Command_COMMAND_TYPE_REMOVE_FILTERED Command_Type = 2
	Command_COMMAND_TYPE_UPDATE          Command_Type = 4
	Command_COMMAND_TYPE_CLEAR           Command_Type = 3
)

// Enum value maps for Command_Type.
var (
	Command_Type_name = map[int32]string{
		0: "COMMAND_TYPE_ADD",
		1: "COMMAND_TYPE_REMOVE",
		2: "COMMAND_TYPE_REMOVE_FILTERED",
		4: "COMMAND_TYPE_UPDATE",
		3: "COMMAND_TYPE_CLEAR",
	}
	Command_Type_value = map[string]int32{
		"COMMAND_TYPE_ADD":             0,
		"COMMAND_TYPE_REMOVE":          1,
		"COMMAND_TYPE_REMOVE_FILTERED": 2,
		"COMMAND_TYPE_UPDATE":          4,
		"COMMAND_TYPE_CLEAR":           3,
	}
)

func (x Command_Type) Enum() *Command_Type {
	p := new(Command_Type)
	*p = x
	return p
}

func (x Command_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_command_command_proto_enumTypes[0].Descriptor()
}

func (Command_Type) Type() protoreflect.EnumType {
	return &file_command_command_proto_enumTypes[0]
}

func (x Command_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command_Type.Descriptor instead.
func (Command_Type) EnumDescriptor() ([]byte, []int) {
	return file_command_command_proto_rawDescGZIP(), []int{5, 0}
}

type StringArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *StringArray) Reset() {
	*x = StringArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringArray) ProtoMessage() {}

func (x *StringArray) ProtoReflect() protoreflect.Message {
	mi := &file_command_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringArray.ProtoReflect.Descriptor instead.
func (*StringArray) Descriptor() ([]byte, []int) {
	return file_command_command_proto_rawDescGZIP(), []int{0}
}

func (x *StringArray) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

type AddPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sec   string         `protobuf:"bytes,1,opt,name=sec,proto3" json:"sec,omitempty"`
	PType string         `protobuf:"bytes,2,opt,name=pType,proto3" json:"pType,omitempty"`
	Rules []*StringArray `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *AddPolicyRequest) Reset() {
	*x = AddPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPolicyRequest) ProtoMessage() {}

func (x *AddPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPolicyRequest.ProtoReflect.Descriptor instead.
func (*AddPolicyRequest) Descriptor() ([]byte, []int) {
	return file_command_command_proto_rawDescGZIP(), []int{1}
}

func (x *AddPolicyRequest) GetSec() string {
	if x != nil {
		return x.Sec
	}
	return ""
}

func (x *AddPolicyRequest) GetPType() string {
	if x != nil {
		return x.PType
	}
	return ""
}

func (x *AddPolicyRequest) GetRules() []*StringArray {
	if x != nil {
		return x.Rules
	}
	return nil
}

type RemovePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sec   string         `protobuf:"bytes,1,opt,name=sec,proto3" json:"sec,omitempty"`
	PType string         `protobuf:"bytes,2,opt,name=pType,proto3" json:"pType,omitempty"`
	Rules []*StringArray `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *RemovePolicyRequest) Reset() {
	*x = RemovePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePolicyRequest) ProtoMessage() {}

func (x *RemovePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePolicyRequest.ProtoReflect.Descriptor instead.
func (*RemovePolicyRequest) Descriptor() ([]byte, []int) {
	return file_command_command_proto_rawDescGZIP(), []int{2}
}

func (x *RemovePolicyRequest) GetSec() string {
	if x != nil {
		return x.Sec
	}
	return ""
}

func (x *RemovePolicyRequest) GetPType() string {
	if x != nil {
		return x.PType
	}
	return ""
}

func (x *RemovePolicyRequest) GetRules() []*StringArray {
	if x != nil {
		return x.Rules
	}
	return nil
}

type RemoveFilteredPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sec         string   `protobuf:"bytes,1,opt,name=sec,proto3" json:"sec,omitempty"`
	PType       string   `protobuf:"bytes,2,opt,name=pType,proto3" json:"pType,omitempty"`
	FieldIndex  int32    `protobuf:"varint,3,opt,name=fieldIndex,proto3" json:"fieldIndex,omitempty"`
	FieldValues []string `protobuf:"bytes,4,rep,name=fieldValues,proto3" json:"fieldValues,omitempty"`
}

func (x *RemoveFilteredPolicyRequest) Reset() {
	*x = RemoveFilteredPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFilteredPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFilteredPolicyRequest) ProtoMessage() {}

func (x *RemoveFilteredPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFilteredPolicyRequest.ProtoReflect.Descriptor instead.
func (*RemoveFilteredPolicyRequest) Descriptor() ([]byte, []int) {
	return file_command_command_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveFilteredPolicyRequest) GetSec() string {
	if x != nil {
		return x.Sec
	}
	return ""
}

func (x *RemoveFilteredPolicyRequest) GetPType() string {
	if x != nil {
		return x.PType
	}
	return ""
}

func (x *RemoveFilteredPolicyRequest) GetFieldIndex() int32 {
	if x != nil {
		return x.FieldIndex
	}
	return 0
}

func (x *RemoveFilteredPolicyRequest) GetFieldValues() []string {
	if x != nil {
		return x.FieldValues
	}
	return nil
}

type UpdatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sec     string   `protobuf:"bytes,1,opt,name=sec,proto3" json:"sec,omitempty"`
	PType   string   `protobuf:"bytes,2,opt,name=pType,proto3" json:"pType,omitempty"`
	NewRule []string `protobuf:"bytes,3,rep,name=newRule,proto3" json:"newRule,omitempty"`
	OldRule []string `protobuf:"bytes,4,rep,name=oldRule,proto3" json:"oldRule,omitempty"`
}

func (x *UpdatePolicyRequest) Reset() {
	*x = UpdatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePolicyRequest) ProtoMessage() {}

func (x *UpdatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_command_command_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePolicyRequest) GetSec() string {
	if x != nil {
		return x.Sec
	}
	return ""
}

func (x *UpdatePolicyRequest) GetPType() string {
	if x != nil {
		return x.PType
	}
	return ""
}

func (x *UpdatePolicyRequest) GetNewRule() []string {
	if x != nil {
		return x.NewRule
	}
	return nil
}

func (x *UpdatePolicyRequest) GetOldRule() []string {
	if x != nil {
		return x.OldRule
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Command_Type `protobuf:"varint,1,opt,name=type,proto3,enum=command.Command_Type" json:"type,omitempty"`
	Data []byte       `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_command_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_command_command_proto_rawDescGZIP(), []int{5}
}

func (x *Command) GetType() Command_Type {
	if x != nil {
		return x.Type
	}
	return Command_COMMAND_TYPE_ADD
}

func (x *Command) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddNodeRequest) Reset() {
	*x = AddNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_command_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNodeRequest) ProtoMessage() {}

func (x *AddNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_command_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNodeRequest.ProtoReflect.Descriptor instead.
func (*AddNodeRequest) Descriptor() ([]byte, []int) {
	return file_command_command_proto_rawDescGZIP(), []int{6}
}

func (x *AddNodeRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddNodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveNodeRequest) Reset() {
	*x = RemoveNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_command_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveNodeRequest) ProtoMessage() {}

func (x *RemoveNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_command_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveNodeRequest.ProtoReflect.Descriptor instead.
func (*RemoveNodeRequest) Descriptor() ([]byte, []int) {
	return file_command_command_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveNodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_command_command_proto protoreflect.FileDescriptor

var file_command_command_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0x23, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x66, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2a, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x69, 0x0a,
	0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x1b, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0x71, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x52, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c,
	0x64, 0x52, 0x75, 0x6c, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x88, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4d,
	0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x10, 0x00, 0x12,
	0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4d, 0x4d,
	0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f,
	0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f,
	0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x10, 0x03, 0x22, 0x3a, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x63,
	0x65, 0x2f, 0x63, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x2d, 0x68, 0x72, 0x61, 0x66, 0x74, 0x2d, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_command_proto_rawDescOnce sync.Once
	file_command_command_proto_rawDescData = file_command_command_proto_rawDesc
)

func file_command_command_proto_rawDescGZIP() []byte {
	file_command_command_proto_rawDescOnce.Do(func() {
		file_command_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_command_proto_rawDescData)
	})
	return file_command_command_proto_rawDescData
}

var file_command_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_command_command_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_command_command_proto_goTypes = []interface{}{
	(Command_Type)(0),                   // 0: command.Command.Type
	(*StringArray)(nil),                 // 1: command.StringArray
	(*AddPolicyRequest)(nil),            // 2: command.AddPolicyRequest
	(*RemovePolicyRequest)(nil),         // 3: command.RemovePolicyRequest
	(*RemoveFilteredPolicyRequest)(nil), // 4: command.RemoveFilteredPolicyRequest
	(*UpdatePolicyRequest)(nil),         // 5: command.UpdatePolicyRequest
	(*Command)(nil),                     // 6: command.Command
	(*AddNodeRequest)(nil),              // 7: command.AddNodeRequest
	(*RemoveNodeRequest)(nil),           // 8: command.RemoveNodeRequest
}
var file_command_command_proto_depIdxs = []int32{
	1, // 0: command.AddPolicyRequest.rules:type_name -> command.StringArray
	1, // 1: command.RemovePolicyRequest.rules:type_name -> command.StringArray
	0, // 2: command.Command.type:type_name -> command.Command.Type
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_command_command_proto_init() }
func file_command_command_proto_init() {
	if File_command_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_command_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringArray); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPolicyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePolicyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFilteredPolicyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePolicyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_command_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_command_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_command_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_command_command_proto_goTypes,
		DependencyIndexes: file_command_command_proto_depIdxs,
		EnumInfos:         file_command_command_proto_enumTypes,
		MessageInfos:      file_command_command_proto_msgTypes,
	}.Build()
	File_command_command_proto = out.File
	file_command_command_proto_rawDesc = nil
	file_command_command_proto_goTypes = nil
	file_command_command_proto_depIdxs = nil
}