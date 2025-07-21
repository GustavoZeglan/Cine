package postgres

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
	"gorm.io/gorm"
)

type SessionRepo struct {
	DB *gorm.DB
}

func NewSessionRepository(DB *gorm.DB) repositories.SessionRepository {
	return &SessionRepo{DB}
}

func (s *SessionRepo) Create(ctx context.Context, session *entities.Session) error {
	return s.DB.WithContext(ctx).Create(session).Error
}

func (s *SessionRepo) GetByID(ctx context.Context, sessionID uint) (*entities.Session, error) {
	var session entities.Session
	err := s.DB.WithContext(ctx).First(&session, sessionID).Error
	return &session, err
}

func (s *SessionRepo) GetAll(ctx context.Context) ([]*entities.Session, error) {
	var sessions []*entities.Session
	err := s.DB.WithContext(ctx).Find(&sessions).Error
	return sessions, err
}

func (s *SessionRepo) Update(ctx context.Context, session *entities.Session, sessionID uint) error {
	session.ID = sessionID
	return s.DB.WithContext(ctx).Save(session).Error
}
