// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: users/v1/users.proto

package usersv1

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

type GetMeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMeRequest) Reset() {
	*x = GetMeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_v1_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeRequest) ProtoMessage() {}

func (x *GetMeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_v1_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeRequest.ProtoReflect.Descriptor instead.
func (*GetMeRequest) Descriptor() ([]byte, []int) {
	return file_users_v1_users_proto_rawDescGZIP(), []int{0}
}

type GetMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrgId     string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Nickname  string `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *GetMeResponse) Reset() {
	*x = GetMeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_v1_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeResponse) ProtoMessage() {}

func (x *GetMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_v1_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeResponse.ProtoReflect.Descriptor instead.
func (*GetMeResponse) Descriptor() ([]byte, []int) {
	return file_users_v1_users_proto_rawDescGZIP(), []int{1}
}

func (x *GetMeResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetMeResponse) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *GetMeResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetMeResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetMeResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetMeResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type CreateOrganisationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOrganisationRequest) Reset() {
	*x = CreateOrganisationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_v1_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrganisationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrganisationRequest) ProtoMessage() {}

func (x *CreateOrganisationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_v1_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrganisationRequest.ProtoReflect.Descriptor instead.
func (*CreateOrganisationRequest) Descriptor() ([]byte, []int) {
	return file_users_v1_users_proto_rawDescGZIP(), []int{2}
}

type CreateOrganisationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId    string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OrgId      string   `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	OrgName    string   `protobuf:"bytes,3,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	OrgMembers []string `protobuf:"bytes,4,rep,name=org_members,json=orgMembers,proto3" json:"org_members,omitempty"`
}

func (x *CreateOrganisationResponse) Reset() {
	*x = CreateOrganisationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_v1_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrganisationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrganisationResponse) ProtoMessage() {}

func (x *CreateOrganisationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_v1_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrganisationResponse.ProtoReflect.Descriptor instead.
func (*CreateOrganisationResponse) Descriptor() ([]byte, []int) {
	return file_users_v1_users_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrganisationResponse) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *CreateOrganisationResponse) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CreateOrganisationResponse) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *CreateOrganisationResponse) GetOrgMembers() []string {
	if x != nil {
		return x.OrgMembers
	}
	return nil
}

var File_users_v1_users_proto protoreflect.FileDescriptor

var file_users_v1_users_proto_rawDesc = []byte{
	0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xad, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f,
	0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x1b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8a, 0x01,
	0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x67,
	0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x6f, 0x72, 0x67, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x32, 0xad, 0x01, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x65, 0x67, 0x61, 0x6e, 0x63,
	0x72, 0x61, 0x69, 0x67, 0x63, 0x72, 0x75, 0x69, 0x63, 0x6b, 0x73, 0x68, 0x61, 0x6e, 0x6b, 0x2f,
	0x76, 0x6f, 0x6c, 0x6c, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_v1_users_proto_rawDescOnce sync.Once
	file_users_v1_users_proto_rawDescData = file_users_v1_users_proto_rawDesc
)

func file_users_v1_users_proto_rawDescGZIP() []byte {
	file_users_v1_users_proto_rawDescOnce.Do(func() {
		file_users_v1_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_v1_users_proto_rawDescData)
	})
	return file_users_v1_users_proto_rawDescData
}

var file_users_v1_users_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_users_v1_users_proto_goTypes = []interface{}{
	(*GetMeRequest)(nil),               // 0: users.v1.GetMeRequest
	(*GetMeResponse)(nil),              // 1: users.v1.GetMeResponse
	(*CreateOrganisationRequest)(nil),  // 2: users.v1.CreateOrganisationRequest
	(*CreateOrganisationResponse)(nil), // 3: users.v1.CreateOrganisationResponse
}
var file_users_v1_users_proto_depIdxs = []int32{
	0, // 0: users.v1.UsersService.GetMe:input_type -> users.v1.GetMeRequest
	2, // 1: users.v1.UsersService.CreateOrganisation:input_type -> users.v1.CreateOrganisationRequest
	1, // 2: users.v1.UsersService.GetMe:output_type -> users.v1.GetMeResponse
	3, // 3: users.v1.UsersService.CreateOrganisation:output_type -> users.v1.CreateOrganisationResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_users_v1_users_proto_init() }
func file_users_v1_users_proto_init() {
	if File_users_v1_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_v1_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeRequest); i {
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
		file_users_v1_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeResponse); i {
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
		file_users_v1_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrganisationRequest); i {
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
		file_users_v1_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrganisationResponse); i {
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
			RawDescriptor: file_users_v1_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_v1_users_proto_goTypes,
		DependencyIndexes: file_users_v1_users_proto_depIdxs,
		MessageInfos:      file_users_v1_users_proto_msgTypes,
	}.Build()
	File_users_v1_users_proto = out.File
	file_users_v1_users_proto_rawDesc = nil
	file_users_v1_users_proto_goTypes = nil
	file_users_v1_users_proto_depIdxs = nil
}
