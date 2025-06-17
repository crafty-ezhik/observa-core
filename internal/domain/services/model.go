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
	ID             uint           `gorm:"primarykey" json:"id"`
	Name           string         `gorm:"varchar(128)" json:"name"`
	Description    string         `gorm:"text" json:"description"`
	HealthCheckUrl string         `gorm:"varchar(128)" json:"health_check_url"`
	OwnerEmail     string         `gorm:"varchar(128)" json:"owner_email"`
	Tags           []string       `gorm:"column:tags;type:jsonb;serializer:json" json:"tags"`
	Status         string         `gorm:"varchar(10)" json:"status"`
	LastCheckedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"last_checked_at"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"uniqueIndex" json:"-"`
}
