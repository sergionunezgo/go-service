package service

import (
	"github.com/sergionunezgo/go-reuse/v2/pkg/logger"
	"github.com/sergionunezgo/go-service/app/service/http"
)

// Config defines the values that can be loaded from env vars or other config files.
type Config struct {
	Port     int
	LogLevel string
}

// Service defines the methods that are required to operate a web service.
type Service interface {
	Start() error
	Close() error
}

// New will return a Service that can be used to handle client requests.
func New(cfg *Config) Service {
	logger.Log.Info("initializing api service")
	return http.NewService(cfg.Port)
}
