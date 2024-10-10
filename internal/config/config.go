package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string     `yaml:"storagePath" env-required:"true"`
	HTTPServer  HttpServer `yaml:"httpServer"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8082"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idleTimeout" env-default:"60s"`
}

func MustLoad(configPath string) *Config {
	if _, err := os.Stat(configPath); err != nil {
		//panic(fmt.Sprintf("ConfigPath error: %v", err))
		log.Fatalf("ConfigPath error: %v", err)
	}

	var conf Config

	if err := cleanenv.ReadConfig(configPath, &conf); err != nil {
		log.Fatalf("Cannot read from confi file: %v", err)
	}

	return &conf
}
