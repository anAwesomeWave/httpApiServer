package main

import (
	"fmt"
	"log/slog"
	"os"
	"urlShortener/internal/config"
	"urlShortener/internal/lib/logger/slg"
	"urlShortener/internal/storage/sqlite"
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
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage", slg.Err(err))
		os.Exit(1)
	}
	_ = storage
	//id, err := storage.SaveURL("https://www.google.com/", "myAlias")
	//if err != nil {
	//	log.Error("Failed to insert", slg.Err(err))
	//	os.Exit(1)
	//}
	//log.Info("Saved url with id: ", slog.String("id", strconv.FormatInt(id, 10)))
	//id, err = storage.SaveURL("https://www.google.com/", "myNewAlias")
	//if err != nil {
	//	log.Error("Failed to insert", slg.Err(err))
	//	os.Exit(1)
	//}
	//log.Info("Saved url with id: ", id)
	//id, err = storage.SaveURL("https://www.google.com/", "myAlias")
	//if err != nil {
	//	log.Error("Failed to insert", slg.Err(err))
	//	os.Exit(1)
	//}
	//log.Info("Saved url with id: ", id)
	//url, err := storage.GetURL("myAlias")
	//if err != nil {
	//	log.Error("Failed to get url by alias", slg.Err(err))
	//	os.Exit(1)
	//}
	//log.Info("Url name from alias", slog.String("url", url))
	//id, err = storage.DeleteURL("myAliass")
	//if err != nil {
	//	log.Error("Failed to get url by alias", slg.Err(err))
	//	os.Exit(1)
	//}
	//log.Info("DELETED", slog.String("id", strconv.FormatInt(id, 10)))

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
