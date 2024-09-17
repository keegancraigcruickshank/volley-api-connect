// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: remote-scoring/v1/private/remote-scoring.proto

package remotescoringv1private

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UploadScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamOneId    string                 `protobuf:"bytes,1,opt,name=team_one_id,json=teamOneId,proto3" json:"team_one_id,omitempty"`
	TeamTwoId    string                 `protobuf:"bytes,2,opt,name=team_two_id,json=teamTwoId,proto3" json:"team_two_id,omitempty"`
	TeamOneScore int32                  `protobuf:"varint,3,opt,name=team_one_score,json=teamOneScore,proto3" json:"team_one_score,omitempty"`
	TeamTwoScore int32                  `protobuf:"varint,4,opt,name=team_two_score,json=teamTwoScore,proto3" json:"team_two_score,omitempty"`
	PlayedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=played_at,json=playedAt,proto3" json:"played_at,omitempty"`
}

func (x *UploadScoreRequest) Reset() {
	*x = UploadScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadScoreRequest) ProtoMessage() {}

func (x *UploadScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadScoreRequest.ProtoReflect.Descriptor instead.
func (*UploadScoreRequest) Descriptor() ([]byte, []int) {
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP(), []int{0}
}

func (x *UploadScoreRequest) GetTeamOneId() string {
	if x != nil {
		return x.TeamOneId
	}
	return ""
}

func (x *UploadScoreRequest) GetTeamTwoId() string {
	if x != nil {
		return x.TeamTwoId
	}
	return ""
}

func (x *UploadScoreRequest) GetTeamOneScore() int32 {
	if x != nil {
		return x.TeamOneScore
	}
	return 0
}

func (x *UploadScoreRequest) GetTeamTwoScore() int32 {
	if x != nil {
		return x.TeamTwoScore
	}
	return 0
}

func (x *UploadScoreRequest) GetPlayedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PlayedAt
	}
	return nil
}

type UploadScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadScoreResponse) Reset() {
	*x = UploadScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadScoreResponse) ProtoMessage() {}

func (x *UploadScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadScoreResponse.ProtoReflect.Descriptor instead.
func (*UploadScoreResponse) Descriptor() ([]byte, []int) {
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP(), []int{1}
}

type ListScoresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ListScoresRequest) Reset() {
	*x = ListScoresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScoresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScoresRequest) ProtoMessage() {}

