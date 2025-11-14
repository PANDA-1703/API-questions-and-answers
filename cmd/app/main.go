package main

import (
	"API-quest-answ/internal/config"
	"API-quest-answ/internal/infrastructure/db/postgres"
	"API-quest-answ/pkg/logger"
	"flag"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "", "")
	flag.Parsed()
	cfg, err := config.Init(cfgPath, false)
	if err != nil {
		panic(err)
	}

	logger := logger.New()

	pool, err := postgres.New(cfg.Postgres)
	if err != nil {
		panic(err)
	}

}
