package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type SessionRepository interface {
	Create(ctx context.Context, session *entities.Movie) error
	Update(ctx context.Context, session *entities.Movie, sessionID uint) error
}