func (x *ListScoresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScoresRequest.ProtoReflect.Descriptor instead.
func (*ListScoresRequest) Descriptor() ([]byte, []int) {
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP(), []int{2}
}

func (x *ListScoresRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type ListScoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores []*ListScoresResponse_Score `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
}

func (x *ListScoresResponse) Reset() {
	*x = ListScoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScoresResponse) ProtoMessage() {}

func (x *ListScoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScoresResponse.ProtoReflect.Descriptor instead.
func (*ListScoresResponse) Descriptor() ([]byte, []int) {
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP(), []int{3}
}

func (x *ListScoresResponse) GetScores() []*ListScoresResponse_Score {
	if x != nil {
		return x.Scores
	}
	return nil
}

type DeleteScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScoreId string `protobuf:"bytes,1,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
}

func (x *DeleteScoreRequest) Reset() {
	*x = DeleteScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScoreRequest) ProtoMessage() {}

func (x *DeleteScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScoreRequest.ProtoReflect.Descriptor instead.
func (*DeleteScoreRequest) Descriptor() ([]byte, []int) {
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteScoreRequest) GetScoreId() string {
	if x != nil {
		return x.ScoreId
	}
	return ""
}

type DeleteScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteScoreResponse) Reset() {
	*x = DeleteScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScoreResponse) ProtoMessage() {}

func (x *DeleteScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScoreResponse.ProtoReflect.Descriptor instead.
func (*DeleteScoreResponse) Descriptor() ([]byte, []int) {
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP(), []int{5}
}

type UpdateScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScoreId      string `protobuf:"bytes,1,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
	TeamOneId    string `protobuf:"bytes,2,opt,name=team_one_id,json=teamOneId,proto3" json:"team_one_id,omitempty"`
	TeamTwoId    string `protobuf:"bytes,3,opt,name=team_two_id,json=teamTwoId,proto3" json:"team_two_id,omitempty"`
	TeamOneScore int32  `protobuf:"varint,4,opt,name=team_one_score,json=teamOneScore,proto3" json:"team_one_score,omitempty"`
	TeamTwoScore int32  `protobuf:"varint,5,opt,name=team_two_score,json=teamTwoScore,proto3" json:"team_two_score,omitempty"`
}

func (x *UpdateScoreRequest) Reset() {
	*x = UpdateScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScoreRequest) ProtoMessage() {}

func (x *UpdateScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScoreRequest.ProtoReflect.Descriptor instead.
func (*UpdateScoreRequest) Descriptor() ([]byte, []int) {
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateScoreRequest) GetScoreId() string {
	if x != nil {
		return x.ScoreId
	}
	return ""
}

func (x *UpdateScoreRequest) GetTeamOneId() string {
	if x != nil {
		return x.TeamOneId
	}
	return ""
}

func (x *UpdateScoreRequest) GetTeamTwoId() string {
	if x != nil {
		return x.TeamTwoId
	}
	return ""
}

func (x *UpdateScoreRequest) GetTeamOneScore() int32 {
	if x != nil {
		return x.TeamOneScore
	}
	return 0
}

func (x *UpdateScoreRequest) GetTeamTwoScore() int32 {
	if x != nil {
		return x.TeamTwoScore
	}
	return 0
}

type UpdateScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateScoreResponse) Reset() {
	*x = UpdateScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScoreResponse) ProtoMessage() {}

func (x *UpdateScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScoreResponse.ProtoReflect.Descriptor instead.
func (*UpdateScoreResponse) Descriptor() ([]byte, []int) {
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP(), []int{7}
}

type ListScoresResponse_Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScoreId      string                 `protobuf:"bytes,1,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
	TeamOneId    string                 `protobuf:"bytes,2,opt,name=team_one_id,json=teamOneId,proto3" json:"team_one_id,omitempty"`
	TeamTwoId    string                 `protobuf:"bytes,3,opt,name=team_two_id,json=teamTwoId,proto3" json:"team_two_id,omitempty"`
	TeamOneScore int32                  `protobuf:"varint,4,opt,name=team_one_score,json=teamOneScore,proto3" json:"team_one_score,omitempty"`
	TeamTwoScore int32                  `protobuf:"varint,5,opt,name=team_two_score,json=teamTwoScore,proto3" json:"team_two_score,omitempty"`
	PlayedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=played_at,json=playedAt,proto3" json:"played_at,omitempty"`
}

func (x *ListScoresResponse_Score) Reset() {
	*x = ListScoresResponse_Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScoresResponse_Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScoresResponse_Score) ProtoMessage() {}

func (x *ListScoresResponse_Score) ProtoReflect() protoreflect.Message {
	mi := &file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScoresResponse_Score.ProtoReflect.Descriptor instead.
func (*ListScoresResponse_Score) Descriptor() ([]byte, []int) {
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListScoresResponse_Score) GetScoreId() string {
	if x != nil {
		return x.ScoreId
	}
	return ""
}

func (x *ListScoresResponse_Score) GetTeamOneId() string {
	if x != nil {
		return x.TeamOneId
	}
	return ""
}

func (x *ListScoresResponse_Score) GetTeamTwoId() string {
	if x != nil {
		return x.TeamTwoId
	}
	return ""
}

func (x *ListScoresResponse_Score) GetTeamOneScore() int32 {
	if x != nil {
		return x.TeamOneScore
	}
	return 0
}

func (x *ListScoresResponse_Score) GetTeamTwoScore() int32 {
	if x != nil {
		return x.TeamTwoScore
	}
	return 0
}

func (x *ListScoresResponse_Score) GetPlayedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PlayedAt
	}
	return nil
}

var File_remote_scoring_v1_private_remote_scoring_proto protoreflect.FileDescriptor

var file_remote_scoring_v1_private_remote_scoring_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2d, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a,
	0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x4f, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x77, 0x6f, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x54, 0x77,
	0x6f, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x6e, 0x65, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x65, 0x61,
	0x6d, 0x4f, 0x6e, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x74, 0x65, 0x61, 0x6d, 0x54, 0x77, 0x6f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x37, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x41, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x43, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0xcb, 0x02, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x1a, 0xe7, 0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0b, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x4f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0b, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x54, 0x77, 0x6f, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x65, 0x61, 0x6d, 0x4f, 0x6e, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x77, 0x6f, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x65, 0x61,
	0x6d, 0x54, 0x77, 0x6f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x4f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x54, 0x77, 0x6f, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x65, 0x61, 0x6d, 0x4f, 0x6e, 0x65, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x65, 0x61, 0x6d,
	0x54, 0x77, 0x6f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xd2, 0x03, 0x0a, 0x1b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6c, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2d,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73,
	0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x63,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x67, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x65, 0x67, 0x61, 0x6e, 0x63, 0x72, 0x61, 0x69, 0x67, 0x63, 0x72,
	0x75, 0x69, 0x63, 0x6b, 0x73, 0x68, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x6f, 0x6c, 0x6c, 0x65, 0x79,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2d, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x3b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x63, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_scoring_v1_private_remote_scoring_proto_rawDescOnce sync.Once
	file_remote_scoring_v1_private_remote_scoring_proto_rawDescData = file_remote_scoring_v1_private_remote_scoring_proto_rawDesc
)

