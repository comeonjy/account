// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error)
	SmsLogin(ctx context.Context, in *SmsLoginReq, opts ...grpc.CallOption) (*SmsLoginResp, error)
	MiniLogin(ctx context.Context, in *MiniLoginReq, opts ...grpc.CallOption) (*MiniLoginResp, error)
	GetByID(ctx context.Context, in *GetByIDReq, opts ...grpc.CallOption) (*GetByIDResp, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/task_system.scheduler.v1.Account/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) SmsLogin(ctx context.Context, in *SmsLoginReq, opts ...grpc.CallOption) (*SmsLoginResp, error) {
	out := new(SmsLoginResp)
	err := c.cc.Invoke(ctx, "/task_system.scheduler.v1.Account/SmsLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) MiniLogin(ctx context.Context, in *MiniLoginReq, opts ...grpc.CallOption) (*MiniLoginResp, error) {
	out := new(MiniLoginResp)
	err := c.cc.Invoke(ctx, "/task_system.scheduler.v1.Account/MiniLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetByID(ctx context.Context, in *GetByIDReq, opts ...grpc.CallOption) (*GetByIDResp, error) {
	out := new(GetByIDResp)
	err := c.cc.Invoke(ctx, "/task_system.scheduler.v1.Account/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	Ping(context.Context, *Empty) (*Result, error)
	SmsLogin(context.Context, *SmsLoginReq) (*SmsLoginResp, error)
	MiniLogin(context.Context, *MiniLoginReq) (*MiniLoginResp, error)
	GetByID(context.Context, *GetByIDReq) (*GetByIDResp, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) Ping(context.Context, *Empty) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAccountServer) SmsLogin(context.Context, *SmsLoginReq) (*SmsLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmsLogin not implemented")
}
func (UnimplementedAccountServer) MiniLogin(context.Context, *MiniLoginReq) (*MiniLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MiniLogin not implemented")
}
func (UnimplementedAccountServer) GetByID(context.Context, *GetByIDReq) (*GetByIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task_system.scheduler.v1.Account/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_SmsLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SmsLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task_system.scheduler.v1.Account/SmsLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SmsLogin(ctx, req.(*SmsLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_MiniLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MiniLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).MiniLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task_system.scheduler.v1.Account/MiniLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).MiniLogin(ctx, req.(*MiniLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task_system.scheduler.v1.Account/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetByID(ctx, req.(*GetByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task_system.scheduler.v1.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Account_Ping_Handler,
		},
		{
			MethodName: "SmsLogin",
			Handler:    _Account_SmsLogin_Handler,
		},
		{
			MethodName: "MiniLogin",
			Handler:    _Account_MiniLogin_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Account_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/account.proto",
}
