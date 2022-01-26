// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OauthClient is the client API for Oauth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OauthClient interface {
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	Authenticate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
}

type oauthClient struct {
	cc grpc.ClientConnInterface
}

func NewOauthClient(cc grpc.ClientConnInterface) OauthClient {
	return &oauthClient{cc}
}

func (c *oauthClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.oauth.v1.Oauth/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthClient) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.oauth.v1.Oauth/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthClient) Authenticate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.oauth.v1.Oauth/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.oauth.v1.Oauth/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OauthServer is the server API for Oauth service.
// All implementations must embed UnimplementedOauthServer
// for forward compatibility
type OauthServer interface {
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	Token(context.Context, *TokenRequest) (*TokenResponse, error)
	Authenticate(context.Context, *emptypb.Empty) (*AuthenticateResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	mustEmbedUnimplementedOauthServer()
}

// UnimplementedOauthServer must be embedded to have forward compatible implementations.
type UnimplementedOauthServer struct {
}

func (UnimplementedOauthServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedOauthServer) Token(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedOauthServer) Authenticate(context.Context, *emptypb.Empty) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedOauthServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedOauthServer) mustEmbedUnimplementedOauthServer() {}

// UnsafeOauthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OauthServer will
// result in compilation errors.
type UnsafeOauthServer interface {
	mustEmbedUnimplementedOauthServer()
}

func RegisterOauthServer(s grpc.ServiceRegistrar, srv OauthServer) {
	s.RegisterService(&Oauth_ServiceDesc, srv)
}

func _Oauth_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.oauth.v1.Oauth/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.oauth.v1.Oauth/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).Token(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.oauth.v1.Oauth/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).Authenticate(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.oauth.v1.Oauth/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Oauth_ServiceDesc is the grpc.ServiceDesc for Oauth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oauth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.tkeel.security.api.oauth.v1.Oauth",
	HandlerType: (*OauthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _Oauth_Authorize_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Oauth_Token_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _Oauth_Authenticate_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Oauth_ResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/security_oauth/v1/oauth.proto",
}