func file_remote_scoring_v1_private_remote_scoring_proto_rawDescGZIP() []byte {
	file_remote_scoring_v1_private_remote_scoring_proto_rawDescOnce.Do(func() {
		file_remote_scoring_v1_private_remote_scoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_scoring_v1_private_remote_scoring_proto_rawDescData)
	})
	return file_remote_scoring_v1_private_remote_scoring_proto_rawDescData
}

var file_remote_scoring_v1_private_remote_scoring_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_remote_scoring_v1_private_remote_scoring_proto_goTypes = []interface{}{
	(*UploadScoreRequest)(nil),       // 0: remote_scoring.v1.private.UploadScoreRequest
	(*UploadScoreResponse)(nil),      // 1: remote_scoring.v1.private.UploadScoreResponse
	(*ListScoresRequest)(nil),        // 2: remote_scoring.v1.private.ListScoresRequest
	(*ListScoresResponse)(nil),       // 3: remote_scoring.v1.private.ListScoresResponse
	(*DeleteScoreRequest)(nil),       // 4: remote_scoring.v1.private.DeleteScoreRequest
	(*DeleteScoreResponse)(nil),      // 5: remote_scoring.v1.private.DeleteScoreResponse
	(*UpdateScoreRequest)(nil),       // 6: remote_scoring.v1.private.UpdateScoreRequest
	(*UpdateScoreResponse)(nil),      // 7: remote_scoring.v1.private.UpdateScoreResponse
	(*ListScoresResponse_Score)(nil), // 8: remote_scoring.v1.private.ListScoresResponse.Score
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
}
var file_remote_scoring_v1_private_remote_scoring_proto_depIdxs = []int32{
	9, // 0: remote_scoring.v1.private.UploadScoreRequest.played_at:type_name -> google.protobuf.Timestamp
	9, // 1: remote_scoring.v1.private.ListScoresRequest.date:type_name -> google.protobuf.Timestamp
	8, // 2: remote_scoring.v1.private.ListScoresResponse.scores:type_name -> remote_scoring.v1.private.ListScoresResponse.Score
	9, // 3: remote_scoring.v1.private.ListScoresResponse.Score.played_at:type_name -> google.protobuf.Timestamp
	0, // 4: remote_scoring.v1.private.PrivateRemoteScoringService.UploadScore:input_type -> remote_scoring.v1.private.UploadScoreRequest
	2, // 5: remote_scoring.v1.private.PrivateRemoteScoringService.ListScores:input_type -> remote_scoring.v1.private.ListScoresRequest
	4, // 6: remote_scoring.v1.private.PrivateRemoteScoringService.DeleteScore:input_type -> remote_scoring.v1.private.DeleteScoreRequest
	6, // 7: remote_scoring.v1.private.PrivateRemoteScoringService.UpdateScore:input_type -> remote_scoring.v1.private.UpdateScoreRequest
	1, // 8: remote_scoring.v1.private.PrivateRemoteScoringService.UploadScore:output_type -> remote_scoring.v1.private.UploadScoreResponse
	3, // 9: remote_scoring.v1.private.PrivateRemoteScoringService.ListScores:output_type -> remote_scoring.v1.private.ListScoresResponse
	5, // 10: remote_scoring.v1.private.PrivateRemoteScoringService.DeleteScore:output_type -> remote_scoring.v1.private.DeleteScoreResponse
	7, // 11: remote_scoring.v1.private.PrivateRemoteScoringService.UpdateScore:output_type -> remote_scoring.v1.private.UpdateScoreResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_remote_scoring_v1_private_remote_scoring_proto_init() }
func file_remote_scoring_v1_private_remote_scoring_proto_init() {
	if File_remote_scoring_v1_private_remote_scoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadScoreRequest); i {
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
		file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadScoreResponse); i {
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
		file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListScoresRequest); i {
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
		file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListScoresResponse); i {
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
		file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScoreRequest); i {
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
		file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScoreResponse); i {
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
		file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScoreRequest); i {
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
		file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScoreResponse); i {
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
		file_remote_scoring_v1_private_remote_scoring_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListScoresResponse_Score); i {
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
			RawDescriptor: file_remote_scoring_v1_private_remote_scoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remote_scoring_v1_private_remote_scoring_proto_goTypes,
		DependencyIndexes: file_remote_scoring_v1_private_remote_scoring_proto_depIdxs,
		MessageInfos:      file_remote_scoring_v1_private_remote_scoring_proto_msgTypes,
	}.Build()
	File_remote_scoring_v1_private_remote_scoring_proto = out.File
	file_remote_scoring_v1_private_remote_scoring_proto_rawDesc = nil
	file_remote_scoring_v1_private_remote_scoring_proto_goTypes = nil
	file_remote_scoring_v1_private_remote_scoring_proto_depIdxs = nil
}
