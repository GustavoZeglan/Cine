package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type SessionRepository interface {
	Create(ctx context.Context, session *entities.Session) error
	GetByID(ctx context.Context, sessionID uint) (*entities.Session, error)
	GetAll(ctx context.Context) ([]*entities.Session, error)
	Update(ctx context.Context, session *entities.Session, sessionID uint) error
}
