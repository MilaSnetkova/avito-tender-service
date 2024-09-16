package tenders

import (
	"context"

	"tender-service/internal/models"
)

func (c *tendersCtrl) PublishTender(ctx context.Context, tenderID int) (*models.Tender, error) {
	tender, err := c.repo.Tenders.Publish(ctx, tenderID)
	if err != nil {
		return nil, err
	}

	return tender, nil
}
