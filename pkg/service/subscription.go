package service

import (
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserService) InsertPlanService(p *pb.USubscription) (*pb.Response, error) {

	plan := model.Subscription{
		Plan:       p.Plan,
		Duration:   p.Duration,
		Price:      p.Price,
		GST:        p.Gst,
		TotalPrice: p.TotalPrice,
	}

	err := u.Repo.CreateSubscription(&plan)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error creating subscription plan",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Subscription plan created sucessfully",
	}, nil
}

func (u *UserService) UpdatePlanService(p *pb.USubscription) (*pb.USubscription, error) {
	plan, err := u.Repo.GetPlanByID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	plan.Plan = p.Plan
	plan.Duration = p.Duration
	plan.Price = p.Price
	plan.GST = p.Gst
	plan.TotalPrice = p.TotalPrice

	err = u.Repo.UpdatePlan(plan)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (u *UserService) GetAllPlansService(p *pb.UserNoParam) (*pb.UPlanList, error) {

	result, err := u.Repo.GetAllPlans()
	if err != nil {
		return nil, err
	}

	var planList pb.UPlanList
	for _, plan := range *result {
		pbPlan := &pb.USubscription{
			ID:         uint32(plan.ID),
			Plan:       plan.Plan,
			Duration:   plan.Duration,
			Price:      plan.Price,
			Gst:        plan.GST,
			TotalPrice: plan.TotalPrice,
		}
		planList.Plans = append(planList.Plans, pbPlan)
	}
	return &planList, nil
}

