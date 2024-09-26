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
			ID:   problem.ID,
			Title:       problem.Title,
			Discription: problem.Discription,
			Difficulty:  problem.Difficulty,
			Type:        problem.Type,
			IsPremium:   problem.IsPremium,
		}
		problemList.Problems = append(problemList.Problems, adProblem)
	}

	return &problemList, nil
}

func (a *UserService) GetProblemWithTestCasesService(ctx context.Context, req *pb.UserProblemId) (*pb.UserTestcaseResponse, error) {
	// Call the Problem Service
	result, err := a.ProblemClient.GetProblemWithTestCases(ctx, &problempb.ProblemId{
		ID: req.ID,
	})
	if err != nil {
		return &pb.UserTestcaseResponse{
			Status:  pb.UserTestcaseResponse_ERROR,
			Message: "Failed to fetch problem and test cases",
			Payload: &pb.UserTestcaseResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	// Check if there was an error in the ProblemService response
	if result.Status == problempb.GetProblemResponse_ERROR {
		return &pb.UserTestcaseResponse{
			Status:  pb.UserTestcaseResponse_ERROR,
			Message: result.Message,
			Payload: &pb.UserTestcaseResponse_Error{
				Error: result.GetError(),
			},
		}, nil
	}

	// Map the ProblemService response to the AdminService response format
	problem := result.GetData().Problem
	testCases := result.GetData().TestCases

	var adminTestCases []*pb.UserTestCase
	for _, tc := range testCases {
		adminTestCases = append(adminTestCases, &pb.UserTestCase{
			TestCaseId:     tc.TestCaseId,
			Input:          tc.Input,
			ExpectedOutput: tc.ExpectedOutput,
		})
	}

	// Prepare the response for the AdminService
	return &pb.UserTestcaseResponse{
		Status:  pb.UserTestcaseResponse_OK,
		Message: "Problem and test cases fetched successfully",
		Payload: &pb.UserTestcaseResponse_Data{
			Data: &pb.UserProblemWithTestCases{
				Problem: &pb.UserProblem{
					ID:   problem.ID,
					Title:       problem.Title,
					Discription: problem.Discription,
					Difficulty:  problem.Difficulty,
					Type:        problem.Type,
					IsPremium:   problem.IsPremium,
				},
				TestCases: adminTestCases,
			},
		},
	}, nil
}
