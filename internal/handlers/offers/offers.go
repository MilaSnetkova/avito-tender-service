package offers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tender-service/internal/controller/offers"
	"tender-service/internal/handlers/helpers"
	"tender-service/internal/models"
)

type OffersHandler struct {
	ctrl offers.Offers
}

func NewOffersHandler(ctrl offers.Offers) *OffersHandler {
	return &OffersHandler{ctrl: ctrl}
}

func (h *OffersHandler) ApproveOffer(w http.ResponseWriter, r *http.Request) {
	offerID, err := strconv.Atoi(r.URL.Query().Get("offerId"))
	if err != nil {
		helpers.Response(w, nil, "Invalid offer ID", err, http.StatusBadRequest)
		return
	}

	offer, err := h.ctrl.ApproveOffer(r.Context(), &models.Offer{ID: offerID})
	if err != nil {
		helpers.Response(w, nil, "Error approving offer", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, offer, "Offer approved successfully", nil, http.StatusOK)
}
func (h *OffersHandler) RejectOffer(w http.ResponseWriter, r *http.Request) {
	offerID, err := strconv.Atoi(r.URL.Query().Get("offerId"))
	if err != nil {
		helpers.Response(w, nil, "Invalid offer ID", err, http.StatusBadRequest)
		return
	}

	offer, err := h.ctrl.RejectOffer(r.Context(), &models.Offer{ID: offerID})
	if err != nil {
		helpers.Response(w, nil, "Error rejecting offer", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, offer, "Offer rejected successfully", nil, http.StatusOK)
}

func (h *OffersHandler) CreateOffer(w http.ResponseWriter, r *http.Request) {
	offer, err := h.ctrl.CreateOffer(r.Context(), r.Body)
	if err != nil {
		helpers.Response(w, nil, "Error creating offer", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, offer, "Offer created successfully", nil, http.StatusCreated)
}

func (h *OffersHandler) CancelOffer(w http.ResponseWriter, r *http.Request) {
	offerID, err := strconv.Atoi(r.URL.Query().Get("offerId"))
	if err != nil {
		helpers.Response(w, nil, "Invalid offer ID", err, http.StatusBadRequest)
		return
	}

	offer, err := h.ctrl.CancelOffer(r.Context(), &models.Offer{ID: offerID})
	if err != nil {
		helpers.Response(w, nil, "Error canceling offer", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, offer, "Offer canceled successfully", nil, http.StatusOK)
}

func (h *OffersHandler) UpdateOffer(w http.ResponseWriter, r *http.Request) {
	var offer models.Offer
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		helpers.Response(w, nil, "Invalid request body", err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	updatedOffer, err := h.ctrl.UpdateOffer(r.Context(), &offer)
	if err != nil {
		helpers.Response(w, nil, "Error updating offer", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, updatedOffer, "Offer updated successfully", nil, http.StatusOK)
}

func (h *OffersHandler) PublishOffer(w http.ResponseWriter, r *http.Request) {
	offerID, err := strconv.Atoi(r.URL.Query().Get("offerId"))
	if err != nil {
		helpers.Response(w, nil, "Invalid offer ID", err, http.StatusBadRequest)
		return
	}

	offer, err := h.ctrl.PublishOffer(r.Context(), offerID) // Pass offerID as int
	if err != nil {
		helpers.Response(w, nil, "Error publishing offer", err, http.StatusInternalServerError)
		return
	}

	helpers.Response(w, offer, "Offer published successfully", nil, http.StatusOK)
}
