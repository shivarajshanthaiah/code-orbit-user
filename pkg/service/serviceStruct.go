package service

import (
	"github.com/shivaraj-shanthaiah/code_orbit_user/config"
	problempb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/problem/problempb"
	inter "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/repo/interfaces"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/service/interfaces"
)

type UserService struct {
	Repo          inter.UserRepoInter
	twilio        *config.TwilioService
	redis         *config.RedisService
	ProblemClient problempb.ProblemServiceClient
}

func NewUserService(repo inter.UserRepoInter, redis *config.RedisService, twilio *config.TwilioService, problemClient problempb.ProblemServiceClient) interfaces.UserServiceInter {
	return &UserService{
		Repo:          repo,
		twilio:        twilio,
		redis:         redis,
		ProblemClient: problemClient,
	}
}
