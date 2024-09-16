package tenders

import (
	"context"
	_ "embed"
	"fmt"

	"tender-service/internal/models"
)

//go:embed queries/create_tender.sql
var createTender string

func (r *tenderRepo) Create(ctx context.Context, tender *models.Tender) (*models.Tender, error) {
	var newTender models.Tender

	row := r.pool.QueryRow(
		ctx,
		createTender,
		tender.Title,
		tender.Description,
		tender.Status,
		tender.Version,
		tender.OrganizationID,
	)

	err := row.Scan(&newTender.ID, &newTender.Title, &newTender.Description, &newTender.Status, &newTender.Version, &newTender.CreatedAt, &newTender.UpdatedAt, &newTender.OrganizationID)

	if err != nil {
		return nil, fmt.Errorf("Error creating tender: %w", err)
	}

	return &newTender, nil
}
