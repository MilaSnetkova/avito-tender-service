package tenders

import (
	"context"
	"io"

	"tender-service/internal/models"
	"tender-service/internal/repository/repo"
)

type Tenders interface {
	CreateTender(ctx context.Context, body io.ReadCloser) (*models.Tender, error)
	PublishTender(ctx context.Context, id int) (*models.Tender, error)
	CloseTender(ctx context.Context, id int) (*models.Tender, error)
	UpdateTender(ctx context.Context, tender *models.Tender) (*models.Tender, error)
}

type tendersCtrl struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) Tenders {
	return &tendersCtrl{repo: repo}
}
