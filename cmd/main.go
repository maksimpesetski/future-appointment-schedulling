package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/maksimpesetski/future-appointment-schedulling/internal/service"
)

/*
 I built this service assuming that trainer has already added his/her availability to DB, meaning that it's guaranteed
 we have the most up-to-date available appointments at the time of incoming request.

Areas of improvement:
 - logging
 - routing middleware that checks if url "trainer_id" is valid
 - separate  appointment validation logic into internal "business_logic" pkg

*/
func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("unable to initialize logger: %v", err)
	}

	// server
	s, err := service.NewService(logger)
	if err != nil {
		logger.Fatal("unable create service instance", zap.Error(err))
	}

	s.Start()
}
