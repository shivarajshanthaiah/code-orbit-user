package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserHandler) UserSignup(ctx context.Context, p *pb.Signup) (*pb.Response, error) {
	response, err := u.SVC.SignupService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) VerifyUser(ctx context.Context, p *pb.OTP) (*pb.Response, error) {
	response, err := u.SVC.VerificationService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) UserLogin(ctx context.Context, p *pb.Login) (*pb.Response, error) {
	response, err := u.SVC.LoginService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
