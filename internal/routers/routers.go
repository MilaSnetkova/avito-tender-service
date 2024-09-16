package routers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"tender-service/internal/controller"
	"tender-service/internal/handlers"
)

func New(v string, ctrl controller.Controller, serverAddress string) *chi.Mux {
	h := handlers.New(ctrl)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Recoverer)

	apiVersioned := fmt.Sprintf("/api/%s", v)

	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	tenderRoutes := func(r chi.Router) {
		r.Post("/", h.Tenders.CreateTender)
		r.Patch("/{tenderID}/edit", h.Tenders.UpdateTender)
		r.Post("/{tenderID}/publish", h.Tenders.PublishTender)
		r.Post("/{tenderID}/close", h.Tenders.CloseTender)
	}

	offerRoutes := func(r chi.Router) {
		r.Post("/new", h.Offers.CreateOffer)
		r.Patch("/{offerID}/edit", h.Offers.UpdateOffer)
		r.Post("/{offerID}/publish", h.Offers.PublishOffer)
		r.Post("/{offerID}/cancel", h.Offers.CancelOffer)
		r.Post("/{offerID}/approve", h.Offers.ApproveOffer)
	}

	r.Route(apiVersioned+"/tenders", tenderRoutes)
	r.Route(apiVersioned+"/offers", offerRoutes)

	return r
}
