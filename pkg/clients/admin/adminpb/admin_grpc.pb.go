// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: admin.proto

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	GetAllPlans(ctx context.Context, in *AdNoParam, opts ...grpc.CallOption) (*AdPlanList, error)
	GetSubscriptionByID(ctx context.Context, in *SubscriptionID, opts ...grpc.CallOption) (*AdSubscription, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) GetAllPlans(ctx context.Context, in *AdNoParam, opts ...grpc.CallOption) (*AdPlanList, error) {
	out := new(AdPlanList)
	err := c.cc.Invoke(ctx, "/pb.AdminService/GetAllPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetSubscriptionByID(ctx context.Context, in *SubscriptionID, opts ...grpc.CallOption) (*AdSubscription, error) {
	out := new(AdSubscription)
	err := c.cc.Invoke(ctx, "/pb.AdminService/GetSubscriptionByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	GetAllPlans(context.Context, *AdNoParam) (*AdPlanList, error)
	GetSubscriptionByID(context.Context, *SubscriptionID) (*AdSubscription, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) GetAllPlans(context.Context, *AdNoParam) (*AdPlanList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPlans not implemented")
}
func (UnimplementedAdminServiceServer) GetSubscriptionByID(context.Context, *SubscriptionID) (*AdSubscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionByID not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_GetAllPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetAllPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/GetAllPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetAllPlans(ctx, req.(*AdNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetSubscriptionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetSubscriptionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/GetSubscriptionByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetSubscriptionByID(ctx, req.(*SubscriptionID))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPlans",
			Handler:    _AdminService_GetAllPlans_Handler,
		},
		{
			MethodName: "GetSubscriptionByID",
			Handler:    _AdminService_GetSubscriptionByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
