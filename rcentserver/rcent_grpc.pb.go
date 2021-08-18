// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RCenterClient is the client API for RCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RCenterClient interface {
	GetRoomList(ctx context.Context, opts ...grpc.CallOption) (RCenter_GetRoomListClient, error)
	GetRoom(ctx context.Context, opts ...grpc.CallOption) (RCenter_GetRoomClient, error)
	DelRoom(ctx context.Context, opts ...grpc.CallOption) (RCenter_DelRoomClient, error)
	UpdateRoomList(ctx context.Context, opts ...grpc.CallOption) (RCenter_UpdateRoomListClient, error)
	GetToken(ctx context.Context, opts ...grpc.CallOption) (RCenter_GetTokenClient, error)
	RemovePlayer(ctx context.Context, opts ...grpc.CallOption) (RCenter_RemovePlayerClient, error)
}

type rCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewRCenterClient(cc grpc.ClientConnInterface) RCenterClient {
	return &rCenterClient{cc}
}

func (c *rCenterClient) GetRoomList(ctx context.Context, opts ...grpc.CallOption) (RCenter_GetRoomListClient, error) {
	stream, err := c.cc.NewStream(ctx, &RCenter_ServiceDesc.Streams[0], "/RCenter/GetRoomList", opts...)
	if err != nil {
		return nil, err
	}
	x := &rCenterGetRoomListClient{stream}
	return x, nil
}

type RCenter_GetRoomListClient interface {
	Send(*ReqGetRoomList) error
	Recv() (*RetGetRoomList, error)
	grpc.ClientStream
}

type rCenterGetRoomListClient struct {
	grpc.ClientStream
}

func (x *rCenterGetRoomListClient) Send(m *ReqGetRoomList) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rCenterGetRoomListClient) Recv() (*RetGetRoomList, error) {
	m := new(RetGetRoomList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rCenterClient) GetRoom(ctx context.Context, opts ...grpc.CallOption) (RCenter_GetRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &RCenter_ServiceDesc.Streams[1], "/RCenter/GetRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &rCenterGetRoomClient{stream}
	return x, nil
}

type RCenter_GetRoomClient interface {
	Send(*ReqGetRoom) error
	Recv() (*RetGetRoom, error)
	grpc.ClientStream
}

type rCenterGetRoomClient struct {
	grpc.ClientStream
}

func (x *rCenterGetRoomClient) Send(m *ReqGetRoom) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rCenterGetRoomClient) Recv() (*RetGetRoom, error) {
	m := new(RetGetRoom)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rCenterClient) DelRoom(ctx context.Context, opts ...grpc.CallOption) (RCenter_DelRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &RCenter_ServiceDesc.Streams[2], "/RCenter/DelRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &rCenterDelRoomClient{stream}
	return x, nil
}

type RCenter_DelRoomClient interface {
	Send(*ReqDelRoom) error
	Recv() (*RetDelRoom, error)
	grpc.ClientStream
}

type rCenterDelRoomClient struct {
	grpc.ClientStream
}

func (x *rCenterDelRoomClient) Send(m *ReqDelRoom) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rCenterDelRoomClient) Recv() (*RetDelRoom, error) {
	m := new(RetDelRoom)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rCenterClient) UpdateRoomList(ctx context.Context, opts ...grpc.CallOption) (RCenter_UpdateRoomListClient, error) {
	stream, err := c.cc.NewStream(ctx, &RCenter_ServiceDesc.Streams[3], "/RCenter/UpdateRoomList", opts...)
	if err != nil {
		return nil, err
	}
	x := &rCenterUpdateRoomListClient{stream}
	return x, nil
}

type RCenter_UpdateRoomListClient interface {
	Send(*ReqUpdateRoomList) error
	Recv() (*RetUpdateRoomList, error)
	grpc.ClientStream
}

type rCenterUpdateRoomListClient struct {
	grpc.ClientStream
}

func (x *rCenterUpdateRoomListClient) Send(m *ReqUpdateRoomList) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rCenterUpdateRoomListClient) Recv() (*RetUpdateRoomList, error) {
	m := new(RetUpdateRoomList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rCenterClient) GetToken(ctx context.Context, opts ...grpc.CallOption) (RCenter_GetTokenClient, error) {
	stream, err := c.cc.NewStream(ctx, &RCenter_ServiceDesc.Streams[4], "/RCenter/GetToken", opts...)
	if err != nil {
		return nil, err
	}
	x := &rCenterGetTokenClient{stream}
	return x, nil
}

type RCenter_GetTokenClient interface {
	Send(*ReqGetToken) error
	Recv() (*RetGetToken, error)
	grpc.ClientStream
}

type rCenterGetTokenClient struct {
	grpc.ClientStream
}

func (x *rCenterGetTokenClient) Send(m *ReqGetToken) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rCenterGetTokenClient) Recv() (*RetGetToken, error) {
	m := new(RetGetToken)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rCenterClient) RemovePlayer(ctx context.Context, opts ...grpc.CallOption) (RCenter_RemovePlayerClient, error) {
	stream, err := c.cc.NewStream(ctx, &RCenter_ServiceDesc.Streams[5], "/RCenter/RemovePlayer", opts...)
	if err != nil {
		return nil, err
	}
	x := &rCenterRemovePlayerClient{stream}
	return x, nil
}

type RCenter_RemovePlayerClient interface {
	Send(*ReqRemovePlayer) error
	Recv() (*RetRemovePlayer, error)
	grpc.ClientStream
}

type rCenterRemovePlayerClient struct {
	grpc.ClientStream
}

