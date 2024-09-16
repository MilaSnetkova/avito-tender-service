package offers

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"tender-service/internal/repository/repo"
)

type offerRepo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) repo.Offers {
	return &offerRepo{pool: pool}
}
