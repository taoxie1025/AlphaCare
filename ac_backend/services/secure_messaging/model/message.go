package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Id                   uint   `gorm:"primaryKey"`
	ThreadId             string `gorm:"not null"`
	Sender               string `gorm:"not null"`
	Receiver             string `gorm:"not null"`
	Timestamp            uint   `gorm:"not null"`
	Subject              string `gorm:"not null"`
	MessageBody          string `gorm:"not null"`
	Status               string `gorm:"not null"`
	Urgency              string
	IsViewed             bool          `gorm:"not null"`
	Preview              string        `gorm:"not null"`
	NotificationMethod   pq.Int64Array `gorm:"type:int[]"`
	Expiration           uint
	AttachmentId         string `gorm:"primarykey"`
	AttachmentType       string `gorm:"not null"`
	AttachmentTitle      string `gorm:"not null"`
	AttachmentUrl        string
	AttachmentExpiration uint
}
