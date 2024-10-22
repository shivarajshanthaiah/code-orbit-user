package repo

import "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"

func (u *UserRepository) CreateSubscription(plan *model.Subscription) error {
	if err := u.DB.Create(&plan).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetAllPlans() (*[]model.Subscription, error) {
	var plans []model.Subscription
	if err := u.DB.Find(&plans).Error; err != nil {
		return nil, err
	}
	return &plans, nil
}

func (u *UserRepository) GetPlanByID(planID uint) (*model.Subscription, error) {
	var plan model.Subscription
	err := u.DB.Where("id = ?", planID).First(&plan).Error
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

func (u *UserRepository) UpdatePlan(plan *model.Subscription) error {
	if err := u.DB.Save(&plan).Error; err != nil {
		return err
	}
	return nil
}
