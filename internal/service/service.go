package service

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	router *chi.Mux
}

func NewService(logger *zap.Logger) (*Service, error) {

	s := &Service{logger: logger}

	s.routes()

	return s, nil
}
