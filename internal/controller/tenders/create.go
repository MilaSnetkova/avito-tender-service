package tenders

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"tender-service/internal/models"
)

func (c *tendersCtrl) CreateTender(ctx context.Context, body io.ReadCloser) (*models.Tender, error) {
	var t models.Tender

	if err := json.NewDecoder(body).Decode(&t); err != nil {
		return nil, errors.New("error decoding request body")
	}

	if t.Title == "" {
		return nil, errors.New("missing tender title")
	}

	tender, err := c.repo.Tenders.Create(ctx, &t)
	if err != nil {
		return nil, err
	}

	return tender, nil
}
