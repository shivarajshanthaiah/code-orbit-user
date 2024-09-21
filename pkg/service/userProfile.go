package service

import (
	"errors"
	"fmt"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
	"github.com/shivaraj-shanthaiah/code_orbit_user/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserService) ViewProfileService(p *pb.ID) (*pb.Profile, error) {
	user, err := u.Repo.FindUserByID(p.ID)
	if err != nil {
		return nil, err
	}

	userModel := &pb.Profile{
		User_ID:                user.UserID,
		User_Name:              user.UserName,
		Email:                  user.Email,
		Phone:                  user.Phone,
		Is_Prime_Member:        user.IsPrimeMember,
		Membership_Expiry_Date: timestamppb.New(user.MembershipExpiryDate),
	}
	return userModel, nil
}

func (u *UserService) EditProfileSerivice(p *pb.Profile) (*pb.Profile, error) {
	user, err := u.Repo.FindUserByID(p.User_ID)
	if err != nil {
		return nil, err
	}
	user.UserName = p.User_Name
	user.Phone = p.Phone

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (u *UserService) ChangePasswordService(p *pb.Password) (*pb.Response, error) {
	user, err := u.Repo.FindUserByID(p.User_ID)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting user from database",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	fmt.Println(p.Old_Password, user.Password)
	if !utils.CheckPassword(p.Old_Password, user.Password) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "old password is incorrect",
		}, errors.New("old password mismatch")
	}

	if p.New_Password != p.Confirm_Password {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "NewPassword and confirm password mismatch",
		}, errors.New("newPassword mismatch")
	}

	newPassword, err := utils.HashPassword(p.New_Password)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error while hashing password",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	user.Password = newPassword

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error while updating user password into database",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Password changed successfully",
	}, nil
}

func (u *UserService) BlockUserService(p *pb.ID) (*pb.Response, error) {
	user, err := u.Repo.FindUserByID(p.ID)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting user from database",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	user.IsBlocked = true

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error updating user into database",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "user blocked successfully",
	}, nil
}

func (u *UserService) UnBlockUserService(p *pb.ID) (*pb.Response, error) {
	user, err := u.Repo.FindUserByID(p.ID)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting user from database",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	user.IsBlocked = false

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error updating user into database",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "user unblocked successfully",
	}, nil
}

// FindAllUsersService implements interfaces.UserServiceInter.
func (u *UserService) FindAllUsersService(p *pb.UserNoParam) (*pb.UserList, error) {
	result, err := u.Repo.FindAllUsers()
	if err != nil {
		return nil, err
	}

	var userList pb.UserList
	for _, user := range *result {
		userList.Users = append(userList.Users, &pb.Profile{
			User_ID:                user.UserID,
			User_Name:              user.UserName,
			Email:                  user.Email,
			Phone:                  user.Phone,
			Is_Prime_Member:        user.IsPrimeMember,
			Is_Blocked:             user.IsBlocked,
			Membership_Expiry_Date: timestamppb.New(user.MembershipExpiryDate),
		})
	}
	return &userList, nil
}
