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

// RepoClient is the client API for Repo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepoClient interface {
	CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*DeleteRepoResponse, error)
	ListRepo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRepoResponse, error)
	ListRepoInstaller(ctx context.Context, in *ListRepoInstallerRequest, opts ...grpc.CallOption) (*ListRepoInstallerResponse, error)
	GetRepoInstaller(ctx context.Context, in *GetRepoInstallerRequest, opts ...grpc.CallOption) (*GetRepoInstallerResponse, error)
}

type repoClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoClient(cc grpc.ClientConnInterface) RepoClient {
	return &repoClient{cc}
}

func (c *repoClient) CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.repo.v1.Repo/CreateRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*DeleteRepoResponse, error) {
	out := new(DeleteRepoResponse)
	err := c.cc.Invoke(ctx, "/api.repo.v1.Repo/DeleteRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) ListRepo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRepoResponse, error) {
	out := new(ListRepoResponse)
	err := c.cc.Invoke(ctx, "/api.repo.v1.Repo/ListRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) ListRepoInstaller(ctx context.Context, in *ListRepoInstallerRequest, opts ...grpc.CallOption) (*ListRepoInstallerResponse, error) {
	out := new(ListRepoInstallerResponse)
	err := c.cc.Invoke(ctx, "/api.repo.v1.Repo/ListRepoInstaller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) GetRepoInstaller(ctx context.Context, in *GetRepoInstallerRequest, opts ...grpc.CallOption) (*GetRepoInstallerResponse, error) {
	out := new(GetRepoInstallerResponse)
	err := c.cc.Invoke(ctx, "/api.repo.v1.Repo/GetRepoInstaller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoServer is the server API for Repo service.
// All implementations must embed UnimplementedRepoServer
// for forward compatibility
type RepoServer interface {
	CreateRepo(context.Context, *CreateRepoRequest) (*emptypb.Empty, error)
	DeleteRepo(context.Context, *DeleteRepoRequest) (*DeleteRepoResponse, error)
	ListRepo(context.Context, *emptypb.Empty) (*ListRepoResponse, error)
	ListRepoInstaller(context.Context, *ListRepoInstallerRequest) (*ListRepoInstallerResponse, error)
	GetRepoInstaller(context.Context, *GetRepoInstallerRequest) (*GetRepoInstallerResponse, error)
	mustEmbedUnimplementedRepoServer()
}

// UnimplementedRepoServer must be embedded to have forward compatible implementations.
type UnimplementedRepoServer struct {
}

func (UnimplementedRepoServer) CreateRepo(context.Context, *CreateRepoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepo not implemented")
}
func (UnimplementedRepoServer) DeleteRepo(context.Context, *DeleteRepoRequest) (*DeleteRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepo not implemented")
}
func (UnimplementedRepoServer) ListRepo(context.Context, *emptypb.Empty) (*ListRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepo not implemented")
}
func (UnimplementedRepoServer) ListRepoInstaller(context.Context, *ListRepoInstallerRequest) (*ListRepoInstallerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepoInstaller not implemented")
}
func (UnimplementedRepoServer) GetRepoInstaller(context.Context, *GetRepoInstallerRequest) (*GetRepoInstallerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepoInstaller not implemented")
}
func (UnimplementedRepoServer) mustEmbedUnimplementedRepoServer() {}

// UnsafeRepoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepoServer will
// result in compilation errors.
type UnsafeRepoServer interface {
	mustEmbedUnimplementedRepoServer()
}

func RegisterRepoServer(s grpc.ServiceRegistrar, srv RepoServer) {
	s.RegisterService(&Repo_ServiceDesc, srv)
}

func _Repo_CreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).CreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.repo.v1.Repo/CreateRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).CreateRepo(ctx, req.(*CreateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_DeleteRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).DeleteRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.repo.v1.Repo/DeleteRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).DeleteRepo(ctx, req.(*DeleteRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_ListRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).ListRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.repo.v1.Repo/ListRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).ListRepo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_ListRepoInstaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepoInstallerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).ListRepoInstaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.repo.v1.Repo/ListRepoInstaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).ListRepoInstaller(ctx, req.(*ListRepoInstallerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_GetRepoInstaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepoInstallerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).GetRepoInstaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.repo.v1.Repo/GetRepoInstaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).GetRepoInstaller(ctx, req.(*GetRepoInstallerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Repo_ServiceDesc is the grpc.ServiceDesc for Repo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.repo.v1.Repo",
	HandlerType: (*RepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepo",
			Handler:    _Repo_CreateRepo_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _Repo_DeleteRepo_Handler,
		},
		{
			MethodName: "ListRepo",
			Handler:    _Repo_ListRepo_Handler,
		},
		{
			MethodName: "ListRepoInstaller",
			Handler:    _Repo_ListRepoInstaller_Handler,
		},
		{
			MethodName: "GetRepoInstaller",
			Handler:    _Repo_GetRepoInstaller_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/repo/v1/repo.proto",
}