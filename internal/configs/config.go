package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ServiceName string
	Environment string
	LoggerLevel string
	HTTPHost    string
	HTTPPort    string

	AuthServiceGrpcHost   string
	AuthServiceGrpcPort   string
	CarpetServiceGrpcHost string
	CarpetServiceGrpcPort string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	config := Config{}

	config.HTTPPort = cast.ToString(coalesce("HTTP_Port", ":8082"))

	config.HTTPHost = cast.ToString(coalesce("HTTP_HOST", "localhost"))
	config.HTTPPort = cast.ToString(coalesce("HTTP_PORT", ":8085"))

	config.AuthServiceGrpcPort = cast.ToString(coalesce("AUTH_SERVICE_GRPC_HOST", "localhost"))
	config.AuthServiceGrpcPort = cast.ToString(coalesce("AUTH_SERVICE_GRPC_PORT", ":8081"))

	config.CarpetServiceGrpcHost = cast.ToString(coalesce("CARPET_SERVICE_GRPC_HOST", "localhost"))
	config.CarpetServiceGrpcPort = cast.ToString(coalesce("CARPET_SERVICE_GRPC_PORT", ":8082"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}
	return defaultValue
}
