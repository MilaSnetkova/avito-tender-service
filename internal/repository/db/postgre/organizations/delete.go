package organizations

import (
	"context"
	"tender-service/internal/logger"
)

func (r *organizationsRepo) Delete(ctx context.Context, id int) error {
	logger.Debug("DeleteUser", "id", id)
	return nil
}
