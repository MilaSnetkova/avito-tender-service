package repo

import (
	"context"

	"tender-service/internal/models"
)

type Organizations interface {
	Get(ctx context.Context, id int) (*models.Organization, error)
	Delete(ctx context.Context, id int) error
}
