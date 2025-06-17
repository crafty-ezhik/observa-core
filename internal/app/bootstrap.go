package app

import (
	"github.com/crafty-ezhik/observa-core/internal/config"
	"github.com/crafty-ezhik/observa-core/internal/handlers/http/registered_services"
	"github.com/crafty-ezhik/observa-core/internal/logging"
	"github.com/crafty-ezhik/observa-core/internal/repository"
	"github.com/crafty-ezhik/observa-core/internal/validate"
	"github.com/gofiber/fiber/v2"
)

type AppDependencies struct {
}

// Bootstrap - создает все необходимые зависимости приложения и возвращает структуру с ними для дальнейшей работы
func Bootstrap() (*config.Config, *fiber.App) {
	cfg, err := config.LoadConfig("configs")
	if err != nil {
		panic(err)
	}

	// Инициализация логера
	logger, err := logging.NewZapLogger(&cfg.Log)
	if err != nil {
		panic(err)
	}

	// Инициализация бд и выполнение миграций
	db := InitDatabase(&cfg.Db)
	GoMigrate(db)

	// Инициализация валидатора
	customValidator := validate.NewXValidator()

	// Инициализировать все сервисы, репозитории, чекеры, хендлеры
	regServiceRepo := repository.NewRegisteredServices(db, logger)
	regServiceHandler := registered_services.NewRegisteredServicesHandler(regServiceRepo, logger, customValidator)

	app := fiber.New()
	app.Route("service", func(router fiber.Router) {
		router.Post("/", regServiceHandler.Create)
		router.Get("/", regServiceHandler.GetAll)
	})

	return cfg, app
}
