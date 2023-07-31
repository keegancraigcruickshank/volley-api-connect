// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: divisions/v1/private/divisions.proto

package divisionsv1private

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

type SortDirection int32

const (
	SortDirection_ASC  SortDirection = 0
	SortDirection_DESC SortDirection = 1
)

// Enum value maps for SortDirection.
var (
	SortDirection_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	SortDirection_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x SortDirection) Enum() *SortDirection {
	p := new(SortDirection)
	*p = x
	return p
}

func (x SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_divisions_v1_private_divisions_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_divisions_v1_private_divisions_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_divisions_v1_private_divisions_proto_rawDescGZIP(), []int{0}
}

type Division struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Days     []string `protobuf:"bytes,3,rep,name=days,proto3" json:"days,omitempty"`
	Modified int32    `protobuf:"varint,4,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *Division) Reset() {
	*x = Division{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divisions_v1_private_divisions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Division) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Division) ProtoMessage() {}

func (x *Division) ProtoReflect() protoreflect.Message {
	mi := &file_divisions_v1_private_divisions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Division.ProtoReflect.Descriptor instead.
func (*Division) Descriptor() ([]byte, []int) {
	return file_divisions_v1_private_divisions_proto_rawDescGZIP(), []int{0}
}

func (x *Division) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Division) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Division) GetDays() []string {
	if x != nil {
		return x.Days
	}
	return nil
}

func (x *Division) GetModified() int32 {
	if x != nil {
		return x.Modified
	}
	return 0
}

type AddDivisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Days []string `protobuf:"bytes,2,rep,name=days,proto3" json:"days,omitempty"`
}

func (x *AddDivisionRequest) Reset() {
	*x = AddDivisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divisions_v1_private_divisions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDivisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDivisionRequest) ProtoMessage() {}

func (x *AddDivisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_divisions_v1_private_divisions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDivisionRequest.ProtoReflect.Descriptor instead.
func (*AddDivisionRequest) Descriptor() ([]byte, []int) {
	return file_divisions_v1_private_divisions_proto_rawDescGZIP(), []int{1}
}

func (x *AddDivisionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddDivisionRequest) GetDays() []string {
	if x != nil {
		return x.Days
	}
	return nil
}

type AddDivisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddDivisionResponse) Reset() {
	*x = AddDivisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divisions_v1_private_divisions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDivisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDivisionResponse) ProtoMessage() {}

func (x *AddDivisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_divisions_v1_private_divisions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDivisionResponse.ProtoReflect.Descriptor instead.
func (*AddDivisionResponse) Descriptor() ([]byte, []int) {
	return file_divisions_v1_private_divisions_proto_rawDescGZIP(), []int{2}
}

type ListDivisionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         *string        `protobuf:"bytes,1,opt,name=query,proto3,oneof" json:"query,omitempty"`
	Page          *int32         `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
	PageSize      *int32         `protobuf:"varint,3,opt,name=pageSize,proto3,oneof" json:"pageSize,omitempty"`
	SortDirection *SortDirection `protobuf:"varint,4,opt,name=sortDirection,proto3,enum=divisions.v1.private.SortDirection,oneof" json:"sortDirection,omitempty"`
}

func (x *ListDivisionsRequest) Reset() {
	*x = ListDivisionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divisions_v1_private_divisions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDivisionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDivisionsRequest) ProtoMessage() {}

