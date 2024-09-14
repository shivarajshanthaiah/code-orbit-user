package service

import (
	"github.com/shivaraj-shanthaiah/code_orbit_user/config"
	__ "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
	inter "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/repo/interfaces"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/service/interfaces"
)

type UserService struct {
	Repo   inter.UserRepoInter
	twilio *config.TwilioService
	redis  *config.RedisService
}

// FindAllUsersService implements interfaces.UserServiceInter.
func (u *UserService) FindAllUsersService(p *__.NoParam) (*__.UserList, error) {
	panic("unimplemented")
}

func NewUserService(repo inter.UserRepoInter, redis *config.RedisService, twilio *config.TwilioService) interfaces.UserServiceInter {
	return &UserService{
		Repo:   repo,
		twilio: twilio,
		redis:  redis,
	}
}
