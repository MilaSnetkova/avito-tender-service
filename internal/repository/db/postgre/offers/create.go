package offers

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/create_offer.sql
var createOffer string

func (r *offerRepo) Create(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	row := r.pool.QueryRow(ctx, createOffer, offer.Description, offer.Status, offer.Version, offer.TenderID, offer.UserID)

	err := row.Scan(
		&offer.ID,
		&offer.Description,
		&offer.Status,
		&offer.Version,
		&offer.CreatedAt,
		&offer.UpdatedAt,
		&offer.TenderID,
		&offer.UserID,
	)
	if err != nil {
		return nil, err
	}

	return offer, nil
}
