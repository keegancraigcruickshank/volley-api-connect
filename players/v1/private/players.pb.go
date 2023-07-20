// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: players/v1/private/players.proto

package playersv1private

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
	return file_players_v1_private_players_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_players_v1_private_players_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{0}
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit        *string `protobuf:"bytes,1,opt,name=unit,proto3,oneof" json:"unit,omitempty"`
	HouseNumber *string `protobuf:"bytes,2,opt,name=house_number,json=houseNumber,proto3,oneof" json:"house_number,omitempty"`
	Street      *string `protobuf:"bytes,3,opt,name=street,proto3,oneof" json:"street,omitempty"`
	Suburb      *string `protobuf:"bytes,4,opt,name=suburb,proto3,oneof" json:"suburb,omitempty"`
	State       *string `protobuf:"bytes,5,opt,name=state,proto3,oneof" json:"state,omitempty"`
	Country     *string `protobuf:"bytes,6,opt,name=country,proto3,oneof" json:"country,omitempty"`
	ZipPostcode *string `protobuf:"bytes,7,opt,name=zip_postcode,json=zipPostcode,proto3,oneof" json:"zip_postcode,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetUnit() string {
	if x != nil && x.Unit != nil {
		return *x.Unit
	}
	return ""
}

func (x *Address) GetHouseNumber() string {
	if x != nil && x.HouseNumber != nil {
		return *x.HouseNumber
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil && x.Street != nil {
		return *x.Street
	}
	return ""
}

func (x *Address) GetSuburb() string {
	if x != nil && x.Suburb != nil {
		return *x.Suburb
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil && x.State != nil {
		return *x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *Address) GetZipPostcode() string {
	if x != nil && x.ZipPostcode != nil {
		return *x.ZipPostcode
	}
	return ""
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  *string  `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3,oneof" json:"last_name,omitempty"`
	Email     *string  `protobuf:"bytes,4,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Mobile    *string  `protobuf:"bytes,5,opt,name=mobile,proto3,oneof" json:"mobile,omitempty"`
	Dob       *string  `protobuf:"bytes,6,opt,name=dob,proto3,oneof" json:"dob,omitempty"`
	Address   *Address `protobuf:"bytes,7,opt,name=address,proto3,oneof" json:"address,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{1}
}

func (x *Contact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Contact) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Contact) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *Contact) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *Contact) GetMobile() string {
	if x != nil && x.Mobile != nil {
		return *x.Mobile
	}
	return ""
}

func (x *Contact) GetDob() string {
	if x != nil && x.Dob != nil {
		return *x.Dob
	}
	return ""
}

func (x *Contact) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams            []string `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	PlayerDetails    *Contact `protobuf:"bytes,2,opt,name=player_details,json=playerDetails,proto3" json:"player_details,omitempty"`
	EmergencyContact *Contact `protobuf:"bytes,3,opt,name=emergency_contact,json=emergencyContact,proto3,oneof" json:"emergency_contact,omitempty"`
	Modified         int32    `protobuf:"varint,4,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{2}
}

func (x *Player) GetTeams() []string {
	if x != nil {
		return x.Teams
	}
	return nil
}

func (x *Player) GetPlayerDetails() *Contact {
	if x != nil {
		return x.PlayerDetails
	}
	return nil
}

func (x *Player) GetEmergencyContact() *Contact {
	if x != nil {
		return x.EmergencyContact
	}
	return nil
}

func (x *Player) GetModified() int32 {
	if x != nil {
		return x.Modified
	}
	return 0
}

type CreatePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *CreatePlayerRequest) Reset() {
	*x = CreatePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayerRequest) ProtoMessage() {}

func (x *CreatePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayerRequest.ProtoReflect.Descriptor instead.
func (*CreatePlayerRequest) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePlayerRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type CreatePlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePlayerResponse) Reset() {
	*x = CreatePlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayerResponse) ProtoMessage() {}

func (x *CreatePlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayerResponse.ProtoReflect.Descriptor instead.
func (*CreatePlayerResponse) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{4}
}

type RemovePlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *RemovePlayersRequest) Reset() {
	*x = RemovePlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePlayersRequest) ProtoMessage() {}

func (x *RemovePlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePlayersRequest.ProtoReflect.Descriptor instead.
func (*RemovePlayersRequest) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{5}
}

func (x *RemovePlayersRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type RemovePlayersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemovePlayersResponse) Reset() {
	*x = RemovePlayersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePlayersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePlayersResponse) ProtoMessage() {}

