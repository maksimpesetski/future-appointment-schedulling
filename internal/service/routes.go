package service

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Service) routes() {
	router := chi.NewRouter()

	router.Use(middleware.Timeout(time.Duration(3000) * time.Millisecond))

	// temporarily disabled accessLogger for the routs below
	router.Group(func(r chi.Router) {

	})

	s.router = router
}
