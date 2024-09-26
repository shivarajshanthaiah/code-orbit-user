package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserHandler) SubmitCode(ctx context.Context, p *pb.UserSubmissionRequest) (*pb.UserSubmissionResponse, error) {
	response, err := u.SVC.SubmitCodeService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}
