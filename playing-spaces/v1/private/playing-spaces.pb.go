// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: playing-spaces/v1/private/playing-spaces.proto

package playingspacesv1private

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
	return file_playing_spaces_v1_private_playing_spaces_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_playing_spaces_v1_private_playing_spaces_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{0}
}

type PlayingSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LocationId   string `protobuf:"bytes,2,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	LocationName string `protobuf:"bytes,3,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PlayingSpace) Reset() {
	*x = PlayingSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayingSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayingSpace) ProtoMessage() {}

func (x *PlayingSpace) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayingSpace.ProtoReflect.Descriptor instead.
func (*PlayingSpace) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{0}
}

func (x *PlayingSpace) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlayingSpace) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *PlayingSpace) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *PlayingSpace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreatePlayingSpaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationId string `protobuf:"bytes,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreatePlayingSpaceRequest) Reset() {
	*x = CreatePlayingSpaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayingSpaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayingSpaceRequest) ProtoMessage() {}

func (x *CreatePlayingSpaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayingSpaceRequest.ProtoReflect.Descriptor instead.
func (*CreatePlayingSpaceRequest) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePlayingSpaceRequest) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *CreatePlayingSpaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreatePlayingSpaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayingSpace *PlayingSpace `protobuf:"bytes,1,opt,name=playingSpace,proto3" json:"playingSpace,omitempty"`
}

func (x *CreatePlayingSpaceResponse) Reset() {
	*x = CreatePlayingSpaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayingSpaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayingSpaceResponse) ProtoMessage() {}

func (x *CreatePlayingSpaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayingSpaceResponse.ProtoReflect.Descriptor instead.
func (*CreatePlayingSpaceResponse) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePlayingSpaceResponse) GetPlayingSpace() *PlayingSpace {
	if x != nil {
		return x.PlayingSpace
	}
	return nil
}

type RemovePlayingSpacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayingSpaceIds []string `protobuf:"bytes,1,rep,name=playingSpaceIds,proto3" json:"playingSpaceIds,omitempty"`
}

func (x *RemovePlayingSpacesRequest) Reset() {
	*x = RemovePlayingSpacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePlayingSpacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePlayingSpacesRequest) ProtoMessage() {}

func (x *RemovePlayingSpacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePlayingSpacesRequest.ProtoReflect.Descriptor instead.
func (*RemovePlayingSpacesRequest) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{3}
}

func (x *RemovePlayingSpacesRequest) GetPlayingSpaceIds() []string {
	if x != nil {
		return x.PlayingSpaceIds
	}
	return nil
}

type RemovePlayingSpacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemovePlayingSpacesResponse) Reset() {
	*x = RemovePlayingSpacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePlayingSpacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePlayingSpacesResponse) ProtoMessage() {}

func (x *RemovePlayingSpacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePlayingSpacesResponse.ProtoReflect.Descriptor instead.
func (*RemovePlayingSpacesResponse) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{4}
}

type ModifyPlayingSpaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	LocationId *string `protobuf:"bytes,3,opt,name=location_id,json=locationId,proto3,oneof" json:"location_id,omitempty"`
}

func (x *ModifyPlayingSpaceRequest) Reset() {
	*x = ModifyPlayingSpaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyPlayingSpaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyPlayingSpaceRequest) ProtoMessage() {}

func (x *ModifyPlayingSpaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyPlayingSpaceRequest.ProtoReflect.Descriptor instead.
func (*ModifyPlayingSpaceRequest) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{5}
}

func (x *ModifyPlayingSpaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ModifyPlayingSpaceRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ModifyPlayingSpaceRequest) GetLocationId() string {
	if x != nil && x.LocationId != nil {
		return *x.LocationId
	}
	return ""
}

type ModifyPlayingSpaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayingSpace *PlayingSpace `protobuf:"bytes,1,opt,name=playingSpace,proto3" json:"playingSpace,omitempty"`
}

func (x *ModifyPlayingSpaceResponse) Reset() {
	*x = ModifyPlayingSpaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyPlayingSpaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyPlayingSpaceResponse) ProtoMessage() {}

func (x *ModifyPlayingSpaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyPlayingSpaceResponse.ProtoReflect.Descriptor instead.
func (*ModifyPlayingSpaceResponse) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{6}
}

func (x *ModifyPlayingSpaceResponse) GetPlayingSpace() *PlayingSpace {
	if x != nil {
		return x.PlayingSpace
	}
	return nil
}

type ListPlayingSpacesFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayingSpaceIds []string `protobuf:"bytes,1,rep,name=playing_space_ids,json=playingSpaceIds,proto3" json:"playing_space_ids,omitempty"`
}

func (x *ListPlayingSpacesFilter) Reset() {
	*x = ListPlayingSpacesFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlayingSpacesFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlayingSpacesFilter) ProtoMessage() {}

func (x *ListPlayingSpacesFilter) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlayingSpacesFilter.ProtoReflect.Descriptor instead.
func (*ListPlayingSpacesFilter) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{7}
}

func (x *ListPlayingSpacesFilter) GetPlayingSpaceIds() []string {
	if x != nil {
		return x.PlayingSpaceIds
	}
	return nil
}

type ListPlayingSpacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         *string                  `protobuf:"bytes,1,opt,name=query,proto3,oneof" json:"query,omitempty"`
	Page          *int32                   `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
	PageSize      *int32                   `protobuf:"varint,3,opt,name=pageSize,proto3,oneof" json:"pageSize,omitempty"`
	SortDirection *SortDirection           `protobuf:"varint,4,opt,name=sortDirection,proto3,enum=playing_spaces.v1.private.SortDirection,oneof" json:"sortDirection,omitempty"`
	Filter        *ListPlayingSpacesFilter `protobuf:"bytes,5,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
}

func (x *ListPlayingSpacesRequest) Reset() {
	*x = ListPlayingSpacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlayingSpacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlayingSpacesRequest) ProtoMessage() {}

func (x *ListPlayingSpacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlayingSpacesRequest.ProtoReflect.Descriptor instead.
func (*ListPlayingSpacesRequest) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{8}
}

func (x *ListPlayingSpacesRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *ListPlayingSpacesRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListPlayingSpacesRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListPlayingSpacesRequest) GetSortDirection() SortDirection {
	if x != nil && x.SortDirection != nil {
		return *x.SortDirection
	}
	return SortDirection_ASC
}

func (x *ListPlayingSpacesRequest) GetFilter() *ListPlayingSpacesFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListPlayingSpacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayingSpaces []*PlayingSpace `protobuf:"bytes,1,rep,name=playingSpaces,proto3" json:"playingSpaces,omitempty"`
	TotalCount    int32           `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListPlayingSpacesResponse) Reset() {
	*x = ListPlayingSpacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlayingSpacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlayingSpacesResponse) ProtoMessage() {}

