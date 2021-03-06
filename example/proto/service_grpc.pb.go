// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package file

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

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadFileClient, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadService_ServiceDesc.Streams[0], "/UploadService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadFileClient{stream}
	return x, nil
}

type UploadService_UploadFileClient interface {
	Send(*File) error
	CloseAndRecv() (*FileResponse, error)
	grpc.ClientStream
}

type uploadServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadFileClient) Send(m *File) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadFileClient) CloseAndRecv() (*FileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServiceServer is the server API for UploadService service.
// All implementations should embed UnimplementedUploadServiceServer
// for forward compatibility
type UploadServiceServer interface {
	UploadFile(UploadService_UploadFileServer) error
}

// UnimplementedUploadServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUploadServiceServer struct {
}

func (UnimplementedUploadServiceServer) UploadFile(UploadService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}

// UnsafeUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServiceServer will
// result in compilation errors.
type UnsafeUploadServiceServer interface {
	mustEmbedUnimplementedUploadServiceServer()
}

func RegisterUploadServiceServer(s grpc.ServiceRegistrar, srv UploadServiceServer) {
	s.RegisterService(&UploadService_ServiceDesc, srv)
}

func _UploadService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).UploadFile(&uploadServiceUploadFileServer{stream})
}

type UploadService_UploadFileServer interface {
	SendAndClose(*FileResponse) error
	Recv() (*File, error)
	grpc.ServerStream
}

type uploadServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadFileServer) SendAndClose(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadFileServer) Recv() (*File, error) {
	m := new(File)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadService_ServiceDesc is the grpc.ServiceDesc for UploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _UploadService_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
