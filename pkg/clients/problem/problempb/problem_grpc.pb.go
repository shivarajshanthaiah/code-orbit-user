// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: problem.proto

package __

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

// ProblemServiceClient is the client API for ProblemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemServiceClient interface {
	GetAllProblems(ctx context.Context, in *ProbNoParam, opts ...grpc.CallOption) (*ProblemList, error)
	GetProblemWithTestCases(ctx context.Context, in *ProblemId, opts ...grpc.CallOption) (*GetProblemResponse, error)
}

type problemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemServiceClient(cc grpc.ClientConnInterface) ProblemServiceClient {
	return &problemServiceClient{cc}
}

func (c *problemServiceClient) GetAllProblems(ctx context.Context, in *ProbNoParam, opts ...grpc.CallOption) (*ProblemList, error) {
	out := new(ProblemList)
	err := c.cc.Invoke(ctx, "/pb.ProblemService/GetAllProblems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) GetProblemWithTestCases(ctx context.Context, in *ProblemId, opts ...grpc.CallOption) (*GetProblemResponse, error) {
	out := new(GetProblemResponse)
	err := c.cc.Invoke(ctx, "/pb.ProblemService/GetProblemWithTestCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemServiceServer is the server API for ProblemService service.
// All implementations must embed UnimplementedProblemServiceServer
// for forward compatibility
type ProblemServiceServer interface {
	GetAllProblems(context.Context, *ProbNoParam) (*ProblemList, error)
	GetProblemWithTestCases(context.Context, *ProblemId) (*GetProblemResponse, error)
	mustEmbedUnimplementedProblemServiceServer()
}

// UnimplementedProblemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProblemServiceServer struct {
}

func (UnimplementedProblemServiceServer) GetAllProblems(context.Context, *ProbNoParam) (*ProblemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProblems not implemented")
}
func (UnimplementedProblemServiceServer) GetProblemWithTestCases(context.Context, *ProblemId) (*GetProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemWithTestCases not implemented")
}
func (UnimplementedProblemServiceServer) mustEmbedUnimplementedProblemServiceServer() {}

// UnsafeProblemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemServiceServer will
// result in compilation errors.
type UnsafeProblemServiceServer interface {
	mustEmbedUnimplementedProblemServiceServer()
}

func RegisterProblemServiceServer(s grpc.ServiceRegistrar, srv ProblemServiceServer) {
	s.RegisterService(&ProblemService_ServiceDesc, srv)
}

func _ProblemService_GetAllProblems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetAllProblems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProblemService/GetAllProblems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetAllProblems(ctx, req.(*ProbNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_GetProblemWithTestCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProblemId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetProblemWithTestCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProblemService/GetProblemWithTestCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetProblemWithTestCases(ctx, req.(*ProblemId))
	}
	return interceptor(ctx, in, info, handler)
}

// ProblemService_ServiceDesc is the grpc.ServiceDesc for ProblemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProblemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProblemService",
	HandlerType: (*ProblemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllProblems",
			Handler:    _ProblemService_GetAllProblems_Handler,
		},
		{
			MethodName: "GetProblemWithTestCases",
			Handler:    _ProblemService_GetProblemWithTestCases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "problem.proto",
}
