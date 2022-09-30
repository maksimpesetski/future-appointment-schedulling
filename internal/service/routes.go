package service

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/maksimpesetski/future-appointment-schedulling/internal/controllers"
)

func (s *Service) routes(controller *controllers.AppointmentsController) {
	router := chi.NewRouter()

	router.Use(middleware.Timeout(time.Duration(3000) * time.Millisecond))

	// temporarily disabled accessLogger for the routs below
	router.Group(func(r chi.Router) {

		r.Route("/appointments/{trainerId}", func(r chi.Router) {
			r.Post("/", controller.ScheduleAppointment)
			r.Get("/scheduled", controller.GetScheduledAppointments)
			r.Get("/available", controller.GetAvailableAppointments)
		})
	})

	s.router = router
}
