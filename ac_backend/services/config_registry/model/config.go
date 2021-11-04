package model

import "gorm.io/gorm"

type Plugin struct {
	gorm.Model
	Id       string `gorm:"not null"`
	ConfigID string
	Name     string `gorm:"not null"`
	Version  string `gorm:"not null"`
}

type Config struct {
	gorm.Model
	Id               string   `gorm:"primaryKey"`
	UserId           uint     `gorm:"not null"`
	InstalledPlugins []Plugin `gorm:"not null"`
}
