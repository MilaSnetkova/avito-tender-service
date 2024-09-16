package tenders

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/update_tender.sql
var updateTender string

func (r *tenderRepo) Update(ctx context.Context, tender *models.Tender) (*models.Tender, error) {
	var updatedTender models.Tender

	row := r.pool.QueryRow(ctx, updateTender, tender.Title, tender.Description, tender.ID)
	err := row.Scan(
		&updatedTender.ID,
		&updatedTender.Title,
		&updatedTender.Description,
		&updatedTender.Status,
		&updatedTender.Version,
		&updatedTender.UpdatedAt,
		&updatedTender.OrganizationID,
	)
	if err != nil {
		return nil, err
	}

	return &updatedTender, nil
}
