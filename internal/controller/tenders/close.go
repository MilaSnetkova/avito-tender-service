package tenders

import (
	"context"

	"tender-service/internal/models"
)

func (c *tendersCtrl) CloseTender(ctx context.Context, tenderID int) (*models.Tender, error) {
	tender, err := c.repo.Tenders.Close(ctx, tenderID)

	if err != nil {
		return nil, err
	}

	return tender, nil
}
