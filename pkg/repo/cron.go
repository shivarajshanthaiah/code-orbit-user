package repo

import (
	"errors"
	"time"

	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"
)

func (u *UserRepository) FindAllPrimeUsers() ([]model.User, error) {
	var users []model.User
	err := u.DB.Where("is_prime_member = ? AND membership_expiry_date > ?", true, time.Now()).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) UpdateUserMembership(user model.User) error {
	if user.UserID == "" {
		return errors.New("invalid user ID")
	}

	err := u.DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
