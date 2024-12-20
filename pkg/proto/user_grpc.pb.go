// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserSignup(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*Response, error)
	VerifyUser(ctx context.Context, in *OTP, opts ...grpc.CallOption) (*Response, error)
	UserLogin(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Response, error)
	ViewProfile(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Profile, error)
	EditProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Profile, error)
	ChangePassword(ctx context.Context, in *Password, opts ...grpc.CallOption) (*Response, error)
	BlockUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error)
	UnBlockUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error)
	GetAllUsers(ctx context.Context, in *UserNoParam, opts ...grpc.CallOption) (*UserList, error)
	UserGetAllProblems(ctx context.Context, in *UserNoParam, opts ...grpc.CallOption) (*UserProblemList, error)
	UserGetProblemWithTestCases(ctx context.Context, in *UserProblemId, opts ...grpc.CallOption) (*UserTestcaseResponse, error)
	SubmitCode(ctx context.Context, in *UserSubmissionRequest, opts ...grpc.CallOption) (*UserSubmissionResponse, error)
	GetUserStats(ctx context.Context, in *ID, opts ...grpc.CallOption) (*UserStatsResponse, error)
	AddSubscriptionPlan(ctx context.Context, in *USubscription, opts ...grpc.CallOption) (*Response, error)
	GetAllPlans(ctx context.Context, in *UserNoParam, opts ...grpc.CallOption) (*UPlanList, error)
	UpdatePlan(ctx context.Context, in *USubscription, opts ...grpc.CallOption) (*USubscription, error)
	GenerateInvoice(ctx context.Context, in *InvoiceRequest, opts ...grpc.CallOption) (*Response, error)
	GetUserInvoices(ctx context.Context, in *ID, opts ...grpc.CallOption) (*InvoiceList, error)
	MakePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	PaymentSuccess(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*ConfirmResponse, error)
	GetUserProfileStats(ctx context.Context, in *UserStatsRequest, opts ...grpc.CallOption) (*UserStatsProfileResponse, error)
	GetSubscriptionStats(ctx context.Context, in *SubscriptionStatsRequest, opts ...grpc.CallOption) (*SubscriptionStatsResponse, error)
	UserGetProblemStats(ctx context.Context, in *UProblemStatsRequest, opts ...grpc.CallOption) (*UProblemStatsResponse, error)
	UserGetLeaderboardStats(ctx context.Context, in *ULeaderboardRequest, opts ...grpc.CallOption) (*ULeaderboardResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserSignup(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/UserSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyUser(ctx context.Context, in *OTP, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/VerifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ViewProfile(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/pb.UserService/ViewProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/pb.UserService/EditProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *Password, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BlockUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UnBlockUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/UnBlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *UserNoParam, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserGetAllProblems(ctx context.Context, in *UserNoParam, opts ...grpc.CallOption) (*UserProblemList, error) {
	out := new(UserProblemList)
	err := c.cc.Invoke(ctx, "/pb.UserService/UserGetAllProblems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserGetProblemWithTestCases(ctx context.Context, in *UserProblemId, opts ...grpc.CallOption) (*UserTestcaseResponse, error) {
	out := new(UserTestcaseResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/UserGetProblemWithTestCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SubmitCode(ctx context.Context, in *UserSubmissionRequest, opts ...grpc.CallOption) (*UserSubmissionResponse, error) {
	out := new(UserSubmissionResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/SubmitCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserStats(ctx context.Context, in *ID, opts ...grpc.CallOption) (*UserStatsResponse, error) {
	out := new(UserStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetUserStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddSubscriptionPlan(ctx context.Context, in *USubscription, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/AddSubscriptionPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllPlans(ctx context.Context, in *UserNoParam, opts ...grpc.CallOption) (*UPlanList, error) {
	out := new(UPlanList)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetAllPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePlan(ctx context.Context, in *USubscription, opts ...grpc.CallOption) (*USubscription, error) {
	out := new(USubscription)
	err := c.cc.Invoke(ctx, "/pb.UserService/UpdatePlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GenerateInvoice(ctx context.Context, in *InvoiceRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/GenerateInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInvoices(ctx context.Context, in *ID, opts ...grpc.CallOption) (*InvoiceList, error) {
	out := new(InvoiceList)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetUserInvoices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) MakePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/MakePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PaymentSuccess(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/PaymentSuccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserProfileStats(ctx context.Context, in *UserStatsRequest, opts ...grpc.CallOption) (*UserStatsProfileResponse, error) {
	out := new(UserStatsProfileResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetUserProfileStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetSubscriptionStats(ctx context.Context, in *SubscriptionStatsRequest, opts ...grpc.CallOption) (*SubscriptionStatsResponse, error) {
	out := new(SubscriptionStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetSubscriptionStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserGetProblemStats(ctx context.Context, in *UProblemStatsRequest, opts ...grpc.CallOption) (*UProblemStatsResponse, error) {
	out := new(UProblemStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/UserGetProblemStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserGetLeaderboardStats(ctx context.Context, in *ULeaderboardRequest, opts ...grpc.CallOption) (*ULeaderboardResponse, error) {
	out := new(ULeaderboardResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/UserGetLeaderboardStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserSignup(context.Context, *Signup) (*Response, error)
	VerifyUser(context.Context, *OTP) (*Response, error)
	UserLogin(context.Context, *Login) (*Response, error)
	ViewProfile(context.Context, *ID) (*Profile, error)
	EditProfile(context.Context, *Profile) (*Profile, error)
	ChangePassword(context.Context, *Password) (*Response, error)
	BlockUser(context.Context, *ID) (*Response, error)
	UnBlockUser(context.Context, *ID) (*Response, error)
	GetAllUsers(context.Context, *UserNoParam) (*UserList, error)
	UserGetAllProblems(context.Context, *UserNoParam) (*UserProblemList, error)
	UserGetProblemWithTestCases(context.Context, *UserProblemId) (*UserTestcaseResponse, error)
	SubmitCode(context.Context, *UserSubmissionRequest) (*UserSubmissionResponse, error)
	GetUserStats(context.Context, *ID) (*UserStatsResponse, error)
	AddSubscriptionPlan(context.Context, *USubscription) (*Response, error)
	GetAllPlans(context.Context, *UserNoParam) (*UPlanList, error)
	UpdatePlan(context.Context, *USubscription) (*USubscription, error)
	GenerateInvoice(context.Context, *InvoiceRequest) (*Response, error)
	GetUserInvoices(context.Context, *ID) (*InvoiceList, error)
	MakePayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	PaymentSuccess(context.Context, *ConfirmRequest) (*ConfirmResponse, error)
	GetUserProfileStats(context.Context, *UserStatsRequest) (*UserStatsProfileResponse, error)
	GetSubscriptionStats(context.Context, *SubscriptionStatsRequest) (*SubscriptionStatsResponse, error)
	UserGetProblemStats(context.Context, *UProblemStatsRequest) (*UProblemStatsResponse, error)
	UserGetLeaderboardStats(context.Context, *ULeaderboardRequest) (*ULeaderboardResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserSignup(context.Context, *Signup) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignup not implemented")
}
func (UnimplementedUserServiceServer) VerifyUser(context.Context, *OTP) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}
func (UnimplementedUserServiceServer) UserLogin(context.Context, *Login) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) ViewProfile(context.Context, *ID) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewProfile not implemented")
}
func (UnimplementedUserServiceServer) EditProfile(context.Context, *Profile) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProfile not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *Password) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceServer) BlockUser(context.Context, *ID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedUserServiceServer) UnBlockUser(context.Context, *ID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockUser not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *UserNoParam) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) UserGetAllProblems(context.Context, *UserNoParam) (*UserProblemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetAllProblems not implemented")
}
func (UnimplementedUserServiceServer) UserGetProblemWithTestCases(context.Context, *UserProblemId) (*UserTestcaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetProblemWithTestCases not implemented")
}
func (UnimplementedUserServiceServer) SubmitCode(context.Context, *UserSubmissionRequest) (*UserSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCode not implemented")
}
func (UnimplementedUserServiceServer) GetUserStats(context.Context, *ID) (*UserStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStats not implemented")
}
func (UnimplementedUserServiceServer) AddSubscriptionPlan(context.Context, *USubscription) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscriptionPlan not implemented")
}
func (UnimplementedUserServiceServer) GetAllPlans(context.Context, *UserNoParam) (*UPlanList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPlans not implemented")
}
func (UnimplementedUserServiceServer) UpdatePlan(context.Context, *USubscription) (*USubscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlan not implemented")
}
func (UnimplementedUserServiceServer) GenerateInvoice(context.Context, *InvoiceRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInvoice not implemented")
}
func (UnimplementedUserServiceServer) GetUserInvoices(context.Context, *ID) (*InvoiceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInvoices not implemented")
}
func (UnimplementedUserServiceServer) MakePayment(context.Context, *PaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePayment not implemented")
}
func (UnimplementedUserServiceServer) PaymentSuccess(context.Context, *ConfirmRequest) (*ConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentSuccess not implemented")
}
func (UnimplementedUserServiceServer) GetUserProfileStats(context.Context, *UserStatsRequest) (*UserStatsProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfileStats not implemented")
}
func (UnimplementedUserServiceServer) GetSubscriptionStats(context.Context, *SubscriptionStatsRequest) (*SubscriptionStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionStats not implemented")
}
func (UnimplementedUserServiceServer) UserGetProblemStats(context.Context, *UProblemStatsRequest) (*UProblemStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetProblemStats not implemented")
}
func (UnimplementedUserServiceServer) UserGetLeaderboardStats(context.Context, *ULeaderboardRequest) (*ULeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetLeaderboardStats not implemented")
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

func _UserService_UserSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UserSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserSignup(ctx, req.(*Signup))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/VerifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyUser(ctx, req.(*OTP))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*Login))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ViewProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ViewProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/ViewProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ViewProfile(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/EditProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditProfile(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Password)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*Password))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BlockUser(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UnBlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UnBlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UnBlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UnBlockUser(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*UserNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserGetAllProblems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserGetAllProblems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UserGetAllProblems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserGetAllProblems(ctx, req.(*UserNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserGetProblemWithTestCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProblemId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserGetProblemWithTestCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UserGetProblemWithTestCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserGetProblemWithTestCases(ctx, req.(*UserProblemId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SubmitCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SubmitCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/SubmitCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SubmitCode(ctx, req.(*UserSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetUserStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserStats(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddSubscriptionPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(USubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddSubscriptionPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/AddSubscriptionPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddSubscriptionPlan(ctx, req.(*USubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetAllPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllPlans(ctx, req.(*UserNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(USubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UpdatePlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePlan(ctx, req.(*USubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GenerateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GenerateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GenerateInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GenerateInvoice(ctx, req.(*InvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInvoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInvoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetUserInvoices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInvoices(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_MakePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).MakePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/MakePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).MakePayment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PaymentSuccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PaymentSuccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/PaymentSuccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PaymentSuccess(ctx, req.(*ConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserProfileStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserProfileStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetUserProfileStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserProfileStats(ctx, req.(*UserStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetSubscriptionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSubscriptionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetSubscriptionStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSubscriptionStats(ctx, req.(*SubscriptionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserGetProblemStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UProblemStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserGetProblemStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UserGetProblemStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserGetProblemStats(ctx, req.(*UProblemStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserGetLeaderboardStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ULeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserGetLeaderboardStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UserGetLeaderboardStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserGetLeaderboardStats(ctx, req.(*ULeaderboardRequest))
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
			MethodName: "UserSignup",
			Handler:    _UserService_UserSignup_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _UserService_VerifyUser_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "ViewProfile",
			Handler:    _UserService_ViewProfile_Handler,
		},
		{
			MethodName: "EditProfile",
			Handler:    _UserService_EditProfile_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _UserService_BlockUser_Handler,
		},
		{
			MethodName: "UnBlockUser",
			Handler:    _UserService_UnBlockUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "UserGetAllProblems",
			Handler:    _UserService_UserGetAllProblems_Handler,
		},
		{
			MethodName: "UserGetProblemWithTestCases",
			Handler:    _UserService_UserGetProblemWithTestCases_Handler,
		},
		{
			MethodName: "SubmitCode",
			Handler:    _UserService_SubmitCode_Handler,
		},
		{
			MethodName: "GetUserStats",
			Handler:    _UserService_GetUserStats_Handler,
		},
		{
			MethodName: "AddSubscriptionPlan",
			Handler:    _UserService_AddSubscriptionPlan_Handler,
		},
		{
			MethodName: "GetAllPlans",
			Handler:    _UserService_GetAllPlans_Handler,
		},
		{
			MethodName: "UpdatePlan",
			Handler:    _UserService_UpdatePlan_Handler,
		},
		{
			MethodName: "GenerateInvoice",
			Handler:    _UserService_GenerateInvoice_Handler,
		},
		{
			MethodName: "GetUserInvoices",
			Handler:    _UserService_GetUserInvoices_Handler,
		},
		{
			MethodName: "MakePayment",
			Handler:    _UserService_MakePayment_Handler,
		},
		{
			MethodName: "PaymentSuccess",
			Handler:    _UserService_PaymentSuccess_Handler,
		},
		{
			MethodName: "GetUserProfileStats",
			Handler:    _UserService_GetUserProfileStats_Handler,
		},
		{
			MethodName: "GetSubscriptionStats",
			Handler:    _UserService_GetSubscriptionStats_Handler,
		},
		{
			MethodName: "UserGetProblemStats",
			Handler:    _UserService_UserGetProblemStats_Handler,
		},
		{
			MethodName: "UserGetLeaderboardStats",
			Handler:    _UserService_UserGetLeaderboardStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
