package controllers

import "time"

type GetAvailableRequest struct {
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
}

type CreateNewAppointmentRequest struct {
	UserID   int       `json:"user_id"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
}

type AppointmentResponse struct {
	AppointmentId int       `json:"appointment_id"`
	TrainerId     int       `json:"trainer_id"`
	CustomerId    *int      `json:"customer_id,omitempty"`
	StartsAt      time.Time `json:"starts_at"`
	EndsAt        time.Time `json:"ends_at"`
	Name          string    `json:"name"`
}

type ScheduledAppointmentResponse struct {
	AppointmentID int `json:"appointment_id"`
}
