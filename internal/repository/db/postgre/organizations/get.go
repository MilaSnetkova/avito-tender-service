package organizations

import (
	"context"
	_ "embed"
	"fmt"

	"tender-service/internal/models"

	"github.com/jackc/pgx/v5"
)

//go:embed queries/get.sql
var get string

func (r *organizationsRepo) Get(ctx context.Context, id int) (*models.Organization, error) {
	row := r.pool.QueryRow(ctx, get, id)

	var org models.Organization
	err := row.Scan(&org.ID, &org.Name, &org.Description, &org.Type, &org.CreatedAt, &org.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching organization: %w", err)
	}
	return &org, nil
}
