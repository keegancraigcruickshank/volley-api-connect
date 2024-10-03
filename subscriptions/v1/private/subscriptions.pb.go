// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: subscriptions/v1/private/subscriptions.proto

package subscriptionsv1private

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

type SubscriptionPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubscriptionKey string  `protobuf:"bytes,2,opt,name=subscription_key,json=subscriptionKey,proto3" json:"subscription_key,omitempty"`
	Name            string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description     string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price           float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SubscriptionPlan) Reset() {
	*x = SubscriptionPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_v1_private_subscriptions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionPlan) ProtoMessage() {}

func (x *SubscriptionPlan) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_v1_private_subscriptions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionPlan.ProtoReflect.Descriptor instead.
func (*SubscriptionPlan) Descriptor() ([]byte, []int) {
	return file_subscriptions_v1_private_subscriptions_proto_rawDescGZIP(), []int{0}
}

func (x *SubscriptionPlan) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubscriptionPlan) GetSubscriptionKey() string {
	if x != nil {
		return x.SubscriptionKey
	}
	return ""
}

func (x *SubscriptionPlan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubscriptionPlan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SubscriptionPlan) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetOrganisationPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrganisationPlanRequest) Reset() {
	*x = GetOrganisationPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_v1_private_subscriptions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrganisationPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrganisationPlanRequest) ProtoMessage() {}

func (x *GetOrganisationPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_v1_private_subscriptions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrganisationPlanRequest.ProtoReflect.Descriptor instead.
func (*GetOrganisationPlanRequest) Descriptor() ([]byte, []int) {
	return file_subscriptions_v1_private_subscriptions_proto_rawDescGZIP(), []int{1}
}

type GetOrganisationPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan *SubscriptionPlan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *GetOrganisationPlanResponse) Reset() {
	*x = GetOrganisationPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_v1_private_subscriptions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrganisationPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrganisationPlanResponse) ProtoMessage() {}

func (x *GetOrganisationPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_v1_private_subscriptions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrganisationPlanResponse.ProtoReflect.Descriptor instead.
func (*GetOrganisationPlanResponse) Descriptor() ([]byte, []int) {
	return file_subscriptions_v1_private_subscriptions_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrganisationPlanResponse) GetPlan() *SubscriptionPlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

var File_subscriptions_v1_private_subscriptions_proto protoreflect.FileDescriptor

var file_subscriptions_v1_private_subscriptions_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x10, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x5d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61,
	0x6e, 0x32, 0xa4, 0x01, 0x0a, 0x1b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x84, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x34, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x35, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x66, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x65, 0x67, 0x61, 0x6e, 0x63, 0x72, 0x61,
	0x69, 0x67, 0x63, 0x72, 0x75, 0x69, 0x63, 0x6b, 0x73, 0x68, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x6f,
	0x6c, 0x6c, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x3b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x76, 0x31, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscriptions_v1_private_subscriptions_proto_rawDescOnce sync.Once
	file_subscriptions_v1_private_subscriptions_proto_rawDescData = file_subscriptions_v1_private_subscriptions_proto_rawDesc
)

func file_subscriptions_v1_private_subscriptions_proto_rawDescGZIP() []byte {
	file_subscriptions_v1_private_subscriptions_proto_rawDescOnce.Do(func() {
		file_subscriptions_v1_private_subscriptions_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscriptions_v1_private_subscriptions_proto_rawDescData)
	})
	return file_subscriptions_v1_private_subscriptions_proto_rawDescData
}

var file_subscriptions_v1_private_subscriptions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_subscriptions_v1_private_subscriptions_proto_goTypes = []any{
	(*SubscriptionPlan)(nil),            // 0: subscriptions.v1.private.SubscriptionPlan
	(*GetOrganisationPlanRequest)(nil),  // 1: subscriptions.v1.private.GetOrganisationPlanRequest
	(*GetOrganisationPlanResponse)(nil), // 2: subscriptions.v1.private.GetOrganisationPlanResponse
}
var file_subscriptions_v1_private_subscriptions_proto_depIdxs = []int32{
	0, // 0: subscriptions.v1.private.GetOrganisationPlanResponse.plan:type_name -> subscriptions.v1.private.SubscriptionPlan
	1, // 1: subscriptions.v1.private.PrivateSubscriptionsService.GetOrganisationPlan:input_type -> subscriptions.v1.private.GetOrganisationPlanRequest
	2, // 2: subscriptions.v1.private.PrivateSubscriptionsService.GetOrganisationPlan:output_type -> subscriptions.v1.private.GetOrganisationPlanResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_subscriptions_v1_private_subscriptions_proto_init() }
func file_subscriptions_v1_private_subscriptions_proto_init() {
	if File_subscriptions_v1_private_subscriptions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscriptions_v1_private_subscriptions_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SubscriptionPlan); i {
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
		file_subscriptions_v1_private_subscriptions_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrganisationPlanRequest); i {
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
		file_subscriptions_v1_private_subscriptions_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrganisationPlanResponse); i {
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
			RawDescriptor: file_subscriptions_v1_private_subscriptions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscriptions_v1_private_subscriptions_proto_goTypes,
		DependencyIndexes: file_subscriptions_v1_private_subscriptions_proto_depIdxs,
		MessageInfos:      file_subscriptions_v1_private_subscriptions_proto_msgTypes,
	}.Build()
	File_subscriptions_v1_private_subscriptions_proto = out.File
	file_subscriptions_v1_private_subscriptions_proto_rawDesc = nil
	file_subscriptions_v1_private_subscriptions_proto_goTypes = nil
	file_subscriptions_v1_private_subscriptions_proto_depIdxs = nil
}
