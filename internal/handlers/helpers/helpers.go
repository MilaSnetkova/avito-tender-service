package helpers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetURLParamsTenderID(r *http.Request) string {
	return chi.URLParam(r, "tenderID")
}

func GetURLParamsOfferID(r *http.Request) string {
	return chi.URLParam(r, "offerID")
}
