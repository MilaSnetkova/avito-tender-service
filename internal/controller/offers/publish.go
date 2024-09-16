package offers

import (
	"context"

	"tender-service/internal/models"
)

func (c *offersCtrl) PublishOffer(ctx context.Context, offerID int) (*models.Offer, error) {
	offer, err := c.repo.Offers.Publish(ctx, offerID)

	if err != nil {
		return nil, err
	}

	return offer, nil
}
