package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func (c AppointmentsController) GetAvailableAppointments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	trainerIDParam := chi.URLParam(r, "trainerId")
	if trainerIDParam == "" {
		c.logger.Error("unable to extract trainer ID")
		http.Error(w, "invalid trainer id", http.StatusBadRequest)
		return
	}

	// validate trainerID
	trainerID, err := strconv.Atoi(trainerIDParam)
	if err != nil {
		c.logger.Error("missing trainer id")
		http.Error(w, "invalid trainer id", http.StatusBadRequest)
		return
	}

	if trainerID == 0 {
		c.logger.Error("invalid trainer id")
		http.Error(w, "invalid trainer id", http.StatusBadRequest)
		return
	}

	payload := GetAvailableRequest{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		c.logger.Error("unable to decode request payload")
		http.Error(w, "unable to decode request payload", http.StatusBadRequest)
		return
	}

	// sync time with PS timezone
	startAt := payload.StartsAt.In(c.location)
	endAt := payload.EndsAt.In(c.location)

	// check if startAt is in future and start < end
	if time.Now().In(c.location).After(startAt) || startAt.After(startAt) {
		c.logger.Error("invalid start or end time")
		http.Error(w, "invalid start or end time", http.StatusBadRequest)
		return
	}

	// db query to fetch available appointments within the time range
	apps, err := c.db.GetAvailableAppointments(ctx, trainerID, startAt, endAt)
	if err != nil {
		c.logger.Error("unable to query available appointments", zap.Error(err))
		http.Error(w, "unable to query available appointments", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(apps)
	if err != nil {
		c.logger.Error("unable to encode response payload", zap.Error(err))
		http.Error(w, "unable to encode response payload", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
