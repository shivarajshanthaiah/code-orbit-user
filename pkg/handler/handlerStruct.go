package handler

import (
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
	inter "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/service/interfaces"
)

type UserHandler struct {
	SVC inter.UserServiceInter
	pb.UserServiceServer
}

func NewUserHandler(svc inter.UserServiceInter) *UserHandler {
	return &UserHandler{
		SVC: svc,
	}
}
