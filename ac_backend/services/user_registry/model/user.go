package model

import (
	"gorm.io/gorm"
)

const (
	USER_SCOPE_PATIENT     = 1000
	USER_SCOPE_DOCTOR      = 2000
	USER_SCOPE_CLERK       = 3000
	USER_SCOPE_ADMIN       = 4000
	USER_SCOPE_SUPER_ADMIN = 5000
)

type User struct {
	gorm.Model
	Id                           string `gorm:"unique;not null"`
	UserScope                    int    `gorm:"not null"`
	Firstname                    string `gorm:"not null"`
	Lastname                     string `gorm:"not null"`
	Email                        string `gorm:"unique;not null"`
	SaltedPassword               string `gorm:"not null"`
	ResetPassword                bool   `gorm:"not null"`
	RefreshUrl                   string `gorm:""`
	ReturnUrl                    string `gorm:""`
	StripeOnboardingUrl          string `gorm:""`
	StripeOnboardingUrlExpiresAt int64  `gorm:""`
	StripeAccountId              string `gorm:""`
	StripeAccountCreatedAt       int64  `gorm:""`
	StripeCustomerId             string `gorm:""`
}
