package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/jsonapi/internal/leveledlog"
	"example.com/jsonapi/internal/server"
	"example.com/jsonapi/internal/version"
)

type config struct {
	addr    string
	baseURL string
	env     string
	version bool
}

type application struct {
	config config
	logger *leveledlog.Logger
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", "localhost:4444", "server address to listen on")
	flag.StringVar(&cfg.baseURL, "base-url", "http://localhost:4444", "base URL for the application")
	flag.StringVar(&cfg.env, "env", "development", "operating environment: development, testing, staging or production")
	flag.BoolVar(&cfg.version, "version", false, "display version and exit")

	flag.Parse()

	if cfg.version {
		fmt.Printf("version: %s\n", version.Get())
		return
	}

	logger := leveledlog.NewLogger(os.Stdout, leveledlog.LevelAll, true)

	app := &application{
		config: cfg,
		logger: logger,
	}

	logger.Info("starting server on %s (version %s)", cfg.addr, version.Get())

	err := server.Run(cfg.addr, app.routes())
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("server stopped")
}
