package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserHandler) GetAllPlans(ctx context.Context, p *pb.UserNoParam) (*pb.UPlanList, error) {
	response, err := u.SVC.GetAllPlansService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) AddSubscriptionPlan(ctx context.Context, p *pb.USubscription) (*pb.Response, error) {
	response, err := u.SVC.InsertPlanService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) UpdatePlan(ctx context.Context, p *pb.USubscription) (*pb.USubscription, error) {
	response, err := u.SVC.UpdatePlanService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}