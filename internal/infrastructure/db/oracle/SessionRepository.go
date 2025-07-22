package oracle

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
	"gorm.io/gorm"
)

type SessionRepositoryImpl struct {
	db *gorm.DB
}

func NewSessionRepository() repositories.SessionRepository {
	return &SessionRepositoryImpl{db: GetConnection()}
}

func (s *SessionRepositoryImpl) Create(ctx context.Context, session *entities.Session) error {
	return s.db.WithContext(ctx).Create(session).Error
}

func (s *SessionRepositoryImpl) GetByID(ctx context.Context, sessionID uint) (*entities.Session, error) {
	var session entities.Session
	err := s.db.WithContext(ctx).First(&session, sessionID).Error
	return &session, err
}

func (s *SessionRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Session, error) {
	var sessions []*entities.Session
	err := s.db.WithContext(ctx).Find(&sessions).Error
	return sessions, err
}

func (s *SessionRepositoryImpl) Update(ctx context.Context, session *entities.Session, sessionID uint) error {
	session.ID = sessionID
	return s.db.WithContext(ctx).Save(session).Error
}
