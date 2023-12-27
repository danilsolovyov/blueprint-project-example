package main

import (
    "context"
    "os/signal"
	"syscall"

    "blueprint-project-example/internal/app"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    "github.com/caarlos0/env/v8"
)

func main() {
    cfgLogger := zap.NewDevelopmentConfig()
	cfgLogger.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfgLogger.EncoderConfig.ConsoleSeparator = "    "

	logger, err := cfgLogger.Build()
	if err != nil {
		panic(err)
	}

    ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer cancel()

    var config app.Config
    
    err = env.Parse(config)
	if err != nil {
		panic(err)
	}

    if err := app.Run(ctx, config, logger); err != nil {
        logger.Fatal("Failed to run app", zap.Any("Error", err))
    }
}

