package tenders

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tender-service/internal/controller/tenders"
	"tender-service/internal/handlers/helpers"
	"tender-service/internal/models"
)

type TendersHandler struct {
	ctrl tenders.Tenders
}

func NewTendersHandler(ctrl tenders.Tenders) *TendersHandler {
	return &TendersHandler{ctrl: ctrl}
}

func (h *TendersHandler) CreateTender(w http.ResponseWriter, r *http.Request) {
	tender, err := h.ctrl.CreateTender(r.Context(), r.Body)
	if err != nil {
		helpers.Response(w, nil, "Error creating tender", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, tender, "Tender created successfully", nil, http.StatusCreated)
}

func (h *TendersHandler) PublishTender(w http.ResponseWriter, r *http.Request) {
	tenderID, err := strconv.Atoi(r.URL.Query().Get("tenderId"))
	if err != nil {
		helpers.Response(w, nil, "Invalid tender ID", err, http.StatusBadRequest)
		return
	}

	tender, err := h.ctrl.PublishTender(r.Context(), tenderID)
	if err != nil {
		helpers.Response(w, nil, "Error publishing tender", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, tender, "Tender published successfully", nil, http.StatusOK)
}

func (h *TendersHandler) CloseTender(w http.ResponseWriter, r *http.Request) {
	tenderID, err := strconv.Atoi(r.URL.Query().Get("tenderId"))
	if err != nil {
		helpers.Response(w, nil, "Invalid tender ID", err, http.StatusBadRequest)
		return
	}

	tender, err := h.ctrl.CloseTender(r.Context(), tenderID)
	if err != nil {
		helpers.Response(w, nil, "Error closing tender", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, tender, "Tender closed successfully", nil, http.StatusOK)
}

func (h *TendersHandler) UpdateTender(w http.ResponseWriter, r *http.Request) {
	var tender models.Tender
	if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
		helpers.Response(w, nil, "Invalid request body", err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	updatedTender, err := h.ctrl.UpdateTender(r.Context(), &tender)
	if err != nil {
		helpers.Response(w, nil, "Error updating tender", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, updatedTender, "Tender updated successfully", nil, http.StatusOK)
}
