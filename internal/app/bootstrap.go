package app

import "github.com/crafty-ezhik/observa-core/internal/config"

type AppDependencies struct {
}

// Bootstrap - создает все необходимые зависимости приложения и возвращает структуру с ними для дальнейшей работы
func Bootstrap() *config.Config {
	cfg, err := config.LoadConfig("configs")
	if err != nil {
		panic(err)
	}

	db := InitDatabase(&cfg.Db)
	return cfg
}
