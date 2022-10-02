package storage

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lann/builder"

	"github.com/maksimpesetski/future-appointment-schedulling/internal/models"
)

//dbProcessor marks methods that interact with the database, (and live in this package),
//so that they can be easily referenced and mocked elsewhere.
type processor struct {
	db *sql.DB
	sq sq.StatementBuilderType
}

// NewDBProcessor creates a new db client wrapper to query sql database
func NewDBProcessor() (DBProcessor, error) {
	db, err := newSQLConnection("postgres://postgres:postgres@db:5432/postgres?connect_timeout=10&sslmode=disable")
	if err != nil {
		return nil, err
	}
	sqBuilder := sq.StatementBuilderType(builder.EmptyBuilder).PlaceholderFormat(sq.Dollar)
	return &processor{db, sqBuilder}, nil
}

type DBProcessor interface {
	GetAvailableAppointments(ctx context.Context, trainerID int, startTime, endTime time.Time) ([]models.Appointment, error)
	GetScheduledAppointments(ctx context.Context, trainerID int) ([]models.Appointment, error)
	BookAppointment(appointmentID, customerId int) error
	Ping() error
	Close() error
}

func (p processor) Close() error {
	return p.db.Close()
}

func (p processor) Ping() error {
	return p.db.Ping()
}
