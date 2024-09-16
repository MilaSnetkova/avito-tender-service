package tenders

import (
	"context"
	_ "embed"

	"tender-service/internal/models"
)

//go:embed queries/publish_tender.sql
var publishTender string

func (r *tenderRepo) Publish(ctx context.Context, id int) (*models.Tender, error) {
	var publishedTender models.Tender

	row := r.pool.QueryRow(ctx, publishTender, id)
	err := row.Scan(
		&publishedTender.ID,
		&publishedTender.Title,
		&publishedTender.Description,
		&publishedTender.Status,
		&publishedTender.Version,
		&publishedTender.UpdatedAt,
		&publishedTender.OrganizationID,
	)
	if err != nil {
		return nil, err
	}

	return &publishedTender, nil
}
