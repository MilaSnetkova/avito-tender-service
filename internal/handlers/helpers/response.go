package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"tender-service/internal/logger"

	"github.com/jackc/pgx/v5"
)

func Response(w http.ResponseWriter, data interface{}, errName string, err error, status int) {
	w.Header().Set("Content-Type", "application/json")

	if errors.Is(err, pgx.ErrNoRows) {
		msg := "No records found"
		logger.Info(msg)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	if errors.Is(err, errors.New("parsing error")) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, errors.New("decoding error")) {
		logger.Error("Error decoding request body", "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err != nil {
		logger.Error(fmt.Sprintf("Error %s", errName), "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if err = json.NewEncoder(w).Encode(data); err != nil {
		logger.Error(fmt.Sprintf("Error encoding %s", errName), "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
