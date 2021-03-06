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
	// 发送验证码
	SendVerificationCode(ctx context.Context, in *SendVerificationCodeReq, opts ...grpc.CallOption) (*Empty, error)
	// 鉴权
	Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error)
	// 登录
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	// 获取小程序授权登录二维码
	GetMiniQRCode(ctx context.Context, in *GetMiniQRCodeReq, opts ...grpc.CallOption) (*GetMiniQRCodeResp, error)
	// 小程序登录
	MiniLogin(ctx context.Context, in *MiniLoginReq, opts ...grpc.CallOption) (*MiniLoginResp, error)
	// 更新用户信息
	UpdatesUser(ctx context.Context, in *UpdatesUserReq, opts ...grpc.CallOption) (*Empty, error)
	// 获取用户信息
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

func (c *accountClient) SendVerificationCode(ctx context.Context, in *SendVerificationCodeReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/task_system.scheduler.v1.Account/SendVerificationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error) {
	out := new(AuthResp)
	err := c.cc.Invoke(ctx, "/task_system.scheduler.v1.Account/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/task_system.scheduler.v1.Account/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetMiniQRCode(ctx context.Context, in *GetMiniQRCodeReq, opts ...grpc.CallOption) (*GetMiniQRCodeResp, error) {
	out := new(GetMiniQRCodeResp)
	err := c.cc.Invoke(ctx, "/task_system.scheduler.v1.Account/GetMiniQRCode", in, out, opts...)
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

func (c *accountClient) UpdatesUser(ctx context.Context, in *UpdatesUserReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/task_system.scheduler.v1.Account/UpdatesUser", in, out, opts...)
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
	// 发送验证码
	SendVerificationCode(context.Context, *SendVerificationCodeReq) (*Empty, error)
	// 鉴权
	Auth(context.Context, *AuthReq) (*AuthResp, error)
	// 登录
	Login(context.Context, *LoginReq) (*LoginResp, error)
	// 获取小程序授权登录二维码
	GetMiniQRCode(context.Context, *GetMiniQRCodeReq) (*GetMiniQRCodeResp, error)
	// 小程序登录
	MiniLogin(context.Context, *MiniLoginReq) (*MiniLoginResp, error)
	// 更新用户信息
	UpdatesUser(context.Context, *UpdatesUserReq) (*Empty, error)
	// 获取用户信息
	GetByID(context.Context, *GetByIDReq) (*GetByIDResp, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) Ping(context.Context, *Empty) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAccountServer) SendVerificationCode(context.Context, *SendVerificationCodeReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerificationCode not implemented")
}
func (UnimplementedAccountServer) Auth(context.Context, *AuthReq) (*AuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedAccountServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAccountServer) GetMiniQRCode(context.Context, *GetMiniQRCodeReq) (*GetMiniQRCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMiniQRCode not implemented")
}
func (UnimplementedAccountServer) MiniLogin(context.Context, *MiniLoginReq) (*MiniLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MiniLogin not implemented")
}
func (UnimplementedAccountServer) UpdatesUser(context.Context, *UpdatesUserReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatesUser not implemented")
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

func _Account_SendVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerificationCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SendVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task_system.scheduler.v1.Account/SendVerificationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SendVerificationCode(ctx, req.(*SendVerificationCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task_system.scheduler.v1.Account/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Auth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task_system.scheduler.v1.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetMiniQRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMiniQRCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetMiniQRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task_system.scheduler.v1.Account/GetMiniQRCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetMiniQRCode(ctx, req.(*GetMiniQRCodeReq))
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

func _Account_UpdatesUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatesUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdatesUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task_system.scheduler.v1.Account/UpdatesUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdatesUser(ctx, req.(*UpdatesUserReq))
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
			MethodName: "SendVerificationCode",
			Handler:    _Account_SendVerificationCode_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Account_Auth_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
		{
			MethodName: "GetMiniQRCode",
			Handler:    _Account_GetMiniQRCode_Handler,
		},
		{
			MethodName: "MiniLogin",
			Handler:    _Account_MiniLogin_Handler,
		},
		{
			MethodName: "UpdatesUser",
			Handler:    _Account_UpdatesUser_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Account_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/account.proto",
}
