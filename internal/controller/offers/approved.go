package offers

import (
	"context"

	"tender-service/internal/models"
)

func (c *offersCtrl) ApproveOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	offer, err := c.repo.Offers.Approve(ctx, offer)
	if err != nil {
		return nil, err
	}
	return offer, nil
}
