package storage

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/maksimpesetski/future-appointment-schedulling/internal/models"
)

// GetAvailableAppointments queries appointments db for appointments that are in time range and customer_id is NULL(available)
func (p processor) GetAvailableAppointments(ctx context.Context, trainerID int, startTime, endTime time.Time) ([]models.Appointment, error) {
	appointments := []models.Appointment{}

	rows, err := p.sq.
		Select("appointment_id", "starts_at", "ends_at", "name", "trainer_id", "customer_id").
		From("appointments").
		Where(sq.And{
			sq.GtOrEq{"starts_at": startTime},
			sq.LtOrEq{"ends_at": endTime},
			sq.Eq{"trainer_id": trainerID},
			sq.Eq{"customer_id": nil},
		}).
		RunWith(p.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment models.Appointment
		err := rows.Scan(
			&appointment.AppointmentId,
			&appointment.StartsAt,
			&appointment.EndsAt,
			&appointment.Name,
			&appointment.TrainerId,
			&appointment.CustomerId)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// GetScheduledAppointments queries appointments db for appointments that are in time range and customer_id is NOT NULL(booked)
func (p processor) GetScheduledAppointments(ctx context.Context, trainerID int) ([]models.Appointment, error) {
	appointments := []models.Appointment{}

	rows, err := p.sq.
		Select("appointment_id", "starts_at", "ends_at", "name", "trainer_id", "customer_id").
		From("appointments").
		Where(sq.And{
			sq.Eq{"trainer_id": trainerID},
			sq.NotEq{"customer_id": nil},
		}).
		RunWith(p.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment models.Appointment
		err := rows.Scan(
			&appointment.AppointmentId,
			&appointment.StartsAt,
			&appointment.EndsAt,
			&appointment.Name,
			&appointment.TrainerId,
			&appointment.CustomerId)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// BookAppointment queries db to mark the selected appointment as booked (assign customerID)
func (p processor) BookAppointment(appointmentID, customerId int) error {

	_, err := p.sq.Update("appointments").
		SetMap(map[string]interface{}{
			"customer_id": customerId,
		}).
		Where(sq.Eq{"appointment_id": appointmentID}).
		RunWith(p.db).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
