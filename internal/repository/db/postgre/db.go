package db

import (
	"context"

	"tender-service/internal/logger"
	"tender-service/internal/repository/db/postgre/offers"
	"tender-service/internal/repository/db/postgre/organizations"
	"tender-service/internal/repository/db/postgre/tenders"
	"tender-service/internal/repository/db/postgre/users"
	"tender-service/internal/repository/repo"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, dsn string) (*repo.Repo, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to postgres")

	go func(ctx context.Context, pool *pgxpool.Pool) {
		<-ctx.Done()
		logger.Info("Close postgres connection")
		pool.Close()
	}(ctx, pool)

	return &repo.Repo{
		Users:         users.New(pool),
		Organizations: organizations.New(pool),
		Tenders:       tenders.New(pool),
		Offers:        offers.New(pool),
	}, nil
}
