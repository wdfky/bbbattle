// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: roomserver.proto

package main

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

type RoomCmd int32

const (
	RoomCmd_RoomCmdNone       RoomCmd = 0
	RoomCmd_Login             RoomCmd = 1
	RoomCmd_HeartBeat         RoomCmd = 2
	RoomCmd_Action            RoomCmd = 3
	RoomCmd_AddDelPlayer      RoomCmd = 4 //添加删除用户
	RoomCmd_UpdatePlayer      RoomCmd = 5 //更新用户状态,这里为重量
	RoomCmd_BatchUpdatePlayer RoomCmd = 6 //批量更新
	//GameProgressNotice = 7;
	RoomCmd_Move       RoomCmd = 7
	RoomCmd_BatchMove  RoomCmd = 8
	RoomCmd_SpawnFood  RoomCmd = 9
	RoomCmd_Reborn     RoomCmd = 10 //玩家重生
	RoomCmd_Start      RoomCmd = 11 //开始游戏
	RoomCmd_End        RoomCmd = 12 //结束游戏
	RoomCmd_MapInit    RoomCmd = 13
	RoomCmd_AddFood    RoomCmd = 14
	RoomCmd_DelFood    RoomCmd = 15
	RoomCmd_RoomNotice RoomCmd = 16
	RoomCmd_EatFood    RoomCmd = 17
	RoomCmd_EatPlayer  RoomCmd = 18
)

// Enum value maps for RoomCmd.
var (
	RoomCmd_name = map[int32]string{
		0:  "RoomCmdNone",
		1:  "Login",
		2:  "HeartBeat",
		3:  "Action",
		4:  "AddDelPlayer",
		5:  "UpdatePlayer",
		6:  "BatchUpdatePlayer",
		7:  "Move",
		8:  "BatchMove",
		9:  "SpawnFood",
		10: "Reborn",
		11: "Start",
		12: "End",
		13: "MapInit",
		14: "AddFood",
		15: "DelFood",
		16: "RoomNotice",
		17: "EatFood",
		18: "EatPlayer",
	}
	RoomCmd_value = map[string]int32{
		"RoomCmdNone":       0,
		"Login":             1,
		"HeartBeat":         2,
		"Action":            3,
		"AddDelPlayer":      4,
		"UpdatePlayer":      5,
		"BatchUpdatePlayer": 6,
		"Move":              7,
		"BatchMove":         8,
		"SpawnFood":         9,
		"Reborn":            10,
		"Start":             11,
		"End":               12,
		"MapInit":           13,
		"AddFood":           14,
		"DelFood":           15,
		"RoomNotice":        16,
		"EatFood":           17,
		"EatPlayer":         18,
	}
)

func (x RoomCmd) Enum() *RoomCmd {
	p := new(RoomCmd)
	*p = x
	return p
}

func (x RoomCmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomCmd) Descriptor() protoreflect.EnumDescriptor {
	return file_roomserver_proto_enumTypes[0].Descriptor()
}

func (RoomCmd) Type() protoreflect.EnumType {
	return &file_roomserver_proto_enumTypes[0]
}

func (x RoomCmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomCmd.Descriptor instead.
func (RoomCmd) EnumDescriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{0}
}

type RoomStateType int32

const (
	RoomStateType_Open  RoomStateType = 0
	RoomStateType_Close RoomStateType = 1
)

// Enum value maps for RoomStateType.
var (
	RoomStateType_name = map[int32]string{
		0: "Open",
		1: "Close",
	}
	RoomStateType_value = map[string]int32{
		"Open":  0,
		"Close": 1,
	}
)

func (x RoomStateType) Enum() *RoomStateType {
	p := new(RoomStateType)
	*p = x
	return p
}

func (x RoomStateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomStateType) Descriptor() protoreflect.EnumDescriptor {
	return file_roomserver_proto_enumTypes[1].Descriptor()
}

func (RoomStateType) Type() protoreflect.EnumType {
	return &file_roomserver_proto_enumTypes[1]
}

func (x RoomStateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomStateType.Descriptor instead.
func (RoomStateType) EnumDescriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{1}
}

