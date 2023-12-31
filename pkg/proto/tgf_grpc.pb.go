// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: proto/tgf.proto

package pkg

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

// TwitchGeneralFollowsClient is the client API for TwitchGeneralFollows service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwitchGeneralFollowsClient interface {
	GetGeneralFollows(ctx context.Context, in *GetTGFRequest, opts ...grpc.CallOption) (*GetTGFResponse, error)
}

type twitchGeneralFollowsClient struct {
	cc grpc.ClientConnInterface
}

func NewTwitchGeneralFollowsClient(cc grpc.ClientConnInterface) TwitchGeneralFollowsClient {
	return &twitchGeneralFollowsClient{cc}
}

func (c *twitchGeneralFollowsClient) GetGeneralFollows(ctx context.Context, in *GetTGFRequest, opts ...grpc.CallOption) (*GetTGFResponse, error) {
	out := new(GetTGFResponse)
	err := c.cc.Invoke(ctx, "/tgf.TwitchGeneralFollows/GetGeneralFollows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwitchGeneralFollowsServer is the server API for TwitchGeneralFollows service.
// All implementations must embed UnimplementedTwitchGeneralFollowsServer
// for forward compatibility
type TwitchGeneralFollowsServer interface {
	GetGeneralFollows(context.Context, *GetTGFRequest) (*GetTGFResponse, error)
	mustEmbedUnimplementedTwitchGeneralFollowsServer()
}

// UnimplementedTwitchGeneralFollowsServer must be embedded to have forward compatible implementations.
type UnimplementedTwitchGeneralFollowsServer struct {
}

func (UnimplementedTwitchGeneralFollowsServer) GetGeneralFollows(context.Context, *GetTGFRequest) (*GetTGFResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeneralFollows not implemented")
}
func (UnimplementedTwitchGeneralFollowsServer) mustEmbedUnimplementedTwitchGeneralFollowsServer() {}

// UnsafeTwitchGeneralFollowsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwitchGeneralFollowsServer will
// result in compilation errors.
type UnsafeTwitchGeneralFollowsServer interface {
	mustEmbedUnimplementedTwitchGeneralFollowsServer()
}

func RegisterTwitchGeneralFollowsServer(s grpc.ServiceRegistrar, srv TwitchGeneralFollowsServer) {
	s.RegisterService(&TwitchGeneralFollows_ServiceDesc, srv)
}

func _TwitchGeneralFollows_GetGeneralFollows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTGFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitchGeneralFollowsServer).GetGeneralFollows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tgf.TwitchGeneralFollows/GetGeneralFollows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitchGeneralFollowsServer).GetGeneralFollows(ctx, req.(*GetTGFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TwitchGeneralFollows_ServiceDesc is the grpc.ServiceDesc for TwitchGeneralFollows service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TwitchGeneralFollows_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tgf.TwitchGeneralFollows",
	HandlerType: (*TwitchGeneralFollowsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGeneralFollows",
			Handler:    _TwitchGeneralFollows_GetGeneralFollows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tgf.proto",
}
