package interfaces

import pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"

type UserServiceInter interface {
	SignupService(p *pb.Signup) (*pb.Response, error)
	LoginService(p *pb.Login) (*pb.Response, error)
	VerificationService(p *pb.OTP) (*pb.Response, error)

	ViewProfileService(p *pb.ID) (*pb.Profile, error)
	EditProfileSerivice(p *pb.Profile) (*pb.Profile, error)
	ChangePasswordService(p *pb.Password) (*pb.Response, error)
	BlockUserService(p *pb.ID) (*pb.Response, error)
	UnBlockUserService(p *pb.ID) (*pb.Response, error)
	FindAllUsersService(p *pb.NoParam)(*pb.UserList, error)

}