func (x *ListDivisionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_divisions_v1_private_divisions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDivisionsRequest.ProtoReflect.Descriptor instead.
func (*ListDivisionsRequest) Descriptor() ([]byte, []int) {
	return file_divisions_v1_private_divisions_proto_rawDescGZIP(), []int{3}
}

func (x *ListDivisionsRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *ListDivisionsRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListDivisionsRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListDivisionsRequest) GetSortDirection() SortDirection {
	if x != nil && x.SortDirection != nil {
		return *x.SortDirection
	}
	return SortDirection_ASC
}

type ListDivisionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Divisions  []*Division `protobuf:"bytes,1,rep,name=divisions,proto3" json:"divisions,omitempty"`
	TotalCount int32       `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListDivisionsResponse) Reset() {
	*x = ListDivisionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divisions_v1_private_divisions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDivisionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDivisionsResponse) ProtoMessage() {}

func (x *ListDivisionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_divisions_v1_private_divisions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDivisionsResponse.ProtoReflect.Descriptor instead.
func (*ListDivisionsResponse) Descriptor() ([]byte, []int) {
	return file_divisions_v1_private_divisions_proto_rawDescGZIP(), []int{4}
}

func (x *ListDivisionsResponse) GetDivisions() []*Division {
	if x != nil {
		return x.Divisions
	}
	return nil
}

func (x *ListDivisionsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type RemoveDivisionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *RemoveDivisionsRequest) Reset() {
	*x = RemoveDivisionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divisions_v1_private_divisions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDivisionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDivisionsRequest) ProtoMessage() {}

func (x *RemoveDivisionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_divisions_v1_private_divisions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDivisionsRequest.ProtoReflect.Descriptor instead.
func (*RemoveDivisionsRequest) Descriptor() ([]byte, []int) {
	return file_divisions_v1_private_divisions_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveDivisionsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type RemoveDivisionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveDivisionsResponse) Reset() {
	*x = RemoveDivisionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_divisions_v1_private_divisions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDivisionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDivisionsResponse) ProtoMessage() {}

func (x *RemoveDivisionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_divisions_v1_private_divisions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDivisionsResponse.ProtoReflect.Descriptor instead.
func (*RemoveDivisionsResponse) Descriptor() ([]byte, []int) {
	return file_divisions_v1_private_divisions_proto_rawDescGZIP(), []int{6}
}

var File_divisions_v1_private_divisions_proto protoreflect.FileDescriptor

var file_divisions_v1_private_divisions_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x5e, 0x0a, 0x08,
	0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x3c, 0x0a, 0x12,
	0x41, 0x64, 0x64, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x64,
	0x64, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xed, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x02, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x4e, 0x0a, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x0d, 0x73,
	0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x75, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x64, 0x69,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x69,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a,
	0x22, 0x0a, 0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53,
	0x43, 0x10, 0x01, 0x32, 0xdd, 0x02, 0x0a, 0x17, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x64, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x2e, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x65, 0x65, 0x67, 0x61, 0x6e, 0x63, 0x72, 0x61, 0x69, 0x67, 0x63, 0x72, 0x75,
	0x69, 0x63, 0x6b, 0x73, 0x68, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x6f, 0x6c, 0x6c, 0x65, 0x79, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x64, 0x69, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x3b, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x76, 0x31, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_divisions_v1_private_divisions_proto_rawDescOnce sync.Once
	file_divisions_v1_private_divisions_proto_rawDescData = file_divisions_v1_private_divisions_proto_rawDesc
)

func file_divisions_v1_private_divisions_proto_rawDescGZIP() []byte {
	file_divisions_v1_private_divisions_proto_rawDescOnce.Do(func() {
		file_divisions_v1_private_divisions_proto_rawDescData = protoimpl.X.CompressGZIP(file_divisions_v1_private_divisions_proto_rawDescData)
	})
	return file_divisions_v1_private_divisions_proto_rawDescData
}

var file_divisions_v1_private_divisions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_divisions_v1_private_divisions_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_divisions_v1_private_divisions_proto_goTypes = []interface{}{
	(SortDirection)(0),              // 0: divisions.v1.private.SortDirection
	(*Division)(nil),                // 1: divisions.v1.private.Division
	(*AddDivisionRequest)(nil),      // 2: divisions.v1.private.AddDivisionRequest
	(*AddDivisionResponse)(nil),     // 3: divisions.v1.private.AddDivisionResponse
	(*ListDivisionsRequest)(nil),    // 4: divisions.v1.private.ListDivisionsRequest
	(*ListDivisionsResponse)(nil),   // 5: divisions.v1.private.ListDivisionsResponse
	(*RemoveDivisionsRequest)(nil),  // 6: divisions.v1.private.RemoveDivisionsRequest
	(*RemoveDivisionsResponse)(nil), // 7: divisions.v1.private.RemoveDivisionsResponse
}
var file_divisions_v1_private_divisions_proto_depIdxs = []int32{
	0, // 0: divisions.v1.private.ListDivisionsRequest.sortDirection:type_name -> divisions.v1.private.SortDirection
	1, // 1: divisions.v1.private.ListDivisionsResponse.divisions:type_name -> divisions.v1.private.Division
	2, // 2: divisions.v1.private.PrivateDivisionsService.AddDivision:input_type -> divisions.v1.private.AddDivisionRequest
	6, // 3: divisions.v1.private.PrivateDivisionsService.RemoveDivisions:input_type -> divisions.v1.private.RemoveDivisionsRequest
	4, // 4: divisions.v1.private.PrivateDivisionsService.ListDivisions:input_type -> divisions.v1.private.ListDivisionsRequest
	3, // 5: divisions.v1.private.PrivateDivisionsService.AddDivision:output_type -> divisions.v1.private.AddDivisionResponse
	7, // 6: divisions.v1.private.PrivateDivisionsService.RemoveDivisions:output_type -> divisions.v1.private.RemoveDivisionsResponse
	5, // 7: divisions.v1.private.PrivateDivisionsService.ListDivisions:output_type -> divisions.v1.private.ListDivisionsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_divisions_v1_private_divisions_proto_init() }
func file_divisions_v1_private_divisions_proto_init() {
	if File_divisions_v1_private_divisions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_divisions_v1_private_divisions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Division); i {
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
		file_divisions_v1_private_divisions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDivisionRequest); i {
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
		file_divisions_v1_private_divisions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDivisionResponse); i {
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
		file_divisions_v1_private_divisions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDivisionsRequest); i {
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
		file_divisions_v1_private_divisions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDivisionsResponse); i {
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
		file_divisions_v1_private_divisions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDivisionsRequest); i {
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
		file_divisions_v1_private_divisions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDivisionsResponse); i {
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
	file_divisions_v1_private_divisions_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_divisions_v1_private_divisions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_divisions_v1_private_divisions_proto_goTypes,
		DependencyIndexes: file_divisions_v1_private_divisions_proto_depIdxs,
		EnumInfos:         file_divisions_v1_private_divisions_proto_enumTypes,
		MessageInfos:      file_divisions_v1_private_divisions_proto_msgTypes,
	}.Build()
	File_divisions_v1_private_divisions_proto = out.File
	file_divisions_v1_private_divisions_proto_rawDesc = nil
	file_divisions_v1_private_divisions_proto_goTypes = nil
	file_divisions_v1_private_divisions_proto_depIdxs = nil
}
