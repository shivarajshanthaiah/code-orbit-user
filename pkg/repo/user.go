package repo

import "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"

func (u *UserRepository) CreateUser(user *model.User) (string, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return "", err
	}
	return user.UserID, nil
}

func (u *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User

	if err := u.DB.Model(&model.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindUserByID(userID string) (*model.User, error) {
	var user model.User

	if err := u.DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	if err := u.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) FindAllUsers() (*[]model.User, error) {
	var users []model.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}