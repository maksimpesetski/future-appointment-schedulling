package service

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/maksimpesetski/future-appointment-schedulling/internal/controllers"
	"github.com/maksimpesetski/future-appointment-schedulling/internal/storage"
)

type Service struct {
	logger *zap.Logger
	router *chi.Mux
}

func NewService(logger *zap.Logger, db storage.DBProcessor) (*Service, error) {

	s := &Service{logger: logger}

	appController, err := controllers.NewAppointmentsController(logger, db)
	if err != nil {
		return nil, err
	}

	s.routes(appController)

	return s, nil
}
