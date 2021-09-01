// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package discordproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DiscordClient is the client API for Discord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscordClient interface {
	SendMsgToChannel(ctx context.Context, in *SendMsgToChannelRequest, opts ...grpc.CallOption) (*SendMsgToChannelResponse, error)
	SendMsgToPrivateChannel(ctx context.Context, in *SendMsgToPrivateChannelRequest, opts ...grpc.CallOption) (*SendMsgToPrivateChannelResponse, error)
	ReadUserRoles(ctx context.Context, in *ReadUserRolesRequest, opts ...grpc.CallOption) (*ReadUserRolesResponse, error)
}

type discordClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscordClient(cc grpc.ClientConnInterface) DiscordClient {
	return &discordClient{cc}
}

func (c *discordClient) SendMsgToChannel(ctx context.Context, in *SendMsgToChannelRequest, opts ...grpc.CallOption) (*SendMsgToChannelResponse, error) {
	out := new(SendMsgToChannelResponse)
	err := c.cc.Invoke(ctx, "/discord/SendMsgToChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordClient) SendMsgToPrivateChannel(ctx context.Context, in *SendMsgToPrivateChannelRequest, opts ...grpc.CallOption) (*SendMsgToPrivateChannelResponse, error) {
	out := new(SendMsgToPrivateChannelResponse)
	err := c.cc.Invoke(ctx, "/discord/SendMsgToPrivateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordClient) ReadUserRoles(ctx context.Context, in *ReadUserRolesRequest, opts ...grpc.CallOption) (*ReadUserRolesResponse, error) {
	out := new(ReadUserRolesResponse)
	err := c.cc.Invoke(ctx, "/discord/ReadUserRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscordServer is the server API for Discord service.
// All implementations must embed UnimplementedDiscordServer
// for forward compatibility
type DiscordServer interface {
	SendMsgToChannel(context.Context, *SendMsgToChannelRequest) (*SendMsgToChannelResponse, error)
	SendMsgToPrivateChannel(context.Context, *SendMsgToPrivateChannelRequest) (*SendMsgToPrivateChannelResponse, error)
	ReadUserRoles(context.Context, *ReadUserRolesRequest) (*ReadUserRolesResponse, error)
	mustEmbedUnimplementedDiscordServer()
}

// UnimplementedDiscordServer must be embedded to have forward compatible implementations.
type UnimplementedDiscordServer struct {
}

func (*UnimplementedDiscordServer) SendMsgToChannel(context.Context, *SendMsgToChannelRequest) (*SendMsgToChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsgToChannel not implemented")
}
func (*UnimplementedDiscordServer) SendMsgToPrivateChannel(context.Context, *SendMsgToPrivateChannelRequest) (*SendMsgToPrivateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsgToPrivateChannel not implemented")
}
func (*UnimplementedDiscordServer) ReadUserRoles(context.Context, *ReadUserRolesRequest) (*ReadUserRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUserRoles not implemented")
}
func (*UnimplementedDiscordServer) mustEmbedUnimplementedDiscordServer() {}

func RegisterDiscordServer(s *grpc.Server, srv DiscordServer) {
	s.RegisterService(&_Discord_serviceDesc, srv)
}

func _Discord_SendMsgToChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgToChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServer).SendMsgToChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discord/SendMsgToChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServer).SendMsgToChannel(ctx, req.(*SendMsgToChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discord_SendMsgToPrivateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgToPrivateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServer).SendMsgToPrivateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discord/SendMsgToPrivateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServer).SendMsgToPrivateChannel(ctx, req.(*SendMsgToPrivateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discord_ReadUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServer).ReadUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discord/ReadUserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServer).ReadUserRoles(ctx, req.(*ReadUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discord_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discord",
	HandlerType: (*DiscordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsgToChannel",
			Handler:    _Discord_SendMsgToChannel_Handler,
		},
		{
			MethodName: "SendMsgToPrivateChannel",
			Handler:    _Discord_SendMsgToPrivateChannel_Handler,
		},
		{
			MethodName: "ReadUserRoles",
			Handler:    _Discord_ReadUserRoles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s.discord/proto/discord.proto",
}
