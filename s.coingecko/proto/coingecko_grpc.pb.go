// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package coingeckoproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CoingeckoClient is the client API for Coingecko service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoingeckoClient interface {
	GetAssetLatestPriceByID(ctx context.Context, in *GetAssetLatestPriceByIDRequest, opts ...grpc.CallOption) (*GetAssetLatestPriceByIDResponse, error)
	GetAssetLatestPriceBySymbol(ctx context.Context, in *GetAssetLatestPriceBySymbolRequest, opts ...grpc.CallOption) (*GetAssetLatestPriceBySymbolResponse, error)
	GetATHByID(ctx context.Context, in *GetATHByIDRequest, opts ...grpc.CallOption) (*GetATHByIDResponse, error)
	GetATHBySymbol(ctx context.Context, in *GetATHBySymbolRequest, opts ...grpc.CallOption) (*GetATHBySymbolResponse, error)
	ListCoinIDs(ctx context.Context, in *ListCoinIDsRequest, opts ...grpc.CallOption) (*ListCoinIDsResponse, error)
}

type coingeckoClient struct {
	cc grpc.ClientConnInterface
}

func NewCoingeckoClient(cc grpc.ClientConnInterface) CoingeckoClient {
	return &coingeckoClient{cc}
}

func (c *coingeckoClient) GetAssetLatestPriceByID(ctx context.Context, in *GetAssetLatestPriceByIDRequest, opts ...grpc.CallOption) (*GetAssetLatestPriceByIDResponse, error) {
	out := new(GetAssetLatestPriceByIDResponse)
	err := c.cc.Invoke(ctx, "/coingecko/GetAssetLatestPriceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coingeckoClient) GetAssetLatestPriceBySymbol(ctx context.Context, in *GetAssetLatestPriceBySymbolRequest, opts ...grpc.CallOption) (*GetAssetLatestPriceBySymbolResponse, error) {
	out := new(GetAssetLatestPriceBySymbolResponse)
	err := c.cc.Invoke(ctx, "/coingecko/GetAssetLatestPriceBySymbol", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coingeckoClient) GetATHByID(ctx context.Context, in *GetATHByIDRequest, opts ...grpc.CallOption) (*GetATHByIDResponse, error) {
	out := new(GetATHByIDResponse)
	err := c.cc.Invoke(ctx, "/coingecko/GetATHByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coingeckoClient) GetATHBySymbol(ctx context.Context, in *GetATHBySymbolRequest, opts ...grpc.CallOption) (*GetATHBySymbolResponse, error) {
	out := new(GetATHBySymbolResponse)
	err := c.cc.Invoke(ctx, "/coingecko/GetATHBySymbol", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coingeckoClient) ListCoinIDs(ctx context.Context, in *ListCoinIDsRequest, opts ...grpc.CallOption) (*ListCoinIDsResponse, error) {
	out := new(ListCoinIDsResponse)
	err := c.cc.Invoke(ctx, "/coingecko/ListCoinIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoingeckoServer is the server API for Coingecko service.
// All implementations must embed UnimplementedCoingeckoServer
// for forward compatibility
type CoingeckoServer interface {
	GetAssetLatestPriceByID(context.Context, *GetAssetLatestPriceByIDRequest) (*GetAssetLatestPriceByIDResponse, error)
	GetAssetLatestPriceBySymbol(context.Context, *GetAssetLatestPriceBySymbolRequest) (*GetAssetLatestPriceBySymbolResponse, error)
	GetATHByID(context.Context, *GetATHByIDRequest) (*GetATHByIDResponse, error)
	GetATHBySymbol(context.Context, *GetATHBySymbolRequest) (*GetATHBySymbolResponse, error)
	ListCoinIDs(context.Context, *ListCoinIDsRequest) (*ListCoinIDsResponse, error)
	mustEmbedUnimplementedCoingeckoServer()
}

// UnimplementedCoingeckoServer must be embedded to have forward compatible implementations.
type UnimplementedCoingeckoServer struct {
}

func (*UnimplementedCoingeckoServer) GetAssetLatestPriceByID(context.Context, *GetAssetLatestPriceByIDRequest) (*GetAssetLatestPriceByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetLatestPriceByID not implemented")
}
func (*UnimplementedCoingeckoServer) GetAssetLatestPriceBySymbol(context.Context, *GetAssetLatestPriceBySymbolRequest) (*GetAssetLatestPriceBySymbolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetLatestPriceBySymbol not implemented")
}
func (*UnimplementedCoingeckoServer) GetATHByID(context.Context, *GetATHByIDRequest) (*GetATHByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetATHByID not implemented")
}
func (*UnimplementedCoingeckoServer) GetATHBySymbol(context.Context, *GetATHBySymbolRequest) (*GetATHBySymbolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetATHBySymbol not implemented")
}
func (*UnimplementedCoingeckoServer) ListCoinIDs(context.Context, *ListCoinIDsRequest) (*ListCoinIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCoinIDs not implemented")
}
func (*UnimplementedCoingeckoServer) mustEmbedUnimplementedCoingeckoServer() {}

func RegisterCoingeckoServer(s *grpc.Server, srv CoingeckoServer) {
	s.RegisterService(&_Coingecko_serviceDesc, srv)
}

func _Coingecko_GetAssetLatestPriceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetLatestPriceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoingeckoServer).GetAssetLatestPriceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coingecko/GetAssetLatestPriceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoingeckoServer).GetAssetLatestPriceByID(ctx, req.(*GetAssetLatestPriceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coingecko_GetAssetLatestPriceBySymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetLatestPriceBySymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoingeckoServer).GetAssetLatestPriceBySymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coingecko/GetAssetLatestPriceBySymbol",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoingeckoServer).GetAssetLatestPriceBySymbol(ctx, req.(*GetAssetLatestPriceBySymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coingecko_GetATHByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetATHByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoingeckoServer).GetATHByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coingecko/GetATHByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoingeckoServer).GetATHByID(ctx, req.(*GetATHByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coingecko_GetATHBySymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetATHBySymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoingeckoServer).GetATHBySymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coingecko/GetATHBySymbol",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoingeckoServer).GetATHBySymbol(ctx, req.(*GetATHBySymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coingecko_ListCoinIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoinIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoingeckoServer).ListCoinIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coingecko/ListCoinIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoingeckoServer).ListCoinIDs(ctx, req.(*ListCoinIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coingecko_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coingecko",
	HandlerType: (*CoingeckoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAssetLatestPriceByID",
			Handler:    _Coingecko_GetAssetLatestPriceByID_Handler,
		},
		{
			MethodName: "GetAssetLatestPriceBySymbol",
			Handler:    _Coingecko_GetAssetLatestPriceBySymbol_Handler,
		},
		{
			MethodName: "GetATHByID",
			Handler:    _Coingecko_GetATHByID_Handler,
		},
		{
			MethodName: "GetATHBySymbol",
			Handler:    _Coingecko_GetATHBySymbol_Handler,
		},
		{
			MethodName: "ListCoinIDs",
			Handler:    _Coingecko_ListCoinIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s.coingecko/proto/coingecko.proto",
}
