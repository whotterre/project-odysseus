package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server      ServerConfig `mapstructure:"server"`
	Redis       RedisConfig  `mapstructure:"redis"`
	DatabaseURL string       `mapstructure:"DATABASE_URL"`
	JWTSecret   string       `mapstructure:"JWT_SECRET"`
}

type ServerConfig struct {
	Env      string `mapstructure:"env"`
	HTTPPort string `mapstructure:"http_port"`
	GRPCPort string `mapstructure:"grpc_port"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func LoadConfig() (*Config, error) {
	projectRoot, err := findProjectRoot()
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(filepath.Join(projectRoot, ".env"))

	viper.SetDefault("server.env", "development")
	viper.SetDefault("server.http_port", "8080")
	viper.SetDefault("server.grpc_port", "50051")
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.db", 0)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
		log.Println("No .env file found. Falling back entirely to system environment variables and defaults.")
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if viper.IsSet("SERVER_ENV") {
		viper.Set("server.env", viper.Get("SERVER_ENV"))
	}
	if viper.IsSet("SERVER_HTTP_PORT") {
		viper.Set("server.http_port", viper.Get("SERVER_HTTP_PORT"))
	}
	if viper.IsSet("SERVER_GRPC_PORT") {
		viper.Set("server.grpc_port", viper.Get("SERVER_GRPC_PORT"))
	}

	if viper.IsSet("REDIS_HOST") {
		viper.Set("redis.host", viper.Get("REDIS_HOST"))
	}
	if viper.IsSet("REDIS_PORT") {
		viper.Set("redis.port", viper.Get("REDIS_PORT"))
	}
	if viper.IsSet("REDIS_PASSWORD") {
		viper.Set("redis.password", viper.Get("REDIS_PASSWORD"))
	}
	if viper.IsSet("REDIS_DB") {
		viper.Set("redis.db", viper.Get("REDIS_DB"))
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		} else if !os.IsNotExist(err) {
			return "", err
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("project root not found")
		}
		dir = parent
	}
}
