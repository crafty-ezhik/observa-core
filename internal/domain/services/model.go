package services

import (
	"gorm.io/gorm"
	"time"
)

type RegisteredServices struct {
	ID             uint `gorm:"primarykey"`
	Name           string
	Description    string
	HealthCheckUrl string
	OwnerEmail     string
	Tags           []string
	Status         string
	LastCheckedAt  time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
