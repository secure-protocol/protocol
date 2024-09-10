// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: signer.proto

package signer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Signer_Sign_FullMethodName      = "/signer.Signer/Sign"
	Signer_Gen_FullMethodName       = "/signer.Signer/Gen"
	Signer_Migration_FullMethodName = "/signer.Signer/Migration"
)

// SignerClient is the client API for Signer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignerClient interface {
	Sign(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignRes, error)
	Gen(ctx context.Context, in *GenReq, opts ...grpc.CallOption) (*GenRes, error)
	Migration(ctx context.Context, in *MigrationReq, opts ...grpc.CallOption) (*MigrationRes, error)
}

type signerClient struct {
	cc grpc.ClientConnInterface
}

func NewSignerClient(cc grpc.ClientConnInterface) SignerClient {
	return &signerClient{cc}
}

func (c *signerClient) Sign(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignRes)
	err := c.cc.Invoke(ctx, Signer_Sign_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signerClient) Gen(ctx context.Context, in *GenReq, opts ...grpc.CallOption) (*GenRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenRes)
	err := c.cc.Invoke(ctx, Signer_Gen_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signerClient) Migration(ctx context.Context, in *MigrationReq, opts ...grpc.CallOption) (*MigrationRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MigrationRes)
	err := c.cc.Invoke(ctx, Signer_Migration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignerServer is the server API for Signer service.
// All implementations must embed UnimplementedSignerServer
// for forward compatibility.
type SignerServer interface {
	Sign(context.Context, *SignReq) (*SignRes, error)
	Gen(context.Context, *GenReq) (*GenRes, error)
	Migration(context.Context, *MigrationReq) (*MigrationRes, error)
	mustEmbedUnimplementedSignerServer()
}

// UnimplementedSignerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSignerServer struct{}

func (UnimplementedSignerServer) Sign(context.Context, *SignReq) (*SignRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedSignerServer) Gen(context.Context, *GenReq) (*GenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gen not implemented")
}
func (UnimplementedSignerServer) Migration(context.Context, *MigrationReq) (*MigrationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Migration not implemented")
}
func (UnimplementedSignerServer) mustEmbedUnimplementedSignerServer() {}
func (UnimplementedSignerServer) testEmbeddedByValue()                {}

// UnsafeSignerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignerServer will
// result in compilation errors.
type UnsafeSignerServer interface {
	mustEmbedUnimplementedSignerServer()
}

func RegisterSignerServer(s grpc.ServiceRegistrar, srv SignerServer) {
	// If the following call pancis, it indicates UnimplementedSignerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Signer_ServiceDesc, srv)
}

func _Signer_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_Sign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).Sign(ctx, req.(*SignReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signer_Gen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).Gen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_Gen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).Gen(ctx, req.(*GenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signer_Migration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).Migration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_Migration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).Migration(ctx, req.(*MigrationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Signer_ServiceDesc is the grpc.ServiceDesc for Signer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Signer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "signer.Signer",
	HandlerType: (*SignerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _Signer_Sign_Handler,
		},
		{
			MethodName: "Gen",
			Handler:    _Signer_Gen_Handler,
		},
		{
			MethodName: "Migration",
			Handler:    _Signer_Migration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signer.proto",
}