func (x *RemovePlayersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePlayersResponse.ProtoReflect.Descriptor instead.
func (*RemovePlayersResponse) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{6}
}

type ModifyPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *ModifyPlayerRequest) Reset() {
	*x = ModifyPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyPlayerRequest) ProtoMessage() {}

func (x *ModifyPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyPlayerRequest.ProtoReflect.Descriptor instead.
func (*ModifyPlayerRequest) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{7}
}

func (x *ModifyPlayerRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type ModifyPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ModifyPlayerResponse) Reset() {
	*x = ModifyPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyPlayerResponse) ProtoMessage() {}

func (x *ModifyPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyPlayerResponse.ProtoReflect.Descriptor instead.
func (*ModifyPlayerResponse) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{8}
}

type ListPlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         *string        `protobuf:"bytes,1,opt,name=query,proto3,oneof" json:"query,omitempty"`
	Page          *int32         `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
	PageSize      *int32         `protobuf:"varint,3,opt,name=pageSize,proto3,oneof" json:"pageSize,omitempty"`
	SortDirection *SortDirection `protobuf:"varint,4,opt,name=sortDirection,proto3,enum=players.v1.private.SortDirection,oneof" json:"sortDirection,omitempty"`
}

func (x *ListPlayersRequest) Reset() {
	*x = ListPlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlayersRequest) ProtoMessage() {}

func (x *ListPlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlayersRequest.ProtoReflect.Descriptor instead.
func (*ListPlayersRequest) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{9}
}

func (x *ListPlayersRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *ListPlayersRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListPlayersRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListPlayersRequest) GetSortDirection() SortDirection {
	if x != nil && x.SortDirection != nil {
		return *x.SortDirection
	}
	return SortDirection_ASC
}

type ListPlayersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players    []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	TotalCount int32     `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListPlayersResponse) Reset() {
	*x = ListPlayersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_v1_private_players_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlayersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlayersResponse) ProtoMessage() {}

func (x *ListPlayersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_players_v1_private_players_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlayersResponse.ProtoReflect.Descriptor instead.
func (*ListPlayersResponse) Descriptor() ([]byte, []int) {
	return file_players_v1_private_players_proto_rawDescGZIP(), []int{10}
}

func (x *ListPlayersResponse) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *ListPlayersResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_players_v1_private_players_proto protoreflect.FileDescriptor

var file_players_v1_private_players_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0xbd, 0x02, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x75, 0x72, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x06, 0x73, 0x75, 0x62, 0x75, 0x72, 0x62, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x7a, 0x69, 0x70, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52,
	0x0b, 0x7a, 0x69, 0x70, 0x50, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x75, 0x62, 0x75, 0x72, 0x62, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x7a, 0x69, 0x70, 0x5f, 0x70, 0x6f,
	0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x9c, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x64,
	0x6f, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x88,
	0x01, 0x01, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x48, 0x04, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64, 0x6f, 0x62, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xe3, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x0d, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x4d, 0x0a, 0x11, 0x65, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x48, 0x00, 0x52, 0x10, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x49, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28,
	0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x49, 0x0a, 0x13, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x16, 0x0a, 0x14,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe9, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1f, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x02, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x4c, 0x0a, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x0d, 0x73, 0x6f,
	0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x6b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x22, 0x0a,
	0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10,
	0x01, 0x32, 0xab, 0x03, 0x0a, 0x15, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x66, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x28, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x26, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65,
	0x65, 0x67, 0x61, 0x6e, 0x63, 0x72, 0x61, 0x69, 0x67, 0x63, 0x72, 0x75, 0x69, 0x63, 0x6b, 0x73,
	0x68, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x6f, 0x6c, 0x6c, 0x65, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x3b, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x76, 0x31, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_players_v1_private_players_proto_rawDescOnce sync.Once
	file_players_v1_private_players_proto_rawDescData = file_players_v1_private_players_proto_rawDesc
)

func file_players_v1_private_players_proto_rawDescGZIP() []byte {
	file_players_v1_private_players_proto_rawDescOnce.Do(func() {
		file_players_v1_private_players_proto_rawDescData = protoimpl.X.CompressGZIP(file_players_v1_private_players_proto_rawDescData)
	})
	return file_players_v1_private_players_proto_rawDescData
}

