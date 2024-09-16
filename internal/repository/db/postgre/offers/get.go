package offers

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/create_offer.sql
var getOffer string

func (r *offerRepo) Get(ctx context.Context, id int) (*models.Offer, error) {
	var offer models.Offer

	row := r.pool.QueryRow(ctx, getOffer, id)
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

	return &offer, nil
}
