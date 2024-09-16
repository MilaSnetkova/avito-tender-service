package organizations

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"tender-service/internal/repository/repo"
)

type organizationsRepo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) repo.Organizations {
	return &organizationsRepo{pool: pool}
}
