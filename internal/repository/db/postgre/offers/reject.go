package offers

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/reject_offer.sql
var rejectOffer string

func (r *offerRepo) Reject(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	_, err := r.pool.Exec(ctx, rejectOffer, offer)
	if err != nil {
		return nil, err
	}
	return offer, nil

}