var file_players_v1_private_players_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_players_v1_private_players_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_players_v1_private_players_proto_goTypes = []interface{}{
	(SortDirection)(0),            // 0: players.v1.private.SortDirection
	(*Address)(nil),               // 1: players.v1.private.Address
	(*Contact)(nil),               // 2: players.v1.private.Contact
	(*Player)(nil),                // 3: players.v1.private.Player
	(*CreatePlayerRequest)(nil),   // 4: players.v1.private.CreatePlayerRequest
	(*CreatePlayerResponse)(nil),  // 5: players.v1.private.CreatePlayerResponse
	(*RemovePlayersRequest)(nil),  // 6: players.v1.private.RemovePlayersRequest
	(*RemovePlayersResponse)(nil), // 7: players.v1.private.RemovePlayersResponse
	(*ModifyPlayerRequest)(nil),   // 8: players.v1.private.ModifyPlayerRequest
	(*ModifyPlayerResponse)(nil),  // 9: players.v1.private.ModifyPlayerResponse
	(*ListPlayersRequest)(nil),    // 10: players.v1.private.ListPlayersRequest
	(*ListPlayersResponse)(nil),   // 11: players.v1.private.ListPlayersResponse
}
var file_players_v1_private_players_proto_depIdxs = []int32{
	1,  // 0: players.v1.private.Contact.address:type_name -> players.v1.private.Address
	2,  // 1: players.v1.private.Player.player_details:type_name -> players.v1.private.Contact
	2,  // 2: players.v1.private.Player.emergency_contact:type_name -> players.v1.private.Contact
	3,  // 3: players.v1.private.CreatePlayerRequest.player:type_name -> players.v1.private.Player
	3,  // 4: players.v1.private.ModifyPlayerRequest.player:type_name -> players.v1.private.Player
	0,  // 5: players.v1.private.ListPlayersRequest.sortDirection:type_name -> players.v1.private.SortDirection
	3,  // 6: players.v1.private.ListPlayersResponse.players:type_name -> players.v1.private.Player
	4,  // 7: players.v1.private.PrivatePlayersService.CreatePlayer:input_type -> players.v1.private.CreatePlayerRequest
	6,  // 8: players.v1.private.PrivatePlayersService.RemovePlayers:input_type -> players.v1.private.RemovePlayersRequest
	8,  // 9: players.v1.private.PrivatePlayersService.ModifyPlayer:input_type -> players.v1.private.ModifyPlayerRequest
	10, // 10: players.v1.private.PrivatePlayersService.ListPlayers:input_type -> players.v1.private.ListPlayersRequest
	5,  // 11: players.v1.private.PrivatePlayersService.CreatePlayer:output_type -> players.v1.private.CreatePlayerResponse
	7,  // 12: players.v1.private.PrivatePlayersService.RemovePlayers:output_type -> players.v1.private.RemovePlayersResponse
	9,  // 13: players.v1.private.PrivatePlayersService.ModifyPlayer:output_type -> players.v1.private.ModifyPlayerResponse
	11, // 14: players.v1.private.PrivatePlayersService.ListPlayers:output_type -> players.v1.private.ListPlayersResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_players_v1_private_players_proto_init() }
func file_players_v1_private_players_proto_init() {
	if File_players_v1_private_players_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_players_v1_private_players_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_players_v1_private_players_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_players_v1_private_players_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_players_v1_private_players_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayerRequest); i {
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
		file_players_v1_private_players_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayerResponse); i {
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
		file_players_v1_private_players_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePlayersRequest); i {
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
		file_players_v1_private_players_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePlayersResponse); i {
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
		file_players_v1_private_players_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyPlayerRequest); i {
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
		file_players_v1_private_players_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyPlayerResponse); i {
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
		file_players_v1_private_players_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlayersRequest); i {
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
		file_players_v1_private_players_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlayersResponse); i {
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
	file_players_v1_private_players_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_players_v1_private_players_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_players_v1_private_players_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_players_v1_private_players_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_players_v1_private_players_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_players_v1_private_players_proto_goTypes,
		DependencyIndexes: file_players_v1_private_players_proto_depIdxs,
		EnumInfos:         file_players_v1_private_players_proto_enumTypes,
		MessageInfos:      file_players_v1_private_players_proto_msgTypes,
	}.Build()
	File_players_v1_private_players_proto = out.File
	file_players_v1_private_players_proto_rawDesc = nil
	file_players_v1_private_players_proto_goTypes = nil
	file_players_v1_private_players_proto_depIdxs = nil
}
