package service

import (
	"context"

	problempb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/problem/problempb"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserService) SubmitCodeService(ctx context.Context, req *pb.UserSubmissionRequest) (*pb.UserSubmissionResponse, error) {

	submission := &problempb.SubmissionRequest{
		UserId:    req.UserId,
		ProblemId: req.ProblemId,
		Language:  req.Language,
		Code:      req.Code,
	}

	result, err := u.ProblemClient.SubmitCode(ctx, submission)
	if err != nil {
		return &pb.UserSubmissionResponse{
			Status:  pb.UserSubmissionResponse_ERROR,
			Message: "Failed to submit code",
		}, err
	}

	return &pb.UserSubmissionResponse{
		Status:       pb.UserSubmissionResponse_Status(result.Status),
		Message:      result.Message,
		Passed:       result.Passed,
		Failed:       result.Failed,
		SubmissionId: result.SubmissionId,
	}, nil
}
