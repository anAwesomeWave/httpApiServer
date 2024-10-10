package main

import (
	"fmt"
	"urlShortener/internal/config"
)

func main() {
	configPath := parseFlags()
	cfg := config.MustLoad(*configPath)

	fmt.Println(cfg)
	// logger -

	// storage - sqlite

	// router - chi, render
}
