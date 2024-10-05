package service

import (
	"context"

	adminpb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/admin/adminpb"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserService) GetAllPlansService(p *pb.UserNoParam) (*pb.UPlanList, error) {
	ctx := context.Background()

	result, err := u.AdminClient.GetAllPlans(ctx, &adminpb.AdNoParam{})
	if err != nil {
		return nil, err
	}

	var planList pb.UPlanList
	for _, plan := range result.Plans {
		uPlan := &pb.USubscription{
			ID:         plan.ID,
			Plan:       plan.Plan,
			Price:      plan.Price,
			Gst:        plan.Gst,
			TotalPrice: plan.TotalPrice,
		}
		planList.Plans = append(planList.Plans, uPlan)
	}
	return &planList, nil
}
