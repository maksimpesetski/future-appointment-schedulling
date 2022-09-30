package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/maksimpesetski/future-appointment-schedulling/internal/service"
	"github.com/maksimpesetski/future-appointment-schedulling/internal/storage"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("unable to initialize logger: %v", err)
	}

	// init DB
	processor, err := storage.NewDBProcessor()
	if err != nil {
		logger.Fatal("unable to open DB connection", zap.Error(err))
	}
	err = processor.Ping()
	if err != nil {
		logger.Fatal("unable to ping DB", zap.Error(err))
	}
	defer processor.Close()

	// server
	s, err := service.NewService(logger, processor)
	if err != nil {
		logger.Fatal("unable create service instance", zap.Error(err))
	}

	s.Start()
}
