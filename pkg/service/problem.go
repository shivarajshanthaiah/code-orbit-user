package service

import (
	"context"

	problempb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/problem/problempb"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

// GetAllProblemsService implements interfaces.UserServiceInter.
func (u *UserService) GetAllProblemsService(p *pb.UserNoParam) (*pb.UserProblemList, error) {
	ctx := context.Background()

	result, err := u.ProblemClient.GetAllProblems(ctx, &problempb.ProbNoParam{})
	if err != nil {
		return nil, err
	}

	var problemList pb.UserProblemList
	for _, problem := range result.Problems {
		adProblem := &pb.UserProblem{
			ID:          problem.ID,
			Title:       problem.Title,
			Discription: problem.Discription,
			Difficulty:  problem.Difficulty,
			Tags:        problem.Tags,
			Is_Premium:  problem.Is_Premium,
		}
		problemList.Problems = append(problemList.Problems, adProblem)
	}

	return &problemList, nil
}