func (x *rCenterRemovePlayerClient) Send(m *ReqRemovePlayer) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rCenterRemovePlayerClient) Recv() (*RetRemovePlayer, error) {
	m := new(RetRemovePlayer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RCenterServer is the server API for RCenter service.
// All implementations must embed UnimplementedRCenterServer
// for forward compatibility
type RCenterServer interface {
	GetRoomList(RCenter_GetRoomListServer) error
	GetRoom(RCenter_GetRoomServer) error
	DelRoom(RCenter_DelRoomServer) error
	UpdateRoomList(RCenter_UpdateRoomListServer) error
	GetToken(RCenter_GetTokenServer) error
	RemovePlayer(RCenter_RemovePlayerServer) error
	mustEmbedUnimplementedRCenterServer()
}

// UnimplementedRCenterServer must be embedded to have forward compatible implementations.
type UnimplementedRCenterServer struct {
}

func (UnimplementedRCenterServer) GetRoomList(RCenter_GetRoomListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}
func (UnimplementedRCenterServer) GetRoom(RCenter_GetRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedRCenterServer) DelRoom(RCenter_DelRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method DelRoom not implemented")
}
func (UnimplementedRCenterServer) UpdateRoomList(RCenter_UpdateRoomListServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateRoomList not implemented")
}
func (UnimplementedRCenterServer) GetToken(RCenter_GetTokenServer) error {
	return status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedRCenterServer) RemovePlayer(RCenter_RemovePlayerServer) error {
	return status.Errorf(codes.Unimplemented, "method RemovePlayer not implemented")
}
func (UnimplementedRCenterServer) mustEmbedUnimplementedRCenterServer() {}

// UnsafeRCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RCenterServer will
// result in compilation errors.
type UnsafeRCenterServer interface {
	mustEmbedUnimplementedRCenterServer()
}

func RegisterRCenterServer(s grpc.ServiceRegistrar, srv RCenterServer) {
	s.RegisterService(&RCenter_ServiceDesc, srv)
}

func _RCenter_GetRoomList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RCenterServer).GetRoomList(&rCenterGetRoomListServer{stream})
}

type RCenter_GetRoomListServer interface {
	Send(*RetGetRoomList) error
	Recv() (*ReqGetRoomList, error)
	grpc.ServerStream
}

type rCenterGetRoomListServer struct {
	grpc.ServerStream
}

func (x *rCenterGetRoomListServer) Send(m *RetGetRoomList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rCenterGetRoomListServer) Recv() (*ReqGetRoomList, error) {
	m := new(ReqGetRoomList)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RCenter_GetRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RCenterServer).GetRoom(&rCenterGetRoomServer{stream})
}

type RCenter_GetRoomServer interface {
	Send(*RetGetRoom) error
	Recv() (*ReqGetRoom, error)
	grpc.ServerStream
}

type rCenterGetRoomServer struct {
	grpc.ServerStream
}

func (x *rCenterGetRoomServer) Send(m *RetGetRoom) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rCenterGetRoomServer) Recv() (*ReqGetRoom, error) {
	m := new(ReqGetRoom)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RCenter_DelRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RCenterServer).DelRoom(&rCenterDelRoomServer{stream})
}

type RCenter_DelRoomServer interface {
	Send(*RetDelRoom) error
	Recv() (*ReqDelRoom, error)
	grpc.ServerStream
}

type rCenterDelRoomServer struct {
	grpc.ServerStream
}

func (x *rCenterDelRoomServer) Send(m *RetDelRoom) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rCenterDelRoomServer) Recv() (*ReqDelRoom, error) {
	m := new(ReqDelRoom)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RCenter_UpdateRoomList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RCenterServer).UpdateRoomList(&rCenterUpdateRoomListServer{stream})
}

type RCenter_UpdateRoomListServer interface {
	Send(*RetUpdateRoomList) error
	Recv() (*ReqUpdateRoomList, error)
	grpc.ServerStream
}

type rCenterUpdateRoomListServer struct {
	grpc.ServerStream
}

func (x *rCenterUpdateRoomListServer) Send(m *RetUpdateRoomList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rCenterUpdateRoomListServer) Recv() (*ReqUpdateRoomList, error) {
	m := new(ReqUpdateRoomList)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RCenter_GetToken_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RCenterServer).GetToken(&rCenterGetTokenServer{stream})
}

type RCenter_GetTokenServer interface {
	Send(*RetGetToken) error
	Recv() (*ReqGetToken, error)
	grpc.ServerStream
}

type rCenterGetTokenServer struct {
	grpc.ServerStream
}

func (x *rCenterGetTokenServer) Send(m *RetGetToken) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rCenterGetTokenServer) Recv() (*ReqGetToken, error) {
	m := new(ReqGetToken)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RCenter_RemovePlayer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RCenterServer).RemovePlayer(&rCenterRemovePlayerServer{stream})
}

type RCenter_RemovePlayerServer interface {
	Send(*RetRemovePlayer) error
	Recv() (*ReqRemovePlayer, error)
	grpc.ServerStream
}

type rCenterRemovePlayerServer struct {
	grpc.ServerStream
}

func (x *rCenterRemovePlayerServer) Send(m *RetRemovePlayer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rCenterRemovePlayerServer) Recv() (*ReqRemovePlayer, error) {
	m := new(ReqRemovePlayer)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RCenter_ServiceDesc is the grpc.ServiceDesc for RCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RCenter",
	HandlerType: (*RCenterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRoomList",
			Handler:       _RCenter_GetRoomList_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetRoom",
			Handler:       _RCenter_GetRoom_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DelRoom",
			Handler:       _RCenter_DelRoom_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateRoomList",
			Handler:       _RCenter_UpdateRoomList_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetToken",
			Handler:       _RCenter_GetToken_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RemovePlayer",
			Handler:       _RCenter_RemovePlayer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rcent.proto",
}