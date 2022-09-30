package models

import "time"

type Appointment struct {
	AppointmentId int       `json:"appointment_id"`
	TrainerId     int       `json:"trainer_id"`
	CustomerId    *int      `json:"customer_id,omitempty"` // nil means the appointment is available for booking
	StartsAt      time.Time `json:"starts_at"`
	EndsAt        time.Time `json:"ends_at"`
	Name          string    `json:"name"`
}
