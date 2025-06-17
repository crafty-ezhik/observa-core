package services

import (
	"gorm.io/gorm"
	"time"
)

const (
	Healthy   = "Healthy"
	Unhealthy = "Unhealthy"
)

type RegisteredServices struct {
	ID             uint     `gorm:"primarykey"`
	Name           string   `gorm:"varchar(128)"`
	Description    string   `gorm:"text"`
	HealthCheckUrl string   `gorm:"varchar(128)"`
	OwnerEmail     string   `gorm:"varchar(128)"`
	Tags           []string `gorm:"column:tags;type:jsonb;serializer:json"`
	Status         string   `gorm:"varchar(10)"`
	LastCheckedAt  time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"uniqueIndex"`
}
