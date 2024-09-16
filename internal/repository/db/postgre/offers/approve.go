package offers

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/approve_offer.sql
var approveOffer string

func (r *offerRepo) Approve(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	_, err := r.pool.Exec(ctx, approveOffer, offer.ID)
	if err != nil {
		return nil, err
	}
	return offer, nil
}
