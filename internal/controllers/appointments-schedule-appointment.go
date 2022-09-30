package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

var offDays = map[time.Weekday]bool{
	time.Sunday:   true,
	time.Saturday: true,
}

// ScheduleAppointment accepts customer's desired timeframes and books an appointment with trainer
func (c AppointmentsController) ScheduleAppointment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	trainerIDParam := chi.URLParam(r, "trainerId")
	if trainerIDParam == "" {
		c.logger.Error("unable to extract trainer ID")
		http.Error(w, "invalid trainer id", http.StatusBadRequest)
		return
	}

	trainerID, err := strconv.Atoi(trainerIDParam)
	if err != nil {
		c.logger.Error("missing trainer id")
		http.Error(w, "invalid trainer id", http.StatusBadRequest)
		return
	}

	// validate trainerID
	if trainerID == 0 {
		c.logger.Error("invalid trainer id")
		http.Error(w, "invalid trainer id", http.StatusBadRequest)
		return
	}

	payload := CreateNewAppointmentRequest{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		c.logger.Error("unable to decode request payload")
		http.Error(w, "unable to decode request payload", http.StatusBadRequest)
		return
	}

	// sync time with PS timezone
	startAt := payload.StartsAt.In(c.location)
	endAt := payload.EndsAt.In(c.location)

	// check if dates in future or start < end
	if time.Now().In(c.location).After(startAt) || startAt.After(endAt) || startAt.Equal(endAt) {
		c.logger.Error("invalid start/end times")
		http.Error(w, "invalid start/end times", http.StatusBadRequest)
		return
	}

	// check if in Mon - Fri range
	if offDays[startAt.Weekday()] && offDays[endAt.Weekday()] {
		c.logger.Error("unable to schedule on weekend")
		http.Error(w, "unable to schedule on weekend", http.StatusBadRequest)
		return
	}

	// grab hour, min, second
	startHH, startMM, startSS := startAt.Clock()
	endHH, _, _ := startAt.Clock()

	// appointment before 8am PST
	if startHH < 8 || endHH < 8 {
		c.logger.Error("unable to schedule appointment before hours")
		http.Error(w, "unable to schedule appointment before hours", http.StatusBadRequest)
		return
	}

	// appointment after 5pm PST
	if startHH > 17 || endHH > 17 {
		c.logger.Error("unable to schedule appointment after hours")
		http.Error(w, "unable to schedule appointment after hours", http.StatusBadRequest)
		return
	}

	// check if start at --:00:00 or --:30.00 minutes
	if startMM != 0 && startMM != 30 || startSS != 0 {
		c.logger.Error("start and end time must be rounded to 00 or 30")
		http.Error(w, "start and end time must be rounded to 00 or 30", http.StatusBadRequest)
		return
	}

	// check if start + 30 min == end; also covers if end is not --:00:00 or --:30:00
	if startAt.Add(30*time.Minute) != endAt {
		c.logger.Error("session must be 30 minutes")
		http.Error(w, "session must be 30 minutes", http.StatusBadRequest)
		return
	}

	// query db for available appointments within validated time-range
	availableTimes, err := c.db.GetAvailableAppointments(ctx, trainerID, startAt, endAt)
	if err != nil {
		c.logger.Error("unable to query available appointments", zap.Error(err))
		http.Error(w, "unable to query available appointments", http.StatusBadRequest)
		return
	}

	// booked or unavailable
	if len(availableTimes) == 0 {
		c.logger.Error("selected appointment is unavailable")
		http.Error(w, "selected appointment is unavailable", http.StatusBadRequest)
		return

	} else {
		// ad this point we should have exactly one existing non-booked appointment
		if err := c.db.BookAppointment(availableTimes[0].AppointmentId, payload.UserID); err != nil {
			c.logger.Error("unable to book appointment", zap.Error(err))
			http.Error(w, "unable to book appointment", http.StatusBadRequest)
			return
		}
	}

	response := ScheduledAppointmentResponse{
		AppointmentID: availableTimes[0].AppointmentId,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		c.logger.Error("unable to encode response data", zap.Error(err))
		http.Error(w, "unable to encode response data", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBytes)
}