type RoomCloseType int32

const (
	RoomCloseType_NoBody RoomCloseType = 0
)

// Enum value maps for RoomCloseType.
var (
	RoomCloseType_name = map[int32]string{
		0: "NoBody",
	}
	RoomCloseType_value = map[string]int32{
		"NoBody": 0,
	}
)

func (x RoomCloseType) Enum() *RoomCloseType {
	p := new(RoomCloseType)
	*p = x
	return p
}

func (x RoomCloseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomCloseType) Descriptor() protoreflect.EnumDescriptor {
	return file_roomserver_proto_enumTypes[2].Descriptor()
}

func (RoomCloseType) Type() protoreflect.EnumType {
	return &file_roomserver_proto_enumTypes[2]
}

func (x RoomCloseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomCloseType.Descriptor instead.
func (RoomCloseType) EnumDescriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{2}
}

// message Commond{
//     int32 cmd = 1;
//     string data = 2;
// }
type ReqRoomLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ReqRoomLogin) Reset() {
	*x = ReqRoomLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqRoomLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqRoomLogin) ProtoMessage() {}

func (x *ReqRoomLogin) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqRoomLogin.ProtoReflect.Descriptor instead.
func (*ReqRoomLogin) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{0}
}

func (x *ReqRoomLogin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqRoomLogin) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RetRoomLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok          bool   `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	PlayerId    uint64 `protobuf:"varint,2,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	RoomId      uint32 `protobuf:"varint,3,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	Time        uint32 `protobuf:"varint,4,opt,name=Time,proto3" json:"Time,omitempty"`
	MapId       uint32 `protobuf:"varint,5,opt,name=MapId,proto3" json:"MapId,omitempty"`
	MinStartNum uint32 `protobuf:"varint,6,opt,name=MinStartNum,proto3" json:"MinStartNum,omitempty"`
}

func (x *RetRoomLogin) Reset() {
	*x = RetRoomLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetRoomLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetRoomLogin) ProtoMessage() {}

func (x *RetRoomLogin) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetRoomLogin.ProtoReflect.Descriptor instead.
func (*RetRoomLogin) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{1}
}

func (x *RetRoomLogin) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *RetRoomLogin) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *RetRoomLogin) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *RetRoomLogin) GetTime() uint32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RetRoomLogin) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *RetRoomLogin) GetMinStartNum() uint32 {
	if x != nil {
		return x.MinStartNum
	}
	return 0
}

type RetErrorMsgCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode uint32 `protobuf:"varint,1,opt,name=RetCode,proto3" json:"RetCode,omitempty"` //错误码
}

func (x *RetErrorMsgCmd) Reset() {
	*x = RetErrorMsgCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetErrorMsgCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetErrorMsgCmd) ProtoMessage() {}

func (x *RetErrorMsgCmd) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetErrorMsgCmd.ProtoReflect.Descriptor instead.
func (*RetErrorMsgCmd) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{2}
}

func (x *RetErrorMsgCmd) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

type MsgVector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (x *MsgVector) Reset() {
	*x = MsgVector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgVector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgVector) ProtoMessage() {}

func (x *MsgVector) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgVector.ProtoReflect.Descriptor instead.
func (*MsgVector) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{3}
}

func (x *MsgVector) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MsgVector) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

//移动才会触发吃事件，因此一起发过去
type MsgMove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId uint64     `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	Pos      *MsgVector `protobuf:"bytes,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	Rotation *MsgVector `protobuf:"bytes,3,opt,name=Rotation,proto3" json:"Rotation,omitempty"`
	//float Weight  = 4;
	Duration float32 `protobuf:"fixed32,4,opt,name=Duration,proto3" json:"Duration,omitempty"` //uint32 Duration = 4;
}

func (x *MsgMove) Reset() {
	*x = MsgMove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgMove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgMove) ProtoMessage() {}

func (x *MsgMove) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgMove.ProtoReflect.Descriptor instead.
func (*MsgMove) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{4}
}

