package app

import (
	"fmt"
	"github.com/crafty-ezhik/observa-core/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config *config.DBConfig) *gorm.DB {
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
