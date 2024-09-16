package users

import (
	"context"
	"tender-service/internal/logger"
)

func (r *userRepo) Delete(ctx context.Context, id int) error {
	logger.Debug("DeleteUser", "id", id)
	return nil
}
