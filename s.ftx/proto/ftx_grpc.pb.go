// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ftxproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FtxClient is the client API for Ftx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FtxClient interface {
	ListAccountDeposits(ctx context.Context, in *ListAccountDepositsRequest, opts ...grpc.CallOption) (*ListAccountDepositsResponse, error)
}

type ftxClient struct {
	cc grpc.ClientConnInterface
}

func NewFtxClient(cc grpc.ClientConnInterface) FtxClient {
	return &ftxClient{cc}
}

func (c *ftxClient) ListAccountDeposits(ctx context.Context, in *ListAccountDepositsRequest, opts ...grpc.CallOption) (*ListAccountDepositsResponse, error) {
	out := new(ListAccountDepositsResponse)
	err := c.cc.Invoke(ctx, "/ftx/ListAccountDeposits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FtxServer is the server API for Ftx service.
// All implementations must embed UnimplementedFtxServer
// for forward compatibility
type FtxServer interface {
	ListAccountDeposits(context.Context, *ListAccountDepositsRequest) (*ListAccountDepositsResponse, error)
	mustEmbedUnimplementedFtxServer()
}

// UnimplementedFtxServer must be embedded to have forward compatible implementations.
type UnimplementedFtxServer struct {
}

func (*UnimplementedFtxServer) ListAccountDeposits(context.Context, *ListAccountDepositsRequest) (*ListAccountDepositsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccountDeposits not implemented")
}
func (*UnimplementedFtxServer) mustEmbedUnimplementedFtxServer() {}

func RegisterFtxServer(s *grpc.Server, srv FtxServer) {
	s.RegisterService(&_Ftx_serviceDesc, srv)
}

func _Ftx_ListAccountDeposits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountDepositsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FtxServer).ListAccountDeposits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftx/ListAccountDeposits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FtxServer).ListAccountDeposits(ctx, req.(*ListAccountDepositsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ftx_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ftx",
	HandlerType: (*FtxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAccountDeposits",
			Handler:    _Ftx_ListAccountDeposits_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s.ftx/proto/ftx.proto",
}