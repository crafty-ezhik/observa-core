package registered_services

import (
	"github.com/crafty-ezhik/observa-core/internal/domain/services"
	"github.com/crafty-ezhik/observa-core/internal/repository"
	"github.com/crafty-ezhik/observa-core/internal/utils"
	"github.com/crafty-ezhik/observa-core/internal/validate"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

type RegisteredServicesHandler interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type registeredServicesHandler struct {
	RegServiceRepo repository.RegisteredServicesRepository
	log            *zap.Logger
	v              *validate.XValidator
}

func NewRegisteredServicesHandler(regServiceRepo repository.RegisteredServicesRepository, logger *zap.Logger, v *validate.XValidator) RegisteredServicesHandler {
	baseLogger := logger.With(zap.String("type", "RegisteredServicesHandler"))
	return &registeredServicesHandler{
		RegServiceRepo: regServiceRepo,
		log:            baseLogger,
		v:              v,
	}
}

func (h *registeredServicesHandler) Get(c *fiber.Ctx) error {
	return nil
}
func (h *registeredServicesHandler) Create(c *fiber.Ctx) error {
	h.log.Info("Запрос на создание сервиса")
	body, err := utils.HandleBody[CreateRequest](c, h.v)
	if err != nil {
		h.log.Error("Failed to handle request", zap.Error(err))
	}
	service := services.RegisteredServices{
		Name:           body.Name,
		Description:    body.Description,
		HealthCheckUrl: body.HealthUrl,
		OwnerEmail:     body.OwnerEmail,
		Tags:           body.Tags,
		Status:         services.Healthy,
		LastCheckedAt:  time.Now(),
	}
	err = h.RegServiceRepo.CreateService(&service)
	if err != nil {
		h.log.Error("Failed to register service", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(body)
}
