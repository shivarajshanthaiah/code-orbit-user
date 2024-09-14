package interfaces

import "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"

type UserRepoInter interface {
	CreateUser(user *model.User) (string, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(userID string) (*model.User, error)
	UpdateUser(user *model.User) error
	FindAllUsers() (*[]model.User, error)
}