package handlers

import (
	"tender-service/internal/controller"
	"tender-service/internal/handlers/offers"
	"tender-service/internal/handlers/tenders"
)

type Handler struct {
	Tenders *tenders.TendersHandler
	Offers  *offers.OffersHandler
}

func New(controller controller.Controller) *Handler {
	tendersHandler := tenders.NewTendersHandler(controller)
	offersHandler := offers.NewOffersHandler(controller)

	return &Handler{
		Tenders: tendersHandler,
		Offers:  offersHandler,
	}
}
