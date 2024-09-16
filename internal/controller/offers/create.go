package offers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"tender-service/internal/models"
)

func (c *offersCtrl) CreateOffer(ctx context.Context, body io.ReadCloser) (*models.Offer, error) {
	var offer models.Offer

	if err := json.NewDecoder(body).Decode(&offer); err != nil {
		return nil, errors.New("error decoding request body")
	}

	if offer.TenderID == 0 {
		return nil, fmt.Errorf("tenderID is required")
	}

	if offer.UserID == 0 {
		return nil, fmt.Errorf("userID is required")
	}

	if offer.Description == "" {
		return nil, fmt.Errorf("description is required")
	}

	newOffer, err := c.repo.Offers.Create(ctx, &offer)
	if err != nil {
		return nil, err
	}

	return newOffer, nil
}
