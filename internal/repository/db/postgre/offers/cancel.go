package offers

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/cancel_offer.sql
var cancelOffer string

func (r *offerRepo) Cancel(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	_, err := r.pool.Exec(ctx, cancelOffer, offer.ID)
	if err != nil {
		return nil, err
	}
	return offer, nil

}
