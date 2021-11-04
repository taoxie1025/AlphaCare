package model

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Id                      string  `gorm:"primaryKey"`
	FirstName               string  `gorm:"not null"`
	LastName                string  `gorm:"not null"`
	SSN                     string  `gorm:"not null"`
	Gender                  string  `gorm:"not null"`
	YearOfBirth             int32   `gorm:"not null"`
	MonthOfBirth            int32   `gorm:"not null"`
	DayOfBirth              int32   `gorm:"not null"`
	Status                  string  `gorm:"not null"`
	PrimaryCarePhysician    string  `gorm:"not null"`
	PhoneHome               string  `gorm:"not null"`
	PhoneWork               string  `gorm:"not null"`
	PhoneMobile             string  `gorm:"not null"`
	AddressStreet           string  `gorm:"not null"`
	AddressCity             string  `gorm:"not null"`
	AddressState            string  `gorm:"not null"`
	AddressZip              string  `gorm:"not null"`
	WeightInKg              float64 `gorm:"not null"`
	HeightInCm              float64 `gorm:"not null"`
	Ethnicity               string  `gorm:"not null"`
	Email                   string  `gorm:"unique;not null"`
	PatientEmergencyContact string  `gorm:"not null"`
	Referral                string  `gorm:"not null"`
	CreatedAt               int64   `gorm:"not null"`
	UpdatedAt               int64   `gorm:"not null"`
}
