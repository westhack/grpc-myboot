// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	RpcExecute(ctx context.Context, in *RpcRequest, opts ...grpc.CallOption) (*RpcResponse, error)
	RpcJsonExecute(ctx context.Context, in *RpcJsonRequest, opts ...grpc.CallOption) (*RpcJsonResponse, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) RpcExecute(ctx context.Context, in *RpcRequest, opts ...grpc.CallOption) (*RpcResponse, error) {
	out := new(RpcResponse)
	err := c.cc.Invoke(ctx, "/proto.MyService/RpcExecute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) RpcJsonExecute(ctx context.Context, in *RpcJsonRequest, opts ...grpc.CallOption) (*RpcJsonResponse, error) {
	out := new(RpcJsonResponse)
	err := c.cc.Invoke(ctx, "/proto.MyService/RpcJsonExecute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility
type MyServiceServer interface {
	RpcExecute(context.Context, *RpcRequest) (*RpcResponse, error)
	RpcJsonExecute(context.Context, *RpcJsonRequest) (*RpcJsonResponse, error)
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (UnimplementedMyServiceServer) RpcExecute(context.Context, *RpcRequest) (*RpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcExecute not implemented")
}
func (UnimplementedMyServiceServer) RpcJsonExecute(context.Context, *RpcJsonRequest) (*RpcJsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcJsonExecute not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_RpcExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).RpcExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyService/RpcExecute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).RpcExecute(ctx, req.(*RpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_RpcJsonExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcJsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).RpcJsonExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyService/RpcJsonExecute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).RpcJsonExecute(ctx, req.(*RpcJsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcExecute",
			Handler:    _MyService_RpcExecute_Handler,
		},
		{
			MethodName: "RpcJsonExecute",
			Handler:    _MyService_RpcJsonExecute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
