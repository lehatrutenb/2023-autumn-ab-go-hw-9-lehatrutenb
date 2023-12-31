// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: validation_service.proto

package validation

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
	ValidationService_CreateUser_FullMethodName = "/validation.ValidationService/CreateUser"
)

// ValidationServiceClient is the client API for ValidationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValidationServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type validationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidationServiceClient(cc grpc.ClientConnInterface) ValidationServiceClient {
	return &validationServiceClient{cc}
}

func (c *validationServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, ValidationService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidationServiceServer is the server API for ValidationService service.
// All implementations must embed UnimplementedValidationServiceServer
// for forward compatibility
type ValidationServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedValidationServiceServer()
}

// UnimplementedValidationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedValidationServiceServer struct {
}

func (UnimplementedValidationServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedValidationServiceServer) mustEmbedUnimplementedValidationServiceServer() {}

// UnsafeValidationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValidationServiceServer will
// result in compilation errors.
type UnsafeValidationServiceServer interface {
	mustEmbedUnimplementedValidationServiceServer()
}

func RegisterValidationServiceServer(s grpc.ServiceRegistrar, srv ValidationServiceServer) {
	s.RegisterService(&ValidationService_ServiceDesc, srv)
}

func _ValidationService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidationServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValidationService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidationServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ValidationService_ServiceDesc is the grpc.ServiceDesc for ValidationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValidationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "validation.ValidationService",
	HandlerType: (*ValidationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _ValidationService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "validation_service.proto",
}
