package logging

import (
	"github.com/crafty-ezhik/observa-core/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func NewZapLogger(cfg *config.LogConfig) (*zap.Logger, error) {
	// Получаем уровень логирования из переменной окружения
	levelStr := strings.ToLower(cfg.Mode)
	if levelStr == "" {
		levelStr = "info" // дефолтный уровень
	}

	// Определяем уровень логирования
	var level zapcore.Level
	switch levelStr {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	// Настройка кодировщика JSON
	configJSON := zap.NewProductionEncoderConfig()
	configJSON.EncodeTime = zapcore.ISO8601TimeEncoder // или RFC3339Nano
	encoder := zapcore.NewJSONEncoder(configJSON)

	// Создаём кор
	core := zapcore.NewCore(encoder,
		zapcore.AddSync(os.Stdout), // вывод в stdout
		level,
	)

	// Добавляем глобальные поля для всех логов
	logger := zap.New(core)

	return logger, nil
}