func (x *MsgMove) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *MsgMove) GetPos() *MsgVector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *MsgMove) GetRotation() *MsgVector {
	if x != nil {
		return x.Rotation
	}
	return nil
}

func (x *MsgMove) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type MsgBatchMove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moves []*MsgMove `protobuf:"bytes,1,rep,name=Moves,proto3" json:"Moves,omitempty"`
}

func (x *MsgBatchMove) Reset() {
	*x = MsgBatchMove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgBatchMove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgBatchMove) ProtoMessage() {}

func (x *MsgBatchMove) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgBatchMove.ProtoReflect.Descriptor instead.
func (*MsgBatchMove) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{5}
}

func (x *MsgBatchMove) GetMoves() []*MsgMove {
	if x != nil {
		return x.Moves
	}
	return nil
}

type MsgWeight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId uint64  `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	Weight   float32 `protobuf:"fixed32,2,opt,name=Weight,proto3" json:"Weight,omitempty"`
}

func (x *MsgWeight) Reset() {
	*x = MsgWeight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgWeight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgWeight) ProtoMessage() {}

func (x *MsgWeight) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgWeight.ProtoReflect.Descriptor instead.
func (*MsgWeight) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{6}
}

func (x *MsgWeight) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *MsgWeight) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type MsgBatchWeight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weights []*MsgWeight `protobuf:"bytes,1,rep,name=Weights,proto3" json:"Weights,omitempty"`
}

func (x *MsgBatchWeight) Reset() {
	*x = MsgBatchWeight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgBatchWeight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgBatchWeight) ProtoMessage() {}

func (x *MsgBatchWeight) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgBatchWeight.ProtoReflect.Descriptor instead.
func (*MsgBatchWeight) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{7}
}

func (x *MsgBatchWeight) GetWeights() []*MsgWeight {
	if x != nil {
		return x.Weights
	}
	return nil
}

type MsgPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId    string     `protobuf:"bytes,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	PlayerId  uint64     `protobuf:"varint,2,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	Name      string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Pos       *MsgVector `protobuf:"bytes,4,opt,name=Pos,proto3" json:"Pos,omitempty"`
	Rotation  *MsgVector `protobuf:"bytes,5,opt,name=Rotation,proto3" json:"Rotation,omitempty"`
	IsOffline bool       `protobuf:"varint,6,opt,name=IsOffline,proto3" json:"IsOffline,omitempty"`
}

func (x *MsgPlayer) Reset() {
	*x = MsgPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPlayer) ProtoMessage() {}

func (x *MsgPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPlayer.ProtoReflect.Descriptor instead.
func (*MsgPlayer) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{8}
}

func (x *MsgPlayer) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *MsgPlayer) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *MsgPlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MsgPlayer) GetPos() *MsgVector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *MsgPlayer) GetRotation() *MsgVector {
	if x != nil {
		return x.Rotation
	}
	return nil
}

func (x *MsgPlayer) GetIsOffline() bool {
	if x != nil {
		return x.IsOffline
	}
	return false
}

type MsgAddDelPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerList []*MsgPlayer `protobuf:"bytes,1,rep,name=PlayerList,proto3" json:"PlayerList,omitempty"` //玩家id
	Ids        []uint32     `protobuf:"varint,2,rep,packed,name=Ids,proto3" json:"Ids,omitempty"`       //所在的roomid
}

func (x *MsgAddDelPlayer) Reset() {
	*x = MsgAddDelPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAddDelPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAddDelPlayer) ProtoMessage() {}

func (x *MsgAddDelPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgAddDelPlayer.ProtoReflect.Descriptor instead.
func (*MsgAddDelPlayer) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{9}
}

func (x *MsgAddDelPlayer) GetPlayerList() []*MsgPlayer {
	if x != nil {
		return x.PlayerList
	}
	return nil
}

func (x *MsgAddDelPlayer) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//某个room内的信息
type MsgRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*MsgPlayer `protobuf:"bytes,1,rep,name=Players,proto3" json:"Players,omitempty"`
	Length  uint32       `protobuf:"varint,2,opt,name=Length,proto3" json:"Length,omitempty"`
	Width   uint32       `protobuf:"varint,3,opt,name=Width,proto3" json:"Width,omitempty"`
}