func (x *ListPlayingSpacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlayingSpacesResponse.ProtoReflect.Descriptor instead.
func (*ListPlayingSpacesResponse) Descriptor() ([]byte, []int) {
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP(), []int{9}
}

func (x *ListPlayingSpacesResponse) GetPlayingSpaces() []*PlayingSpace {
	if x != nil {
		return x.PlayingSpaces
	}
	return nil
}

func (x *ListPlayingSpacesResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_playing_spaces_v1_private_playing_spaces_proto protoreflect.FileDescriptor

var file_playing_spaces_v1_private_playing_spaces_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79,
	0x69, 0x6e, 0x67, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x78, 0x0a, 0x0c, 0x50,
	0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x22, 0x46, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x69,
	0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x1d, 0x0a, 0x1b, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x19, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22,
	0x69, 0x0a, 0x1a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0c, 0x70, 0x6c,
	0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x73, 0x22, 0xd2, 0x02, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e,
	0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x53, 0x0a, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x04, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x6f,
	0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x8a, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x2a, 0x22, 0x0a, 0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x32, 0xb5, 0x04, 0x0a, 0x1b, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x33, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x34, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x83, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x34, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e,
	0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x86, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69,
	0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x35, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69,
	0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x36, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x83, 0x01, 0x0a, 0x12, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x34, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x67, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65,
	0x65, 0x67, 0x61, 0x6e, 0x63, 0x72, 0x61, 0x69, 0x67, 0x63, 0x72, 0x75, 0x69, 0x63, 0x6b, 0x73,
	0x68, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x6f, 0x6c, 0x6c, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x2d,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x3b, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x76,
	0x31, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playing_spaces_v1_private_playing_spaces_proto_rawDescOnce sync.Once
	file_playing_spaces_v1_private_playing_spaces_proto_rawDescData = file_playing_spaces_v1_private_playing_spaces_proto_rawDesc
)

func file_playing_spaces_v1_private_playing_spaces_proto_rawDescGZIP() []byte {
	file_playing_spaces_v1_private_playing_spaces_proto_rawDescOnce.Do(func() {
		file_playing_spaces_v1_private_playing_spaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_playing_spaces_v1_private_playing_spaces_proto_rawDescData)
	})
	return file_playing_spaces_v1_private_playing_spaces_proto_rawDescData
}

