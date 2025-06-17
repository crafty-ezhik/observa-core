package repository

import (
	"github.com/crafty-ezhik/observa-core/internal/domain/services"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RegisteredServicesRepository interface {
	GetService(id uint) (*services.RegisteredServices, error)
	GetAllServices() (*[]services.RegisteredServices, error)
	CreateService(service *services.RegisteredServices) error
	UpdateService()
	DeleteService()
}

type regServicesRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewRegisteredServices(db *gorm.DB, logger *zap.Logger) RegisteredServicesRepository {
	baseLogger := logger.With(
		zap.String("type", "RegisteredServicesRepository"),
	)
	return &regServicesRepository{
		db:     db,
		logger: baseLogger,
	}
}

func (r *regServicesRepository) GetService(id uint) (*services.RegisteredServices, error) {
	var result services.RegisteredServices
	err := r.db.First(&result, "id = ?", id)
	return &result, err.Error
}

func (r *regServicesRepository) GetAllServices() (*[]services.RegisteredServices, error) {
	var result []services.RegisteredServices
	err := r.db.Find(&result)
	return &result, err.Error
}

func (r *regServicesRepository) CreateService(service *services.RegisteredServices) error {
	return r.db.Create(service).Error

}
func (r *regServicesRepository) UpdateService() {}
func (r *regServicesRepository) DeleteService() {}
