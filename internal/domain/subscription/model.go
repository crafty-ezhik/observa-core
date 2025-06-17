package subscription

import (
	"gorm.io/gorm"
	"time"
)

type Subscription struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
