// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: i18n/proto/i18n.proto

package truffle_i18n

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

// I18NClient is the client API for I18N service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type I18NClient interface {
	T(ctx context.Context, in *I18NRequest, opts ...grpc.CallOption) (*I18NResponse, error)
	// translate stream
	TS(ctx context.Context, opts ...grpc.CallOption) (I18N_TSClient, error)
}

type i18NClient struct {
	cc grpc.ClientConnInterface
}

func NewI18NClient(cc grpc.ClientConnInterface) I18NClient {
	return &i18NClient{cc}
}

func (c *i18NClient) T(ctx context.Context, in *I18NRequest, opts ...grpc.CallOption) (*I18NResponse, error) {
	out := new(I18NResponse)
	err := c.cc.Invoke(ctx, "/truffle.I18N/T", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *i18NClient) TS(ctx context.Context, opts ...grpc.CallOption) (I18N_TSClient, error) {
	stream, err := c.cc.NewStream(ctx, &I18N_ServiceDesc.Streams[0], "/truffle.I18N/TS", opts...)
	if err != nil {
		return nil, err
	}
	x := &i18NTSClient{stream}
	return x, nil
}

type I18N_TSClient interface {
	Send(*I18NRequest) error
	Recv() (*I18NResponse, error)
	grpc.ClientStream
}

type i18NTSClient struct {
	grpc.ClientStream
}

func (x *i18NTSClient) Send(m *I18NRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *i18NTSClient) Recv() (*I18NResponse, error) {
	m := new(I18NResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// I18NServer is the server API for I18N service.
// All implementations must embed UnimplementedI18NServer
// for forward compatibility
type I18NServer interface {
	T(context.Context, *I18NRequest) (*I18NResponse, error)
	// translate stream
	TS(I18N_TSServer) error
	mustEmbedUnimplementedI18NServer()
}

// UnimplementedI18NServer must be embedded to have forward compatible implementations.
type UnimplementedI18NServer struct {
}

func (UnimplementedI18NServer) T(context.Context, *I18NRequest) (*I18NResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method T not implemented")
}
func (UnimplementedI18NServer) TS(I18N_TSServer) error {
	return status.Errorf(codes.Unimplemented, "method TS not implemented")
}
func (UnimplementedI18NServer) mustEmbedUnimplementedI18NServer() {}

// UnsafeI18NServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to I18NServer will
// result in compilation errors.
type UnsafeI18NServer interface {
	mustEmbedUnimplementedI18NServer()
}

func RegisterI18NServer(s grpc.ServiceRegistrar, srv I18NServer) {
	s.RegisterService(&I18N_ServiceDesc, srv)
}

func _I18N_T_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(I18NRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(I18NServer).T(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/truffle.I18N/T",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(I18NServer).T(ctx, req.(*I18NRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _I18N_TS_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(I18NServer).TS(&i18NTSServer{stream})
}

type I18N_TSServer interface {
	Send(*I18NResponse) error
	Recv() (*I18NRequest, error)
	grpc.ServerStream
}

type i18NTSServer struct {
	grpc.ServerStream
}

func (x *i18NTSServer) Send(m *I18NResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *i18NTSServer) Recv() (*I18NRequest, error) {
	m := new(I18NRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// I18N_ServiceDesc is the grpc.ServiceDesc for I18N service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var I18N_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "truffle.I18N",
	HandlerType: (*I18NServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "T",
			Handler:    _I18N_T_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TS",
			Handler:       _I18N_TS_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "i18n/proto/i18n.proto",
}