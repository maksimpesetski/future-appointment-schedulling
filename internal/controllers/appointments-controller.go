package controllers

import (
	"time"

	"go.uber.org/zap"

	"github.com/maksimpesetski/future-appointment-schedulling/internal/storage"
)

type AppointmentsController struct {
	logger   *zap.Logger
	db       storage.DBProcessor
	location *time.Location
}

func NewAppointmentsController(
	logger *zap.Logger,
	db storage.DBProcessor,
) (*AppointmentsController, error) {

	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		logger.Error("unable to load PST timezone")
		return nil, err
	}

	return &AppointmentsController{
		logger,
		db,
		location,
	}, nil
}
