// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: protos/introduce.proto

package pb

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

// IntroInterfaceClient is the client API for IntroInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntroInterfaceClient interface {
	// Executes Grep on a remote machine.
	RequestIntro(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error)
	InformDeath(ctx context.Context, in *DeathInfo, opts ...grpc.CallOption) (*DeathAck, error)
}

type introInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewIntroInterfaceClient(cc grpc.ClientConnInterface) IntroInterfaceClient {
	return &introInterfaceClient{cc}
}

func (c *introInterfaceClient) RequestIntro(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error) {
	out := new(IntroResponse)
	err := c.cc.Invoke(ctx, "/grep.IntroInterface/RequestIntro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *introInterfaceClient) InformDeath(ctx context.Context, in *DeathInfo, opts ...grpc.CallOption) (*DeathAck, error) {
	out := new(DeathAck)
	err := c.cc.Invoke(ctx, "/grep.IntroInterface/InformDeath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntroInterfaceServer is the server API for IntroInterface service.
// All implementations must embed UnimplementedIntroInterfaceServer
// for forward compatibility
type IntroInterfaceServer interface {
	// Executes Grep on a remote machine.
	RequestIntro(context.Context, *IntroRequest) (*IntroResponse, error)
	InformDeath(context.Context, *DeathInfo) (*DeathAck, error)
	mustEmbedUnimplementedIntroInterfaceServer()
}

// UnimplementedIntroInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedIntroInterfaceServer struct {
}

func (UnimplementedIntroInterfaceServer) RequestIntro(context.Context, *IntroRequest) (*IntroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestIntro not implemented")
}
func (UnimplementedIntroInterfaceServer) InformDeath(context.Context, *DeathInfo) (*DeathAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InformDeath not implemented")
}
func (UnimplementedIntroInterfaceServer) mustEmbedUnimplementedIntroInterfaceServer() {}

// UnsafeIntroInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntroInterfaceServer will
// result in compilation errors.
type UnsafeIntroInterfaceServer interface {
	mustEmbedUnimplementedIntroInterfaceServer()
}

func RegisterIntroInterfaceServer(s grpc.ServiceRegistrar, srv IntroInterfaceServer) {
	s.RegisterService(&IntroInterface_ServiceDesc, srv)
}

func _IntroInterface_RequestIntro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntroInterfaceServer).RequestIntro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grep.IntroInterface/RequestIntro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntroInterfaceServer).RequestIntro(ctx, req.(*IntroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntroInterface_InformDeath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeathInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntroInterfaceServer).InformDeath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grep.IntroInterface/InformDeath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntroInterfaceServer).InformDeath(ctx, req.(*DeathInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// IntroInterface_ServiceDesc is the grpc.ServiceDesc for IntroInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntroInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grep.IntroInterface",
	HandlerType: (*IntroInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestIntro",
			Handler:    _IntroInterface_RequestIntro_Handler,
		},
		{
			MethodName: "InformDeath",
			Handler:    _IntroInterface_InformDeath_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/introduce.proto",
}
