package config_test

import (
	"os"
	"testing"

	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/config"
)

func TestLoadConfig_Default(t *testing.T) {
	// Ensure the env variable is not set.
	os.Unsetenv("MAX_CACHE_SIZE")
	cfg := config.LoadConfig()
	if cfg.MaxCacheSize != 10 {
		t.Errorf("expected default max cache size 10, got %d", cfg.MaxCacheSize)
	}
}

func TestLoadConfig_Custom(t *testing.T) {
	os.Setenv("MAX_CACHE_SIZE", "20")
	defer os.Unsetenv("MAX_CACHE_SIZE")

	cfg := config.LoadConfig()
	if cfg.MaxCacheSize != 20 {
		t.Errorf("expected max cache size 20, got %d", cfg.MaxCacheSize)
	}
}

func TestLoadConfig_Invalid(t *testing.T) {
	os.Setenv("MAX_CACHE_SIZE", "-5")
	defer os.Unsetenv("MAX_CACHE_SIZE")

	cfg := config.LoadConfig()
	// For an invalid (non-positive) value, the default should be used.
	if cfg.MaxCacheSize != 10 {
		t.Errorf("expected default max cache size 10 for invalid value, got %d", cfg.MaxCacheSize)
	}
}
