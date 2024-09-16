package offers

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/cancel_offer.sql
var publishOffer string

func (r *offerRepo) Publish(ctx context.Context, id int) (*models.Offer, error) {
	var publishedOffer models.Offer

	row := r.pool.QueryRow(ctx, publishOffer, id)
	err := row.Scan(
		&publishedOffer.ID,
		&publishedOffer.Description,
		&publishedOffer.Status,
		&publishedOffer.Version,
		&publishedOffer.CreatedAt,
		&publishedOffer.UpdatedAt,
		&publishedOffer.TenderID,
		&publishedOffer.UserID,
	)
	if err != nil {
		return nil, err
	}

	return &publishedOffer, nil
}
