package model

import "time"

type User struct {
	UserID               string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserName             string    `gorm:"not null"`
	Email                string    `gorm:"not null;unique"`
	Password             string    `gorm:"not null"`
	Phone                string    `gorm:"not null;unique"`
	IsPrimeMember        bool      `gorm:"default:false"`
	IsBlocked            bool      `gorm:"default:false"`
	MembershipExpiryDate time.Time `gorm:"default:null"`
	CreatedAT            time.Time
	UpdatedAT            time.Time
}