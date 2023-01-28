// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: grades/v1/private/grades.proto

package gradesv1private

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

type Grade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GradeId string `protobuf:"bytes,3,opt,name=grade_id,json=gradeId,proto3" json:"grade_id,omitempty"`
}

func (x *Grade) Reset() {
	*x = Grade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_v1_private_grades_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grade) ProtoMessage() {}

func (x *Grade) ProtoReflect() protoreflect.Message {
	mi := &file_grades_v1_private_grades_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grade.ProtoReflect.Descriptor instead.
func (*Grade) Descriptor() ([]byte, []int) {
	return file_grades_v1_private_grades_proto_rawDescGZIP(), []int{0}
}

func (x *Grade) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Grade) GetGradeId() string {
	if x != nil {
		return x.GradeId
	}
	return ""
}

type CreateGradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateGradeRequest) Reset() {
	*x = CreateGradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_v1_private_grades_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGradeRequest) ProtoMessage() {}

func (x *CreateGradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grades_v1_private_grades_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGradeRequest.ProtoReflect.Descriptor instead.
func (*CreateGradeRequest) Descriptor() ([]byte, []int) {
	return file_grades_v1_private_grades_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGradeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateGradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grade *Grade `protobuf:"bytes,1,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *CreateGradeResponse) Reset() {
	*x = CreateGradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_v1_private_grades_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGradeResponse) ProtoMessage() {}

func (x *CreateGradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grades_v1_private_grades_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGradeResponse.ProtoReflect.Descriptor instead.
func (*CreateGradeResponse) Descriptor() ([]byte, []int) {
	return file_grades_v1_private_grades_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGradeResponse) GetGrade() *Grade {
	if x != nil {
		return x.Grade
	}
	return nil
}

type ListGradesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGradesRequest) Reset() {
	*x = ListGradesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_v1_private_grades_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGradesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGradesRequest) ProtoMessage() {}

func (x *ListGradesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grades_v1_private_grades_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGradesRequest.ProtoReflect.Descriptor instead.
func (*ListGradesRequest) Descriptor() ([]byte, []int) {
	return file_grades_v1_private_grades_proto_rawDescGZIP(), []int{3}
}

type ListGradesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grades []*Grade `protobuf:"bytes,1,rep,name=grades,proto3" json:"grades,omitempty"`
}

func (x *ListGradesResponse) Reset() {
	*x = ListGradesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grades_v1_private_grades_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGradesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGradesResponse) ProtoMessage() {}

func (x *ListGradesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grades_v1_private_grades_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGradesResponse.ProtoReflect.Descriptor instead.
func (*ListGradesResponse) Descriptor() ([]byte, []int) {
	return file_grades_v1_private_grades_proto_rawDescGZIP(), []int{4}
}

func (x *ListGradesResponse) GetGrades() []*Grade {
	if x != nil {
		return x.Grades
	}
	return nil
}

var File_grades_v1_private_grades_proto protoreflect.FileDescriptor

var file_grades_v1_private_grades_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x13, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x46, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x32, 0xd3, 0x01, 0x0a, 0x14, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x25, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x12, 0x24, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65,
	0x65, 0x67, 0x61, 0x6e, 0x63, 0x72, 0x61, 0x69, 0x67, 0x63, 0x72, 0x75, 0x69, 0x63, 0x6b, 0x73,
	0x68, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x6f, 0x6c, 0x6c, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x3b, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x76, 0x31, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_grades_v1_private_grades_proto_rawDescOnce sync.Once
	file_grades_v1_private_grades_proto_rawDescData = file_grades_v1_private_grades_proto_rawDesc
)

func file_grades_v1_private_grades_proto_rawDescGZIP() []byte {
	file_grades_v1_private_grades_proto_rawDescOnce.Do(func() {
		file_grades_v1_private_grades_proto_rawDescData = protoimpl.X.CompressGZIP(file_grades_v1_private_grades_proto_rawDescData)
	})
	return file_grades_v1_private_grades_proto_rawDescData
}

var file_grades_v1_private_grades_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grades_v1_private_grades_proto_goTypes = []interface{}{
	(*Grade)(nil),               // 0: grades.v1.private.Grade
	(*CreateGradeRequest)(nil),  // 1: grades.v1.private.CreateGradeRequest
	(*CreateGradeResponse)(nil), // 2: grades.v1.private.CreateGradeResponse
	(*ListGradesRequest)(nil),   // 3: grades.v1.private.ListGradesRequest
	(*ListGradesResponse)(nil),  // 4: grades.v1.private.ListGradesResponse
}
var file_grades_v1_private_grades_proto_depIdxs = []int32{
	0, // 0: grades.v1.private.CreateGradeResponse.grade:type_name -> grades.v1.private.Grade
	0, // 1: grades.v1.private.ListGradesResponse.grades:type_name -> grades.v1.private.Grade
	1, // 2: grades.v1.private.PrivateGradesService.CreateGrade:input_type -> grades.v1.private.CreateGradeRequest
	3, // 3: grades.v1.private.PrivateGradesService.ListGrades:input_type -> grades.v1.private.ListGradesRequest
	2, // 4: grades.v1.private.PrivateGradesService.CreateGrade:output_type -> grades.v1.private.CreateGradeResponse
	4, // 5: grades.v1.private.PrivateGradesService.ListGrades:output_type -> grades.v1.private.ListGradesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grades_v1_private_grades_proto_init() }
func file_grades_v1_private_grades_proto_init() {
	if File_grades_v1_private_grades_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grades_v1_private_grades_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grade); i {
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
		file_grades_v1_private_grades_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGradeRequest); i {
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
		file_grades_v1_private_grades_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGradeResponse); i {
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
		file_grades_v1_private_grades_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGradesRequest); i {
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
		file_grades_v1_private_grades_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGradesResponse); i {
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
			RawDescriptor: file_grades_v1_private_grades_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grades_v1_private_grades_proto_goTypes,
		DependencyIndexes: file_grades_v1_private_grades_proto_depIdxs,
		MessageInfos:      file_grades_v1_private_grades_proto_msgTypes,
	}.Build()
	File_grades_v1_private_grades_proto = out.File
	file_grades_v1_private_grades_proto_rawDesc = nil
	file_grades_v1_private_grades_proto_goTypes = nil
	file_grades_v1_private_grades_proto_depIdxs = nil
}
