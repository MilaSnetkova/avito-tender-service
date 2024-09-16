package tenders

import (
	"context"
	_ "embed"
	"fmt"

	"tender-service/internal/models"

	"github.com/jackc/pgx/v5"
)

//go:embed queries/get_tender.sql
var getTender string

func (r *tenderRepo) Get(ctx context.Context, id int) (*models.Tender, error) {
	var tender models.Tender

	row := r.pool.QueryRow(ctx, getTender, id)
	err := row.Scan(&tender.ID, &tender.Title, &tender.Description, &tender.Status, &tender.Version, &tender.CreatedAt, &tender.UpdatedAt, &tender.OrganizationID)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	return &tender, nil
}
