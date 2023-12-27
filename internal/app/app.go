package app

import (
	"context"

	"go.uber.org/zap"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"blueprint-project-example"`
}

func Run(ctx context.Context, config Config, logger *zap.Logger) error {
	logger.Info("App started", zap.Any("config", config))
	
	<-ctx.Done()

	logger.Info("App shutting down gracefully")

	return nil
}