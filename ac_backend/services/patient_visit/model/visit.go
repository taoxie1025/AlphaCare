package model

import (
	"gorm.io/gorm"
)

type PatientVisit struct {
	gorm.Model
	Id              string          `gorm:"primaryKey"`
	PatientId       string          `gorm:"not null"`
	ServiceProvider ServiceProvider `json:"service_provider" gorm:"foreignKey:Id"`
	ReasonForVisit  string
	Allergies       Allergies       `json:"allergies" gorm:"embedded"`
	Assessment      Assessment      `json:"assessment" gorm:"embedded"`
	ScheduledTests  ScheduledTests  `json:"scheduled_test" gorm:"embedded"`
	VitalSigns      VitalSigns      `json:"vital_signs" gorm:"embedded"`
	HealthProblems  []HealthProblem `gorm:"foreignKey:Id"`
}

type ServiceProvider struct {
	gorm.Model
	Id        string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Type      string `gorm:"not null"`
}

type Allergies struct {
	gorm.Model
	Allergies        string
	AdverseReactions string
	Alerts           string
}

type Assessment struct {
	gorm.Model
	Assessment string
}

type ScheduledTests struct {
	gorm.Model
	TestName      string
	ScheduledDate string
}

type VitalSigns struct {
	gorm.Model
	SignName string
}

type HealthProblem struct {
	gorm.Model
	Id        string
	Condition string
	Status    string
	Comment   string
}
