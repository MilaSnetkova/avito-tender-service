package users

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"tender-service/internal/repository/repo"
)

type userRepo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) repo.Users {
	return &userRepo{pool: pool}
}
