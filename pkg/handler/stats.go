package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserHandler) GetUserProfileStats(ctx context.Context, p *pb.UserStatsRequest) (*pb.UserStatsProfileResponse, error) {
	response, err := u.SVC.GetAllUsersStatsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) GetSubscriptionStats(ctx context.Context, p *pb.SubscriptionStatsRequest) (*pb.SubscriptionStatsResponse, error) {
	response, err := u.SVC.GetSubscriptionStatsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *UserHandler) UserGetProblemStats(ctx context.Context, p *pb.UProblemStatsRequest) (*pb.UProblemStatsResponse, error) {
	response, err := a.SVC.UserGetProblemStatsService(p)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (a *UserHandler) UserGetLeaderboardStats(ctx context.Context, p *pb.ULeaderboardRequest) (*pb.ULeaderboardResponse, error) {
	response, err := a.SVC.UserGetLeaderboardStatsService(p)
	if err != nil {
		return nil, err
	}
	return response, nil
}