var file_playing_spaces_v1_private_playing_spaces_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_playing_spaces_v1_private_playing_spaces_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_playing_spaces_v1_private_playing_spaces_proto_goTypes = []interface{}{
	(SortDirection)(0),                  // 0: playing_spaces.v1.private.SortDirection
	(*PlayingSpace)(nil),                // 1: playing_spaces.v1.private.PlayingSpace
	(*CreatePlayingSpaceRequest)(nil),   // 2: playing_spaces.v1.private.CreatePlayingSpaceRequest
	(*CreatePlayingSpaceResponse)(nil),  // 3: playing_spaces.v1.private.CreatePlayingSpaceResponse
	(*RemovePlayingSpacesRequest)(nil),  // 4: playing_spaces.v1.private.RemovePlayingSpacesRequest
	(*RemovePlayingSpacesResponse)(nil), // 5: playing_spaces.v1.private.RemovePlayingSpacesResponse
	(*ModifyPlayingSpaceRequest)(nil),   // 6: playing_spaces.v1.private.ModifyPlayingSpaceRequest
	(*ModifyPlayingSpaceResponse)(nil),  // 7: playing_spaces.v1.private.ModifyPlayingSpaceResponse
	(*ListPlayingSpacesFilter)(nil),     // 8: playing_spaces.v1.private.ListPlayingSpacesFilter
	(*ListPlayingSpacesRequest)(nil),    // 9: playing_spaces.v1.private.ListPlayingSpacesRequest
	(*ListPlayingSpacesResponse)(nil),   // 10: playing_spaces.v1.private.ListPlayingSpacesResponse
}
var file_playing_spaces_v1_private_playing_spaces_proto_depIdxs = []int32{
	1,  // 0: playing_spaces.v1.private.CreatePlayingSpaceResponse.playingSpace:type_name -> playing_spaces.v1.private.PlayingSpace
	1,  // 1: playing_spaces.v1.private.ModifyPlayingSpaceResponse.playingSpace:type_name -> playing_spaces.v1.private.PlayingSpace
	0,  // 2: playing_spaces.v1.private.ListPlayingSpacesRequest.sortDirection:type_name -> playing_spaces.v1.private.SortDirection
	8,  // 3: playing_spaces.v1.private.ListPlayingSpacesRequest.filter:type_name -> playing_spaces.v1.private.ListPlayingSpacesFilter
	1,  // 4: playing_spaces.v1.private.ListPlayingSpacesResponse.playingSpaces:type_name -> playing_spaces.v1.private.PlayingSpace
	9,  // 5: playing_spaces.v1.private.PrivatePlayingSpacesService.ListPlayingSpaces:input_type -> playing_spaces.v1.private.ListPlayingSpacesRequest
	2,  // 6: playing_spaces.v1.private.PrivatePlayingSpacesService.CreatePlayingSpace:input_type -> playing_spaces.v1.private.CreatePlayingSpaceRequest
	4,  // 7: playing_spaces.v1.private.PrivatePlayingSpacesService.RemovePlayingSpaces:input_type -> playing_spaces.v1.private.RemovePlayingSpacesRequest
	6,  // 8: playing_spaces.v1.private.PrivatePlayingSpacesService.ModifyPlayingSpace:input_type -> playing_spaces.v1.private.ModifyPlayingSpaceRequest
	10, // 9: playing_spaces.v1.private.PrivatePlayingSpacesService.ListPlayingSpaces:output_type -> playing_spaces.v1.private.ListPlayingSpacesResponse
	3,  // 10: playing_spaces.v1.private.PrivatePlayingSpacesService.CreatePlayingSpace:output_type -> playing_spaces.v1.private.CreatePlayingSpaceResponse
	5,  // 11: playing_spaces.v1.private.PrivatePlayingSpacesService.RemovePlayingSpaces:output_type -> playing_spaces.v1.private.RemovePlayingSpacesResponse
	7,  // 12: playing_spaces.v1.private.PrivatePlayingSpacesService.ModifyPlayingSpace:output_type -> playing_spaces.v1.private.ModifyPlayingSpaceResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_playing_spaces_v1_private_playing_spaces_proto_init() }
func file_playing_spaces_v1_private_playing_spaces_proto_init() {
	if File_playing_spaces_v1_private_playing_spaces_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayingSpace); i {
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
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayingSpaceRequest); i {
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
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayingSpaceResponse); i {
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
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePlayingSpacesRequest); i {
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
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePlayingSpacesResponse); i {
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
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyPlayingSpaceRequest); i {
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
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyPlayingSpaceResponse); i {
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
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlayingSpacesFilter); i {
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
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlayingSpacesRequest); i {
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
		file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlayingSpacesResponse); i {
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
	file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_playing_spaces_v1_private_playing_spaces_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_playing_spaces_v1_private_playing_spaces_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_playing_spaces_v1_private_playing_spaces_proto_goTypes,
		DependencyIndexes: file_playing_spaces_v1_private_playing_spaces_proto_depIdxs,
		EnumInfos:         file_playing_spaces_v1_private_playing_spaces_proto_enumTypes,
		MessageInfos:      file_playing_spaces_v1_private_playing_spaces_proto_msgTypes,
	}.Build()
	File_playing_spaces_v1_private_playing_spaces_proto = out.File
	file_playing_spaces_v1_private_playing_spaces_proto_rawDesc = nil
	file_playing_spaces_v1_private_playing_spaces_proto_goTypes = nil
	file_playing_spaces_v1_private_playing_spaces_proto_depIdxs = nil
}
