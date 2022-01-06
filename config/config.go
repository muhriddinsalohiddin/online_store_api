package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	CatalogServiceHost string
	CatalogServicePort int
	
	OrderServiceHost string
	OrderServicePort int

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.CatalogServiceHost = cast.ToString(getOrReturnDefault("CatalogService_HOST", "localhost"))
	c.CatalogServicePort = cast.ToInt(getOrReturnDefault("CatalogService_PORT", 9005))

	c.OrderServiceHost = cast.ToString(getOrReturnDefault("OrderService_HOST", "localhost"))
	c.OrderServicePort = cast.ToInt(getOrReturnDefault("OrderService_PORT", 9006))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
