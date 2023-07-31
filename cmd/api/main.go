package main

import (
	"flag"
	"fmt"
	"infinitybottle.islandwind.me/internal/vcs"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	version = vcs.Version()
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

// @title Infinity Bottle API
// @version 1.0
// @description This is a REST API built to keep track of whisky infinity bottles and their history
// @termsOfService http://swagger.io/terms/
// @license.name MIT
func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	displayVersion := flag.Bool("version", false, "Display version and exit")

	flag.Parse()

	if *displayVersion {
		fmt.Printf("Version: %s\n", version)
		os.Exit(0)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
