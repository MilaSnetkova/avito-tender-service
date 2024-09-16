package tenders

import (
	"context"

	"tender-service/internal/models"
)

func (c *tendersCtrl) UpdateTender(ctx context.Context, tender *models.Tender) (*models.Tender, error) {
	tender, err := c.repo.Tenders.Update(ctx, tender)

	if err != nil {
		return nil, err
	}

	return tender, nil
}
