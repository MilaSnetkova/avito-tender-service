package tenders

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/close_tender.sql

var closeTender string

func (r *tenderRepo) Close(ctx context.Context, id int) (*models.Tender, error) {
	var closedTender models.Tender

	row := r.pool.QueryRow(ctx, closeTender, id)
	err := row.Scan(
		&closedTender.ID,
		&closedTender.Title,
		&closedTender.Description,
		&closedTender.Status,
		&closedTender.Version,
		&closedTender.UpdatedAt,
		&closedTender.OrganizationID,
	)
	if err != nil {
		return nil, err
	}

	return &closedTender, nil
}
