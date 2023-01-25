// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
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

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Organisation string `protobuf:"bytes,2,opt,name=organisation,proto3" json:"organisation,omitempty"`
	Id           string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Day          string `protobuf:"bytes,4,opt,name=day,proto3" json:"day,omitempty"`
	GradeId      string `protobuf:"bytes,5,opt,name=grade_id,json=gradeId,proto3" json:"grade_id,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetOrganisation() string {
	if x != nil {
		return x.Organisation
	}
	return ""
}

func (x *Team) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Team) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *Team) GetGradeId() string {
	if x != nil {
		return x.GradeId
	}
	return ""
}

type CreateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Day     string `protobuf:"bytes,2,opt,name=day,proto3" json:"day,omitempty"`
	GradeId string `protobuf:"bytes,3,opt,name=grade_id,json=gradeId,proto3" json:"grade_id,omitempty"`
}

func (x *CreateTeamRequest) Reset() {
	*x = CreateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamRequest) ProtoMessage() {}

func (x *CreateTeamRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateTeamRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTeamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTeamRequest) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *CreateTeamRequest) GetGradeId() string {
	if x != nil {
		return x.GradeId
	}
	return ""
}

type CreateTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team *Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *CreateTeamResponse) Reset() {
	*x = CreateTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teams_v1_private_teams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamResponse) ProtoMessage() {}

func (x *CreateTeamResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateTeamResponse.ProtoReflect.Descriptor instead.
func (*CreateTeamResponse) Descriptor() ([]byte, []int) {
	return file_teams_v1_private_teams_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTeamResponse) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

type ListTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
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

type ListTeamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []*Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
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

func (x *ListTeamsResponse) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

var File_teams_v1_private_teams_proto protoreflect.FileDescriptor

var file_teams_v1_private_teams_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x22, 0x7b, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x61, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x22, 0x54, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x32, 0xc8, 0x01, 0x0a,
	0x13, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x12, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x56, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x22, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x65, 0x67, 0x61, 0x6e, 0x63, 0x72, 0x61, 0x69,
	0x67, 0x63, 0x72, 0x75, 0x69, 0x63, 0x6b, 0x73, 0x68, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x6f, 0x6c,
	0x6c, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x3b, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x76, 0x31, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_teams_v1_private_teams_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_teams_v1_private_teams_proto_goTypes = []interface{}{
	(*Team)(nil),               // 0: teams.v1.private.Team
	(*CreateTeamRequest)(nil),  // 1: teams.v1.private.CreateTeamRequest
	(*CreateTeamResponse)(nil), // 2: teams.v1.private.CreateTeamResponse
	(*ListTeamsRequest)(nil),   // 3: teams.v1.private.ListTeamsRequest
	(*ListTeamsResponse)(nil),  // 4: teams.v1.private.ListTeamsResponse
}
var file_teams_v1_private_teams_proto_depIdxs = []int32{
	0, // 0: teams.v1.private.CreateTeamResponse.team:type_name -> teams.v1.private.Team
	0, // 1: teams.v1.private.ListTeamsResponse.teams:type_name -> teams.v1.private.Team
	1, // 2: teams.v1.private.PrivateTeamsService.CreateTeam:input_type -> teams.v1.private.CreateTeamRequest
	3, // 3: teams.v1.private.PrivateTeamsService.ListTeams:input_type -> teams.v1.private.ListTeamsRequest
	2, // 4: teams.v1.private.PrivateTeamsService.CreateTeam:output_type -> teams.v1.private.CreateTeamResponse
	4, // 5: teams.v1.private.PrivateTeamsService.ListTeams:output_type -> teams.v1.private.ListTeamsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
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
			switch v := v.(*Team); i {
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
			switch v := v.(*CreateTeamRequest); i {
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
			switch v := v.(*CreateTeamResponse); i {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teams_v1_private_teams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teams_v1_private_teams_proto_goTypes,
		DependencyIndexes: file_teams_v1_private_teams_proto_depIdxs,
		MessageInfos:      file_teams_v1_private_teams_proto_msgTypes,
	}.Build()
	File_teams_v1_private_teams_proto = out.File
	file_teams_v1_private_teams_proto_rawDesc = nil
	file_teams_v1_private_teams_proto_goTypes = nil
	file_teams_v1_private_teams_proto_depIdxs = nil
}
