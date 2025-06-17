package utils

import (
	"github.com/crafty-ezhik/observa-core/internal/validate"
	"github.com/gofiber/fiber/v2"
)

func HandleBody[T any](c *fiber.Ctx, validator *validate.XValidator) (*T, error) {
	var body T
	if err := c.BodyParser(&body); err != nil {
		return nil, err
	}

	err := validator.Validate(body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
