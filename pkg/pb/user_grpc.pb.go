// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/pb/user.proto

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

const (
	UserService_FindUserByEmail_FullMethodName = "/pb.UserService/FindUserByEmail"
	UserService_SaveUser_FullMethodName        = "/pb.UserService/SaveUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	FindUserByEmail(ctx context.Context, in *FindUserByEmailRequest, opts ...grpc.CallOption) (*FindUserResponse, error)
	SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindUserByEmail(ctx context.Context, in *FindUserByEmailRequest, opts ...grpc.CallOption) (*FindUserResponse, error) {
	out := new(FindUserResponse)
	err := c.cc.Invoke(ctx, UserService_FindUserByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error) {
	out := new(SaveUserResponse)
	err := c.cc.Invoke(ctx, UserService_SaveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	FindUserByEmail(context.Context, *FindUserByEmailRequest) (*FindUserResponse, error)
	SaveUser(context.Context, *SaveUserRequest) (*SaveUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) FindUserByEmail(context.Context, *FindUserByEmailRequest) (*FindUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByEmail not implemented")
}
func (UnimplementedUserServiceServer) SaveUser(context.Context, *SaveUserRequest) (*SaveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_FindUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindUserByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindUserByEmail(ctx, req.(*FindUserByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SaveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SaveUser(ctx, req.(*SaveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUserByEmail",
			Handler:    _UserService_FindUserByEmail_Handler,
		},
		{
			MethodName: "SaveUser",
			Handler:    _UserService_SaveUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/user.proto",
}
