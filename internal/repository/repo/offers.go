package repo

import (
	"context"
	"tender-service/internal/models"
)

type Offers interface {
	Create(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	Get(ctx context.Context, id int) (*models.Offer, error)
	Approve(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	Update(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	Publish(ctx context.Context, id int) (*models.Offer, error)
	Cancel(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	Reject(ctx context.Context, offer *models.Offer) (*models.Offer, error)
}
