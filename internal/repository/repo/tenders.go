package repo

import (
	"context"

	"tender-service/internal/models"
)

type Tenders interface {
	Create(ctx context.Context, tender *models.Tender) (*models.Tender, error)
	Get(ctx context.Context, id int) (*models.Tender, error)
	Publish(ctx context.Context, id int) (*models.Tender, error)
	Close(ctx context.Context, id int) (*models.Tender, error)
	Update(ctx context.Context, tender *models.Tender) (*models.Tender, error)
}
