// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: emacsheadless.proto

package proto

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

// HeadlessClient is the client API for Headless service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeadlessClient interface {
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Headless_SendMessageClient, error)
}

type headlessClient struct {
	cc grpc.ClientConnInterface
}

func NewHeadlessClient(cc grpc.ClientConnInterface) HeadlessClient {
	return &headlessClient{cc}
}

func (c *headlessClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Headless_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Headless_ServiceDesc.Streams[0], "/headless.Headless/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &headlessSendMessageClient{stream}
	return x, nil
}

type Headless_SendMessageClient interface {
	Send(*Cmd) error
	Recv() (*Cmd, error)
	grpc.ClientStream
}

type headlessSendMessageClient struct {
	grpc.ClientStream
}

func (x *headlessSendMessageClient) Send(m *Cmd) error {
	return x.ClientStream.SendMsg(m)
}

func (x *headlessSendMessageClient) Recv() (*Cmd, error) {
	m := new(Cmd)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HeadlessServer is the server API for Headless service.
// All implementations must embed UnimplementedHeadlessServer
// for forward compatibility
type HeadlessServer interface {
	SendMessage(Headless_SendMessageServer) error
	mustEmbedUnimplementedHeadlessServer()
}

// UnimplementedHeadlessServer must be embedded to have forward compatible implementations.
type UnimplementedHeadlessServer struct {
}

func (UnimplementedHeadlessServer) SendMessage(Headless_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedHeadlessServer) mustEmbedUnimplementedHeadlessServer() {}

// UnsafeHeadlessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeadlessServer will
// result in compilation errors.
type UnsafeHeadlessServer interface {
	mustEmbedUnimplementedHeadlessServer()
}

func RegisterHeadlessServer(s grpc.ServiceRegistrar, srv HeadlessServer) {
	s.RegisterService(&Headless_ServiceDesc, srv)
}

func _Headless_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HeadlessServer).SendMessage(&headlessSendMessageServer{stream})
}

type Headless_SendMessageServer interface {
	Send(*Cmd) error
	Recv() (*Cmd, error)
	grpc.ServerStream
}

type headlessSendMessageServer struct {
	grpc.ServerStream
}

func (x *headlessSendMessageServer) Send(m *Cmd) error {
	return x.ServerStream.SendMsg(m)
}

func (x *headlessSendMessageServer) Recv() (*Cmd, error) {
	m := new(Cmd)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Headless_ServiceDesc is the grpc.ServiceDesc for Headless service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Headless_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "headless.Headless",
	HandlerType: (*HeadlessServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _Headless_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "emacsheadless.proto",
}
