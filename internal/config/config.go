package config

import (
	"errors"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
)

// Config - главная структура конфигурации приложения
type Config struct {
	Auth   AuthConfig   `mapstructure:"auth"`
	Db     DBConfig     `mapstructure:"database"`
	Server ServerConfig `mapstructure:"server"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Log    LogConfig    `mapstructure:"log"`
}

// AuthConfig - конфигурация для авторизации в приложении
type AuthConfig struct {
	SigningKey   string        `mapstructure:"signing_key"`
	ApiKeyHeader string        `mapstructure:"api_key_header"`
	AccessTTL    time.Duration `mapstructure:"access_ttl"`
	RefreshTTL   time.Duration `mapstructure:"refresh_ttl"`
}

// DBConfig - конфигурация для настройки подключения к БД
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"sslmode"`
}

// ServerConfig - конфигурация для настройки сервера приложения
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// RedisConfig - конфигурация для настройки redis
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

// LogConfig - базовая конфигурация для настройки логирования
type LogConfig struct {
	Mode       string   `mapstructure:"mode"`
	Encoding   string   `mapstructure:"encoding"`
	OutputPath []string `mapstructure:"output_path"`
}

// LoadConfig - загружает конфигурацию из файла в директории configs и файла .env.
// Возвращает указать на Config и ошибку
func LoadConfig(path string) (*Config, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
		return nil, err
	}

	// Проверка откуда запускает приложение. Локально или из Docker
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	// Указание пути для поиска конфигурационного файла
	viper.AddConfigPath(path)

	// Выбор необходимого файла
	configName := fmt.Sprintf("config.%s", strings.ToLower(env))
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	// Установка префикса для переменных окружения
	viper.SetEnvPrefix("APP")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config *Config
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Println("Config file not found, using environment variables only")
		} else {
			log.Fatalf("Error reading config file, %s", err)
			return nil, err
		}
	}

	err := viper.Unmarshal(&config, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "mapstructure"
		dc.Result = &config
		dc.WeaklyTypedInput = true
	})
	if err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
		return nil, err
	}
	return config, nil
}
