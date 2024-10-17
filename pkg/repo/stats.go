package repo

import (
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"
	"gorm.io/gorm"
)

func (r *UserRepository) FindAllUsersCount() *gorm.DB {
	return r.DB.Model(&model.User{})
}

func (r *UserRepository) FindUsersByCondition(condition string) *gorm.DB {
	return r.DB.Model(&model.User{}).Where(condition)
}

func (r *UserRepository) CountByCondition(condition string) int32 {
	var count int64
	if condition != "" {
		r.DB.Model(&model.Invoice{}).Where(condition).Count(&count)
	} else {
		r.DB.Model(&model.Invoice{}).Count(&count)
	}
	return int32(count)
}

func (r *UserRepository) SumAmountsByCondition(condition string) uint32 {
	var sum uint32
	if condition != "" {
		r.DB.Model(&model.Invoice{}).Where(condition).Select("COALESCE(SUM(amount), 0)").Scan(&sum)
	} else {
		r.DB.Model(&model.Invoice{}).Select("COALESCE(SUM(amount), 0)").Scan(&sum)
	}
	return sum
}
