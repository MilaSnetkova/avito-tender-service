package offers

import (
	"context"
	"io"

	"tender-service/internal/models"
	"tender-service/internal/repository/repo"
)

type Offers interface {
	ApproveOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	CreateOffer(ctx context.Context, body io.ReadCloser) (*models.Offer, error)
	CancelOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	UpdateOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	PublishOffer(ctx context.Context, id int) (*models.Offer, error)
	RejectOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error)
}

type offersCtrl struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) Offers {
	return &offersCtrl{repo: repo}
}
