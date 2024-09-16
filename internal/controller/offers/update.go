package offers

import (
	"context"

	"tender-service/internal/models"
)

func (c *offersCtrl) UpdateOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	offer, err := c.repo.Offers.Update(ctx, offer)

	if err != nil {
		return nil, err
	}

	return offer, nil
}
