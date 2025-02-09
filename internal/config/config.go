package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds configuration for the application.
type Config struct {
	MaxCacheSize int
}

// LoadConfig reads config from environment variables and returns a Config struct.
func LoadConfig() *Config {
	maxCacheSize := 10 // default value
	if val, exists := os.LookupEnv("MAX_CACHE_SIZE"); exists {
		if parsed, err := strconv.Atoi(val); err == nil && parsed > 0 {
			maxCacheSize = parsed
		} else {
			fmt.Printf("Invalid MAX_CACHE_SIZE value '%s'. Using default: %d\n", val, maxCacheSize)
		}
	}
	return &Config{
		MaxCacheSize: maxCacheSize,
	}
}
