package service

import (
	"context"

	problempb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/problem/problempb"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserService) GetAllUsersStatsService(p *pb.UserStatsRequest) (*pb.UserStatsProfileResponse, error) {
	var totalUsers, activeUsers, blockedUsers, primeUsers int64

	if err := u.Repo.FindAllUsersCount().Count(&totalUsers).Error; err != nil {
		return nil, err
	}

	if err := u.Repo.FindUsersByCondition("is_blocked = false").Count(&activeUsers).Error; err != nil {
		return nil, err
	}

	if err := u.Repo.FindUsersByCondition("is_blocked = true").Count(&blockedUsers).Error; err != nil {
		return nil, err
	}

	if err := u.Repo.FindUsersByCondition("is_prime_member = true").Count(&primeUsers).Error; err != nil {
		return nil, err
	}

	nonPrimeUsers := totalUsers - primeUsers

	return &pb.UserStatsProfileResponse{
		TotalUsers:    int32(totalUsers),
		ActiveUsers:   int32(activeUsers),
		BlockedUsers:  int32(blockedUsers),
		PrimeUsers:    int32(primeUsers),
		NonPrimeUsers: int32(nonPrimeUsers),
	}, nil
}

func (u *UserService) GetSubscriptionStatsService(p *pb.SubscriptionStatsRequest) (*pb.SubscriptionStatsResponse, error) {
	var basicPlanSubscribers, standardPlanSubscribers, premiumPlanSubscribers int32
	var totalAmountLifetime, totalAmountWeekly, totalAmountMonthly, totalAmountYearly uint32

	// Count subscribers by plan
	basicPlanSubscribers = u.Repo.CountByCondition("plan = 'Basic'")
	standardPlanSubscribers = u.Repo.CountByCondition("plan = 'Standard'")
	premiumPlanSubscribers = u.Repo.CountByCondition("plan = 'Premium'")

	// Calculate total amounts
	totalAmountLifetime = u.Repo.SumAmountsByCondition("") // All time
	totalAmountWeekly = u.Repo.SumAmountsByCondition("created_at >= NOW() - INTERVAL '7 DAYS'")
	totalAmountMonthly = u.Repo.SumAmountsByCondition("created_at >= NOW() - INTERVAL '1 MONTH'")
	totalAmountYearly = u.Repo.SumAmountsByCondition("created_at >= NOW() - INTERVAL '1 YEAR'")

	return &pb.SubscriptionStatsResponse{
		BasicPlanSubscribers:         int32(basicPlanSubscribers),
		StandardPlanSubscribers:      int32(standardPlanSubscribers),
		PremiumPlanSubscribers:       int32(premiumPlanSubscribers),
		TotalAmountCollectedLifetime: totalAmountLifetime,
		TotalAmountCollectedWeekly:   totalAmountWeekly,
		TotalAmountCollectedMonthly:  totalAmountMonthly,
		TotalAmountCollectedYearly:   totalAmountYearly,
	}, nil
}

func (a *UserService) UserGetProblemStatsService(req *pb.UProblemStatsRequest) (*pb.UProblemStatsResponse, error) {
	ctx := context.Background()

	result, err := a.ProblemClient.GetProblemStats(ctx, &problempb.ProblemStatsRequest{})
	if err != nil {
		return nil, err
	}

	return &pb.UProblemStatsResponse{
		TotalProblems:      result.TotalProblems,
		EasyProblems:       result.EasyProblems,
		MediumProblems:     result.MediumProblems,
		HardProblems:       result.HardProblems,
		TypeProblemCount:   result.TypeProblemCount,
		PremiumProblems:    result.PremiumProblems,
		NonPremiumProblems: result.NonPremiumProblems,
	}, nil
}

func (a *UserService) UserGetLeaderboardStatsService(req *pb.ULeaderboardRequest) (*pb.ULeaderboardResponse, error) {
	ctx := context.Background()
    response, err := a.ProblemClient.GetLeaderboard(ctx, &problempb.LeaderboardRequest{})
    if err != nil {
        return nil, err
    }

    var leaderboardEntries []*pb.ULeaderboardEntry
    for _, entry := range response.Leaderboard {
        leaderboardEntries = append(leaderboardEntries, &pb.ULeaderboardEntry{
            UserId:      entry.UserId,
            SolvedCount: entry.SolvedCount,
			Rank: entry.Rank,
        })
    }

    return &pb.ULeaderboardResponse{
        Leaderboard: leaderboardEntries,
    }, nil
}
