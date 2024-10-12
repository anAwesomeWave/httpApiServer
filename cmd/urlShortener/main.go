package main

import (
	"fmt"
	"log/slog"
	"os"
	"urlShortener/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	configPath := parseFlags()
	cfg := config.MustLoad(*configPath)

	fmt.Println(cfg)
	// logger -
	log := setupLogger(cfg.Env)
	log.Info("Logger created.", slog.String("env", cfg.Env))
	log.Debug("Starting url-shortener")
	// storage - sqlite

	// router - chi, render
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		panic(fmt.Sprintf("Wrong Env value {%v}", env))
	}
	return log
}
