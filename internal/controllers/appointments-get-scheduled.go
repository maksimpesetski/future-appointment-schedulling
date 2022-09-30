package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func (c AppointmentsController) GetScheduledAppointments(w http.ResponseWriter, r *http.Request) {
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

	// query db for appointments within the passed time range and customer id NOT NULL
	apps, err := c.db.GetScheduledAppointments(ctx, trainerID)
	if err != nil {
		c.logger.Error("unable to query scheduled appointments", zap.Error(err))
		http.Error(w, "unable to query scheduled appointments", http.StatusInternalServerError)
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
