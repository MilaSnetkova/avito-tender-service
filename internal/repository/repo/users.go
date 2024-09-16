package repo

import (
	"context"

	"tender-service/internal/models"
)

type Users interface {
	Get(ctx context.Context, id int) (*models.User, error)
	Delete(ctx context.Context, id int) error
}
