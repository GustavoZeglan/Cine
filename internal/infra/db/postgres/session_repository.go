package postgres

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/domain/repositories"
	"gorm.io/gorm"
)

type SessionRepo struct {
	DB *gorm.DB
}

func NewSessionRepository(DB *gorm.DB) repositories.SessionRepository {
	return &SessionRepo{DB}
}

func (s *SessionRepo) Create(ctx context.Context, session *entities.Movie) error {
	return s.DB.WithContext(ctx).Create(session).Error
}

func (s *SessionRepo) Update(ctx context.Context, session *entities.Movie, sessionID uint) error {
	session.ID = sessionID
	return s.DB.WithContext(ctx).Save(session).Error
}
