package repo

import (
	inter "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/repo/interfaces"
	"gorm.io/gorm"
)

// Connecting user repository to DB
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) inter.UserRepoInter {
	return &UserRepository{
		DB: db,
	}
}
