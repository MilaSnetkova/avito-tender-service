package users

import (
	"context"
	_ "embed"
	"fmt"

	"tender-service/internal/models"

	"github.com/jackc/pgx/v5"
)

//go:embed queries/getuser.sql

var get string

func (r *userRepo) Get(ctx context.Context, id int) (*models.User, error) {
	row := r.pool.QueryRow(ctx, get, id)

	var user models.User

	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	return &user, nil
}
