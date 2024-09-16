package controller

import (
	"context"
	"io"
	"tender-service/internal/controller/offers"
	"tender-service/internal/controller/tenders"
	"tender-service/internal/models"
	"tender-service/internal/repository/repo"
)

type ctrl struct {
	tenders tenders.Tenders
	offers  offers.Offers
}

type Controller interface {
	CreateTender(ctx context.Context, body io.ReadCloser) (*models.Tender, error)
	PublishTender(ctx context.Context, tenderID int) (*models.Tender, error)
	CloseTender(ctx context.Context, tednderID int) (*models.Tender, error)
	UpdateTender(ctx context.Context, tender *models.Tender) (*models.Tender, error)
	ApproveOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	CreateOffer(ctx context.Context, body io.ReadCloser) (*models.Offer, error)
	CancelOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	UpdateOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error)
	PublishOffer(ctx context.Context, offerID int) (*models.Offer, error)
	RejectOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error)
}

func New(repo *repo.Repo) Controller {
	tenderCtrl := tenders.New(repo)
	offerCtrl := offers.New(repo)

	return &ctrl{
		tenders: tenderCtrl,
		offers:  offerCtrl,
	}
}

func (c *ctrl) CreateTender(ctx context.Context, body io.ReadCloser) (*models.Tender, error) {
	return c.tenders.CreateTender(ctx, body)
}

func (c *ctrl) PublishTender(ctx context.Context, tenderID int) (*models.Tender, error) {
	return c.tenders.PublishTender(ctx, tenderID)
}

func (c *ctrl) CloseTender(ctx context.Context, tenderID int) (*models.Tender, error) {
	return c.tenders.CloseTender(ctx, tenderID)
}

func (c *ctrl) UpdateTender(ctx context.Context, tender *models.Tender) (*models.Tender, error) {
	return c.tenders.UpdateTender(ctx, tender)
}

func (c *ctrl) CreateOffer(ctx context.Context, body io.ReadCloser) (*models.Offer, error) {
	return c.offers.CreateOffer(ctx, body)
}

func (c *ctrl) ApproveOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	return c.offers.ApproveOffer(ctx, offer)
}

func (c *ctrl) RejectOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	return c.offers.RejectOffer(ctx, offer)
}

func (c *ctrl) PublishOffer(ctx context.Context, offerID int) (*models.Offer, error) {
	return c.offers.PublishOffer(ctx, offerID)
}

func (c *ctrl) CancelOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	return c.offers.CancelOffer(ctx, offer)
}

func (c *ctrl) UpdateOffer(ctx context.Context, offer *models.Offer) (*models.Offer, error) {
	return c.offers.UpdateOffer(ctx, offer)
}
