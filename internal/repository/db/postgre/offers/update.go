package offers

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/update_offer.sql
var updateOffer string

func (r *offerRepo) Update(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	_, err := r.pool.Exec(ctx, updateOffer, offer.Description, offer.Status, offer.Version, offer.ID)
	if err != nil {
		return nil, err
	}
	return offer, err
}
