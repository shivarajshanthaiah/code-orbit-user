package interfaces

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

type UserServiceInter interface {
	SignupService(p *pb.Signup) (*pb.Response, error)
	LoginService(p *pb.Login) (*pb.Response, error)
	VerificationService(p *pb.OTP) (*pb.Response, error)

	ViewProfileService(p *pb.ID) (*pb.Profile, error)
	EditProfileSerivice(p *pb.Profile) (*pb.Profile, error)
	ChangePasswordService(p *pb.Password) (*pb.Response, error)
	BlockUserService(p *pb.ID) (*pb.Response, error)
	UnBlockUserService(p *pb.ID) (*pb.Response, error)
	FindAllUsersService(p *pb.UserNoParam) (*pb.UserList, error)

	GetAllProblemsService(p *pb.UserNoParam) (*pb.UserProblemList, error)
	GetProblemWithTestCasesService(ctx context.Context, req *pb.UserProblemId) (*pb.UserTestcaseResponse, error)

	SubmitCodeService(ctx context.Context, req *pb.UserSubmissionRequest) (*pb.UserSubmissionResponse, error)
	GetUserStatsService(ctx context.Context, p *pb.ID) (*pb.UserStatsResponse, error)

	GetAllPlansService(p *pb.UserNoParam) (*pb.UPlanList, error)
	GenerateInvoiceService(ctx context.Context, req *pb.InvoiceRequest) (*pb.Response, error)
	MakePaymentService(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error)
	PaymentSuccessService(ctx context.Context, req *pb.ConfirmRequest) (*pb.ConfirmResponse, error)

	GetAllUsersStatsService(p *pb.UserStatsRequest) (*pb.UserStatsProfileResponse, error)
	GetSubscriptionStatsService(p *pb.SubscriptionStatsRequest) (*pb.SubscriptionStatsResponse, error)
}
