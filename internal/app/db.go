package app

import (
	"fmt"
	"github.com/crafty-ezhik/observa-core/internal/config"
	"github.com/crafty-ezhik/observa-core/internal/domain/event"
	"github.com/crafty-ezhik/observa-core/internal/domain/services"
	"github.com/crafty-ezhik/observa-core/internal/domain/subscription"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database *gorm.DB

func InitDatabase(config *config.DBConfig) Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
		config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func GoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&services.RegisteredServices{},
		&event.Event{},
		&subscription.Subscription{},
	)
	if err != nil {
		panic(err)
	}
}
