// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: service.proto

package filemsg

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

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	GetFileNames(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FileListResponse, error)
	GetFileInfo(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileInfoResponse, error)
	GetFileData(ctx context.Context, opts ...grpc.CallOption) (FileService_GetFileDataClient, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) GetFileNames(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FileListResponse, error) {
	out := new(FileListResponse)
	err := c.cc.Invoke(ctx, "/filemsg.FileService/GetFileNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileInfo(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileInfoResponse, error) {
	out := new(FileInfoResponse)
	err := c.cc.Invoke(ctx, "/filemsg.FileService/GetFileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileData(ctx context.Context, opts ...grpc.CallOption) (FileService_GetFileDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], "/filemsg.FileService/GetFileData", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceGetFileDataClient{stream}
	return x, nil
}

type FileService_GetFileDataClient interface {
	Send(*FileRequest) error
	Recv() (*FileDataResponse, error)
	grpc.ClientStream
}

type fileServiceGetFileDataClient struct {
	grpc.ClientStream
}

func (x *fileServiceGetFileDataClient) Send(m *FileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceGetFileDataClient) Recv() (*FileDataResponse, error) {
	m := new(FileDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	GetFileNames(context.Context, *emptypb.Empty) (*FileListResponse, error)
	GetFileInfo(context.Context, *FileRequest) (*FileInfoResponse, error)
	GetFileData(FileService_GetFileDataServer) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) GetFileNames(context.Context, *emptypb.Empty) (*FileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileNames not implemented")
}
func (UnimplementedFileServiceServer) GetFileInfo(context.Context, *FileRequest) (*FileInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileInfo not implemented")
}
func (UnimplementedFileServiceServer) GetFileData(FileService_GetFileDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFileData not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_GetFileNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filemsg.FileService/GetFileNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileNames(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filemsg.FileService/GetFileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileInfo(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).GetFileData(&fileServiceGetFileDataServer{stream})
}

type FileService_GetFileDataServer interface {
	Send(*FileDataResponse) error
	Recv() (*FileRequest, error)
	grpc.ServerStream
}

type fileServiceGetFileDataServer struct {
	grpc.ServerStream
}

func (x *fileServiceGetFileDataServer) Send(m *FileDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceGetFileDataServer) Recv() (*FileRequest, error) {
	m := new(FileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filemsg.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileNames",
			Handler:    _FileService_GetFileNames_Handler,
		},
		{
			MethodName: "GetFileInfo",
			Handler:    _FileService_GetFileInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFileData",
			Handler:       _FileService_GetFileData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
