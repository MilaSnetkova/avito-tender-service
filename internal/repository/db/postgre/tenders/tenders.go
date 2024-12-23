package tenders

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"tender-service/internal/repository/repo"
)

type tenderRepo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) repo.Tenders {
	return &tenderRepo{pool: pool}
}