func (x *MsgRoom) Reset() {
	*x = MsgRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRoom) ProtoMessage() {}

func (x *MsgRoom) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRoom.ProtoReflect.Descriptor instead.
func (*MsgRoom) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{10}
}

func (x *MsgRoom) GetPlayers() []*MsgPlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *MsgRoom) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *MsgRoom) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

//食物信息
type MsgFood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64     `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Pos *MsgVector `protobuf:"bytes,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (x *MsgFood) Reset() {
	*x = MsgFood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgFood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgFood) ProtoMessage() {}

func (x *MsgFood) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgFood.ProtoReflect.Descriptor instead.
func (*MsgFood) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{11}
}

func (x *MsgFood) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MsgFood) GetPos() *MsgVector {
	if x != nil {
		return x.Pos
	}
	return nil
}

type BatchMsgFood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodList []*MsgFood `protobuf:"bytes,1,rep,name=FoodList,proto3" json:"FoodList,omitempty"`
}

func (x *BatchMsgFood) Reset() {
	*x = BatchMsgFood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomserver_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchMsgFood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchMsgFood) ProtoMessage() {}

func (x *BatchMsgFood) ProtoReflect() protoreflect.Message {
	mi := &file_roomserver_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchMsgFood.ProtoReflect.Descriptor instead.
func (*BatchMsgFood) Descriptor() ([]byte, []int) {
	return file_roomserver_proto_rawDescGZIP(), []int{12}
}

func (x *BatchMsgFood) GetFoodList() []*MsgFood {
	if x != nil {
		return x.FoodList
	}
	return nil
}

var File_roomserver_proto protoreflect.FileDescriptor

var file_roomserver_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9e, 0x01, 0x0a,
	0x0c, 0x52, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x6b, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x70, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4d,
	0x69, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x2a, 0x0a,
	0x0e, 0x52, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x43, 0x6d, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x27, 0x0a, 0x09, 0x4d, 0x73, 0x67,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x01, 0x59, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x03, 0x50, 0x6f,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x03, 0x50, 0x6f, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x52, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x73, 0x67,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x0c,
	0x4d, 0x73, 0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x05,
	0x4d, 0x6f, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x73,
	0x67, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x05, 0x4d, 0x6f, 0x76, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x09,
	0x4d, 0x73, 0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x36, 0x0a,
	0x0e, 0x4d, 0x73, 0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x24, 0x0a, 0x07, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x4d, 0x73, 0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x07, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0xb7, 0x01, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x50,
	0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x73, 0x67, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x50, 0x6f, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x52, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x73,
	0x67, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x22,
	0x4f, 0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x41, 0x64, 0x64, 0x44, 0x65, 0x6c, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x73, 0x67, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x49, 0x64, 0x73,
	0x22, 0x5d, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x24, 0x0a, 0x07, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d,
	0x73, 0x67, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x22,
	0x37, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x03, 0x50, 0x6f,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x03, 0x50, 0x6f, 0x73, 0x22, 0x34, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4d, 0x73, 0x67, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x24, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x73, 0x67,
	0x46, 0x6f, 0x6f, 0x64, 0x52, 0x08, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x96,
	0x02, 0x0a, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6d, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x6f,
	0x6f, 0x6d, 0x43, 0x6d, 0x64, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x03, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x44, 0x65, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04,
	0x4d, 0x6f, 0x76, 0x65, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4d,
	0x6f, 0x76, 0x65, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x46, 0x6f,
	0x6f, 0x64, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x62, 0x6f, 0x72, 0x6e, 0x10, 0x0a,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0x0b, 0x12, 0x07, 0x0a, 0x03, 0x45,
	0x6e, 0x64, 0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x69, 0x74, 0x10,
	0x0d, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x10, 0x0e, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x65, 0x6c, 0x46, 0x6f, 0x6f, 0x64, 0x10, 0x0f, 0x12, 0x0e, 0x0a, 0x0a, 0x52,
	0x6f, 0x6f, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x10, 0x10, 0x12, 0x0b, 0x0a, 0x07, 0x45,
	0x61, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x10, 0x11, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x61, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x12, 0x2a, 0x24, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x10, 0x01, 0x2a, 0x1b, 0x0a,
	0x0d, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x6f, 0x42, 0x6f, 0x64, 0x79, 0x10, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x6d,
	0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_roomserver_proto_rawDescOnce sync.Once
	file_roomserver_proto_rawDescData = file_roomserver_proto_rawDesc
)

