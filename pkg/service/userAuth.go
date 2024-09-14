package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/shivaraj-shanthaiah/code_orbit_user/config"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
	"github.com/shivaraj-shanthaiah/code_orbit_user/utils"
	"gorm.io/gorm"
)

// SignupService implements interfaces.UserServiceInter.
func (u *UserService) SignupService(p *pb.Signup) (*pb.Response, error) {
	hashedPassword, err := utils.HashPassword(p.Password)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in hashing Password",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, errors.New("unable to hash password")
	}

	user := &model.User{
		UserName: p.User_Name,
		Email:    p.Email,
		Password: hashedPassword,
		Phone:    p.Phone,
	}

	_, err = u.Repo.FindUserByEmail(user.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "email already exists",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	resp, err := u.twilio.SendTwilioOTP(p.Phone)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in sending otp using twilio",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	if resp.Status != nil {
		fmt.Println(*resp.Status)
	} else {
		fmt.Println(resp.Status)
	}

	userData, err := json.Marshal(&user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in marshalling data",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, errors.New("error while marshalling data")
	}

	key := fmt.Sprintf("user_%v", p.Email)
	err = u.redis.SetDataInRedis(key, userData, time.Minute*3)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in setting data in redis",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, errors.New("error while setting data in redis")
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "OTP generayed succesfully. Go to veriication page...",
	}, nil
}

// VerificationService implements interfaces.UserServiceInter.
func (u *UserService) VerificationService(p *pb.OTP) (*pb.Response, error) {
	fmt.Println("Im here")
	key := fmt.Sprintf("user_%v", p.Email)

	userData, err := u.redis.GetFromRedis(key)
	fmt.Println(userData)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting data from redis",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}
	fmt.Println("Im here too")
	var user model.User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in unmarshalling data",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	err = u.twilio.VerifyTwilioOTP(user.Phone, p.OTP)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in verifying twilio OTP",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	userId, err := u.Repo.CreateUser(&user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error creating user in database",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, errors.New("unable to create user")
	}

	log.Printf("User created with ID: %v", userId)

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "OTP verifed successfully. Login to continue",
	}, nil
}

// LoginService implements interfaces.UserServiceInter.
func (u *UserService) LoginService(p *pb.Login) (*pb.Response, error) {
	user, err := u.Repo.FindUserByEmail(p.Email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(p.Password, user.Password) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "password mismatch",
		}, errors.New("incorrect password")
	}

	if user.IsBlocked {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "user is blocked by admin",
		}, errors.New("you are blocked by the admin")
	}

	jwtToken, err := utils.GenerateToken(config.LoadConfig().SECRETKEY, user.Email, user.UserID)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "erros in generating token",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, errors.New("error generating token")
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Login successfull",
		Payload: &pb.Response_Data{Data: jwtToken},
	}, nil
}
