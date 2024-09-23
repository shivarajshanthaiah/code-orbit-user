package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserHandler) UserGetAllProblems(ctx context.Context, p *pb.UserNoParam) (*pb.UserProblemList, error) {
	response, err := u.SVC.GetAllProblemsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) UserGetProblemWithTestCases(ctx context.Context, p *pb.UserProblemId) (*pb.UserTestcaseResponse, error) {
	response, err := u.SVC.GetProblemWithTestCasesService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}
