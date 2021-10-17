// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user_grpc

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

// UserSignUpClient is the client API for UserSignUp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSignUpClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type userSignUpClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSignUpClient(cc grpc.ClientConnInterface) UserSignUpClient {
	return &userSignUpClient{cc}
}

func (c *userSignUpClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/UserSignUp/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSignUpClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/UserSignUp/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSignUpServer is the server API for UserSignUp service.
// All implementations must embed UnimplementedUserSignUpServer
// for forward compatibility
type UserSignUpServer interface {
	SignUp(context.Context, *SignUpRequest) (*LoginResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedUserSignUpServer()
}

// UnimplementedUserSignUpServer must be embedded to have forward compatible implementations.
type UnimplementedUserSignUpServer struct {
}

func (UnimplementedUserSignUpServer) SignUp(context.Context, *SignUpRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedUserSignUpServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserSignUpServer) mustEmbedUnimplementedUserSignUpServer() {}

// UnsafeUserSignUpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSignUpServer will
// result in compilation errors.
type UnsafeUserSignUpServer interface {
	mustEmbedUnimplementedUserSignUpServer()
}

func RegisterUserSignUpServer(s grpc.ServiceRegistrar, srv UserSignUpServer) {
	s.RegisterService(&UserSignUp_ServiceDesc, srv)
}

func _UserSignUp_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSignUpServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserSignUp/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSignUpServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSignUp_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSignUpServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserSignUp/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSignUpServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSignUp_ServiceDesc is the grpc.ServiceDesc for UserSignUp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSignUp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserSignUp",
	HandlerType: (*UserSignUpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _UserSignUp_SignUp_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserSignUp_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sign_up.proto",
}