func file_roomserver_proto_rawDescGZIP() []byte {
	file_roomserver_proto_rawDescOnce.Do(func() {
		file_roomserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_roomserver_proto_rawDescData)
	})
	return file_roomserver_proto_rawDescData
}

var file_roomserver_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_roomserver_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_roomserver_proto_goTypes = []interface{}{
	(RoomCmd)(0),            // 0: RoomCmd
	(RoomStateType)(0),      // 1: RoomStateType
	(RoomCloseType)(0),      // 2: RoomCloseType
	(*ReqRoomLogin)(nil),    // 3: ReqRoomLogin
	(*RetRoomLogin)(nil),    // 4: RetRoomLogin
	(*RetErrorMsgCmd)(nil),  // 5: RetErrorMsgCmd
	(*MsgVector)(nil),       // 6: MsgVector
	(*MsgMove)(nil),         // 7: MsgMove
	(*MsgBatchMove)(nil),    // 8: MsgBatchMove
	(*MsgWeight)(nil),       // 9: MsgWeight
	(*MsgBatchWeight)(nil),  // 10: MsgBatchWeight
	(*MsgPlayer)(nil),       // 11: MsgPlayer
	(*MsgAddDelPlayer)(nil), // 12: MsgAddDelPlayer
	(*MsgRoom)(nil),         // 13: MsgRoom
	(*MsgFood)(nil),         // 14: MsgFood
	(*BatchMsgFood)(nil),    // 15: BatchMsgFood
}
var file_roomserver_proto_depIdxs = []int32{
	6,  // 0: MsgMove.Pos:type_name -> MsgVector
	6,  // 1: MsgMove.Rotation:type_name -> MsgVector
	7,  // 2: MsgBatchMove.Moves:type_name -> MsgMove
	9,  // 3: MsgBatchWeight.Weights:type_name -> MsgWeight
	6,  // 4: MsgPlayer.Pos:type_name -> MsgVector
	6,  // 5: MsgPlayer.Rotation:type_name -> MsgVector
	11, // 6: MsgAddDelPlayer.PlayerList:type_name -> MsgPlayer
	11, // 7: MsgRoom.Players:type_name -> MsgPlayer
	6,  // 8: MsgFood.Pos:type_name -> MsgVector
	14, // 9: BatchMsgFood.FoodList:type_name -> MsgFood
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_roomserver_proto_init() }
func file_roomserver_proto_init() {
	if File_roomserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_roomserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqRoomLogin); i {
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
		file_roomserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetRoomLogin); i {
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
		file_roomserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetErrorMsgCmd); i {
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
		file_roomserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgVector); i {
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
		file_roomserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgMove); i {
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
		file_roomserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgBatchMove); i {
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
		file_roomserver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgWeight); i {
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
		file_roomserver_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgBatchWeight); i {
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
		file_roomserver_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPlayer); i {
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
		file_roomserver_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAddDelPlayer); i {
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
		file_roomserver_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRoom); i {
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
		file_roomserver_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgFood); i {
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
		file_roomserver_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchMsgFood); i {
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
			RawDescriptor: file_roomserver_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_roomserver_proto_goTypes,
		DependencyIndexes: file_roomserver_proto_depIdxs,
		EnumInfos:         file_roomserver_proto_enumTypes,
		MessageInfos:      file_roomserver_proto_msgTypes,
	}.Build()
	File_roomserver_proto = out.File
	file_roomserver_proto_rawDesc = nil
	file_roomserver_proto_goTypes = nil
	file_roomserver_proto_depIdxs = nil
}
