package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserHandler) ViewProfile(ctx context.Context, p *pb.ID) (*pb.Profile, error) {
	response, err := u.SVC.ViewProfileService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) EditProfile(ctx context.Context, p *pb.Profile) (*pb.Profile, error) {
	response, err := u.SVC.EditProfileSerivice(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) ChangePassword(ctx context.Context, p *pb.Password) (*pb.Response, error) {
	response, err := u.SVC.ChangePasswordService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) BlockUser(ctx context.Context, p *pb.ID) (*pb.Response, error) {
	response, err := u.SVC.BlockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) UnBlockUser(ctx context.Context, p *pb.ID) (*pb.Response, error) {
	response, err := u.SVC.UnBlockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}