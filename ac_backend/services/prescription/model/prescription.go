package model

import (
	"gorm.io/gorm"
)

type Prescription struct {
	gorm.Model
	Id                   string `gorm:"primaryKey"`
	PatientId            string `gorm:"not null"`
	PrimaryCarePhysician string `gorm:"not null"`
	HealthcareId         string `gorm:"not null"`
	Timestamp            uint   `gorm:"not null"`
	Subject              string `gorm:"not null"`
	MessageBody          string `gorm:"not null"`
	Urgency              string
	Status               string `gorm:"not null"`
	Quantity             uint
	StartDate            string `gorm:"not null"`
	VisitDate            string `gorm:"not null"`
	EndDate              string `gorm:"not null"`
	Qualifier            string
	AttachmentId         string `gorm:"primarykey"`
	AttachmentType       string `gorm:"not null"`
	AttachmentTitle      string `gorm:"not null"`
	AttachmentUrl        string
	Refills              []Refill  `gorm:"foreignKey:PrescriptionId"`
	Renewals             []Renewal `gorm:"foreignKey:PrescriptionId"`
	Dose                 string
	Frequency            string
}

type Refill struct {
	gorm.Model
	Id             string `gorm:"primaryKey"`
	Date           string `gorm:"not null"`
	count          uint
	PrescriptionId string
}

type Renewal struct {
	gorm.Model
	Id             string `gorm:"primaryKey"`
	Date           string `gorm:"not null"`
	count          uint
	PrescriptionId string
}
