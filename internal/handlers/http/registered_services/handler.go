package registered_services

import (
	"github.com/crafty-ezhik/observa-core/internal/repository"
	"github.com/crafty-ezhik/observa-core/internal/utils"
	"github.com/crafty-ezhik/observa-core/internal/validate"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RegisteredServicesHandler interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type registeredServicesHandler struct {
	RegServiceRepo repository.RegisteredServicesRepository
	logger         *zap.Logger
	v              *validate.XValidator
}

func NewRegisteredServicesHandler(regServiceRepo repository.RegisteredServicesRepository, logger *zap.Logger, v *validate.XValidator) RegisteredServicesHandler {
	baseLogger := logger.With(zap.String("type", "RegisteredServicesHandler"))
	return &registeredServicesHandler{
		RegServiceRepo: regServiceRepo,
		logger:         baseLogger,
		v:              v,
	}
}

func (h *registeredServicesHandler) Get(c *fiber.Ctx) error {
	return nil
}
func (h *registeredServicesHandler) Create(c *fiber.Ctx) error {
	body, err := utils.HandleBody[CreateRequest](c, h.v)
	if err != nil {
		h.logger.Error("Failed to handle request", zap.Error(err))
	}

	return c.Status(fiber.StatusCreated).JSON(body)
}
