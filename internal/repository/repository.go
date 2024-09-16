package repository

import (
	"context"
	"errors"
	"fmt"

	"tender-service/internal/config"
	db "tender-service/internal/repository/db/postgre"
	"tender-service/internal/repository/repo"
)

func New(ctx context.Context, cfg config.Config, repoType repo.Type) (*repo.Repo, error) {
	switch repoType {
	case repo.PSQL:
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			cfg.PostgresUsername,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase)
		return db.New(ctx, dsn)
	default:
		return nil, errors.New("invalid repository type")
	}
}
