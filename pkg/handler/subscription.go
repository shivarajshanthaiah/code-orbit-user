package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserHandler) UserGetAllPlans(ctx context.Context, p *pb.UserNoParam) (*pb.UPlanList, error) {
	response, err := u.SVC.GetAllPlansService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
