// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: teams/v1/private/teams.proto

package teamsv1private

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
	return file_teams_v1_private_teams_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_teams_v1_private_teams_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{0}
}

type ListedTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Players  []string `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
	Modified int64    `protobuf:"varint,4,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *ListedTeam) Reset() {
	*x = ListedTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListedTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListedTeam) ProtoMessage() {}

func (x *ListedTeam) ProtoReflect() protoreflect.Message {
	mi := &file_teams_v1_private_teams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListedTeam.ProtoReflect.Descriptor instead.
func (*ListedTeam) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{0}
}

func (x *ListedTeam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListedTeam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListedTeam) GetPlayers() []string {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *ListedTeam) GetModified() int64 {
	if x != nil {
		return x.Modified
	}
	return 0
}

type AddTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Players []string `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *AddTeamRequest) Reset() {
	*x = AddTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTeamRequest) ProtoMessage() {}

func (x *AddTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teams_v1_private_teams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTeamRequest.ProtoReflect.Descriptor instead.
func (*AddTeamRequest) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{1}
}

func (x *AddTeamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddTeamRequest) GetPlayers() []string {
	if x != nil {
		return x.Players
	}
	return nil
}

type AddTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Players string `protobuf:"bytes,3,opt,name=players,proto3" json:"players,omitempty"`
}

func (x *AddTeamResponse) Reset() {
	*x = AddTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTeamResponse) ProtoMessage() {}

func (x *AddTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teams_v1_private_teams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTeamResponse.ProtoReflect.Descriptor instead.
func (*AddTeamResponse) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{2}
}

func (x *AddTeamResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddTeamResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddTeamResponse) GetPlayers() string {
	if x != nil {
		return x.Players
	}
	return ""
}

type ListTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         string         `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Page          *int32         `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
	PageSize      *int32         `protobuf:"varint,3,opt,name=pageSize,proto3,oneof" json:"pageSize,omitempty"`
	SortDirection *SortDirection `protobuf:"varint,4,opt,name=sortDirection,proto3,enum=teams.v1.private.SortDirection,oneof" json:"sortDirection,omitempty"`
}

func (x *ListTeamsRequest) Reset() {
	*x = ListTeamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsRequest) ProtoMessage() {}

func (x *ListTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teams_v1_private_teams_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsRequest.ProtoReflect.Descriptor instead.
func (*ListTeamsRequest) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{3}
}

func (x *ListTeamsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListTeamsRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListTeamsRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListTeamsRequest) GetSortDirection() SortDirection {
	if x != nil && x.SortDirection != nil {
		return *x.SortDirection
	}
	return SortDirection_ASC
}

type ListTeamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams      []*ListedTeam `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	TotalCount int32         `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListTeamsResponse) Reset() {
	*x = ListTeamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsResponse) ProtoMessage() {}

func (x *ListTeamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teams_v1_private_teams_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsResponse.ProtoReflect.Descriptor instead.
func (*ListTeamsResponse) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{4}
}

func (x *ListTeamsResponse) GetTeams() []*ListedTeam {
	if x != nil {
		return x.Teams
	}
	return nil
}

func (x *ListTeamsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type RemoveTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []string `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *RemoveTeamsRequest) Reset() {
	*x = RemoveTeamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTeamsRequest) ProtoMessage() {}

func (x *RemoveTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teams_v1_private_teams_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTeamsRequest.ProtoReflect.Descriptor instead.
func (*RemoveTeamsRequest) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveTeamsRequest) GetTeams() []string {
	if x != nil {
		return x.Teams
	}
	return nil
}

type RemoveTeamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveTeamsResponse) Reset() {
	*x = RemoveTeamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTeamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTeamsResponse) ProtoMessage() {}

func (x *RemoveTeamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teams_v1_private_teams_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTeamsResponse.ProtoReflect.Descriptor instead.
func (*RemoveTeamsResponse) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{6}
}

type UpdateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    *string  `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Players []string `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *UpdateTeamRequest) Reset() {
	*x = UpdateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeamRequest) ProtoMessage() {}

func (x *UpdateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teams_v1_private_teams_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeamRequest.ProtoReflect.Descriptor instead.
func (*UpdateTeamRequest) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTeamRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTeamRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateTeamRequest) GetPlayers() []string {
	if x != nil {
		return x.Players
	}
	return nil
}

type UpdateTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Players []string `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *UpdateTeamResponse) Reset() {
	*x = UpdateTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeamResponse) ProtoMessage() {}

func (x *UpdateTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teams_v1_private_teams_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeamResponse.ProtoReflect.Descriptor instead.
func (*UpdateTeamResponse) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTeamResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTeamResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTeamResponse) GetPlayers() []string {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_teams_v1_private_teams_proto protoreflect.FileDescriptor

var file_teams_v1_private_teams_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x22, 0x66, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x4f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4a,
	0x0a, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x67, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x12, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x52, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x2a, 0x22, 0x0a, 0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x32, 0xf6, 0x02, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x50, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x61,
	0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x54, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x20,
	0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x65, 0x65, 0x67, 0x61, 0x6e, 0x63, 0x72, 0x61, 0x69, 0x67, 0x63, 0x72, 0x75, 0x69, 0x63, 0x6b,
	0x73, 0x68, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x6f, 0x6c, 0x6c, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x3b, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x76,
	0x31, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teams_v1_private_teams_proto_rawDescOnce sync.Once
	file_teams_v1_private_teams_proto_rawDescData = file_teams_v1_private_teams_proto_rawDesc
)

func file_teams_v1_private_teams_proto_rawDescGZIP() []byte {
	file_teams_v1_private_teams_proto_rawDescOnce.Do(func() {
		file_teams_v1_private_teams_proto_rawDescData = protoimpl.X.CompressGZIP(file_teams_v1_private_teams_proto_rawDescData)
	})
	return file_teams_v1_private_teams_proto_rawDescData
}

var file_teams_v1_private_teams_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teams_v1_private_teams_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_teams_v1_private_teams_proto_goTypes = []interface{}{
	(SortDirection)(0),          // 0: teams.v1.private.SortDirection
	(*ListedTeam)(nil),          // 1: teams.v1.private.ListedTeam
	(*AddTeamRequest)(nil),      // 2: teams.v1.private.AddTeamRequest
	(*AddTeamResponse)(nil),     // 3: teams.v1.private.AddTeamResponse
	(*ListTeamsRequest)(nil),    // 4: teams.v1.private.ListTeamsRequest
	(*ListTeamsResponse)(nil),   // 5: teams.v1.private.ListTeamsResponse
	(*RemoveTeamsRequest)(nil),  // 6: teams.v1.private.RemoveTeamsRequest
	(*RemoveTeamsResponse)(nil), // 7: teams.v1.private.RemoveTeamsResponse
	(*UpdateTeamRequest)(nil),   // 8: teams.v1.private.UpdateTeamRequest
	(*UpdateTeamResponse)(nil),  // 9: teams.v1.private.UpdateTeamResponse
}
var file_teams_v1_private_teams_proto_depIdxs = []int32{
	0, // 0: teams.v1.private.ListTeamsRequest.sortDirection:type_name -> teams.v1.private.SortDirection
	1, // 1: teams.v1.private.ListTeamsResponse.teams:type_name -> teams.v1.private.ListedTeam
	2, // 2: teams.v1.private.PrivateTeamsService.AddTeam:input_type -> teams.v1.private.AddTeamRequest
	2, // 3: teams.v1.private.PrivateTeamsService.ListTeams:input_type -> teams.v1.private.AddTeamRequest
	6, // 4: teams.v1.private.PrivateTeamsService.RemoveTeams:input_type -> teams.v1.private.RemoveTeamsRequest
	8, // 5: teams.v1.private.PrivateTeamsService.UpdateTeam:input_type -> teams.v1.private.UpdateTeamRequest
	3, // 6: teams.v1.private.PrivateTeamsService.AddTeam:output_type -> teams.v1.private.AddTeamResponse
	5, // 7: teams.v1.private.PrivateTeamsService.ListTeams:output_type -> teams.v1.private.ListTeamsResponse
	7, // 8: teams.v1.private.PrivateTeamsService.RemoveTeams:output_type -> teams.v1.private.RemoveTeamsResponse
	9, // 9: teams.v1.private.PrivateTeamsService.UpdateTeam:output_type -> teams.v1.private.UpdateTeamResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_teams_v1_private_teams_proto_init() }
func file_teams_v1_private_teams_proto_init() {
	if File_teams_v1_private_teams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teams_v1_private_teams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListedTeam); i {
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
		file_teams_v1_private_teams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTeamRequest); i {
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
		file_teams_v1_private_teams_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTeamResponse); i {
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
		file_teams_v1_private_teams_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeamsRequest); i {
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
		file_teams_v1_private_teams_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeamsResponse); i {
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
		file_teams_v1_private_teams_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTeamsRequest); i {
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
		file_teams_v1_private_teams_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTeamsResponse); i {
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
		file_teams_v1_private_teams_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTeamRequest); i {
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
		file_teams_v1_private_teams_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTeamResponse); i {
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
	file_teams_v1_private_teams_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_teams_v1_private_teams_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teams_v1_private_teams_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teams_v1_private_teams_proto_goTypes,
		DependencyIndexes: file_teams_v1_private_teams_proto_depIdxs,
		EnumInfos:         file_teams_v1_private_teams_proto_enumTypes,
		MessageInfos:      file_teams_v1_private_teams_proto_msgTypes,
	}.Build()
	File_teams_v1_private_teams_proto = out.File
	file_teams_v1_private_teams_proto_rawDesc = nil
	file_teams_v1_private_teams_proto_goTypes = nil
	file_teams_v1_private_teams_proto_depIdxs = nil
}
