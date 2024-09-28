package service

import (
	"context"
	"errors"

	problempb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/problem/problempb"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
	"github.com/shivaraj-shanthaiah/code_orbit_user/utils"
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
			ID: problem.ID,
			// Discription: problem.Discription,
			Difficulty: problem.Difficulty,
			Type:       problem.Type,
			IsPremium:  problem.IsPremium,
		}
		problemList.Problems = append(problemList.Problems, adProblem)
	}

	return &problemList, nil
}

func (a *UserService) GetProblemWithTestCasesService(ctx context.Context, req *pb.UserProblemId) (*pb.UserTestcaseResponse, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return &pb.UserTestcaseResponse{
			Status:  pb.UserTestcaseResponse_ERROR,
			Message: "User not authenticated",
			Payload: &pb.UserTestcaseResponse_Error{
				Error: "User authentication failed",
			},
		}, err
	}

	user, err := a.Repo.FindUserByID(userID)
	if err != nil {
		return &pb.UserTestcaseResponse{
			Status:  pb.UserTestcaseResponse_ERROR,
			Message: "User not found",
			Payload: &pb.UserTestcaseResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	// Fetch the problem details from Problem Service
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

	// Check if problem is premium and if the user is not a premium member
	if result.GetData().Problem.IsPremium && !user.IsPrimeMember {
		return &pb.UserTestcaseResponse{
			Status:  pb.UserTestcaseResponse_ERROR,
			Message: "This is a premium problem. Please subscribe to access this problem.",
			Payload: &pb.UserTestcaseResponse_Error{
				Error: "User is not a premium member",
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

	// Return the problem and test cases
	return &pb.UserTestcaseResponse{
		Status:  pb.UserTestcaseResponse_OK,
		Message: "Problem and test cases fetched successfully",
		Payload: &pb.UserTestcaseResponse_Data{
			Data: &pb.UserProblemWithTestCases{
				Problem: &pb.UserProblem{
					ID:          problem.ID,
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

func (u *UserService) GetUserStatsService(ctx context.Context, p *pb.ID) (*pb.UserStatsResponse, error) {

	user, err := u.Repo.FindUserByID(p.ID)
	if err != nil {
		return nil, errors.New("error fetching user details: " + err.Error())
	}

	result, err := u.ProblemClient.GetUserStats(ctx, &problempb.UserID{ID: p.ID})
	if err != nil {
		return nil, err
	}

	return &pb.UserStatsResponse{
		UserId:                     p.ID,
		UserName:                   user.UserName,
		TotalAttempts:              result.TotalAttempts,
		TotalSuccessfulSubmissions: result.TotalSuccessfulSubmissions,
		TotalFailedSubmissions:     result.TotalFailedSubmissions,
		EasyAttempts:               result.EasyAttempts,
		MediumAttempts:             result.MediumAttempts,
		HardAttempts:               result.HardAttempts,
	}, nil
}
