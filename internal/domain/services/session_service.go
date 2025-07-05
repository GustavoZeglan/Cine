package services

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/domain/repositories"
)

type SessionService interface {
	CreateSession(ctx context.Context, session *entities.Session) error
	GetSessionByID(ctx context.Context, sessionID uint) (*entities.Session, error)
	GetAllSessions(ctx context.Context) ([]*entities.Session, error)
	UpdateSession(ctx context.Context, session *entities.Session, sessionID uint) error
}

type SessionServiceImpl struct {
	SessionRepo repositories.SessionRepository
}

func NewSessionService(sessionRepo repositories.SessionRepository) SessionService {
	return &SessionServiceImpl{SessionRepo: sessionRepo}
}

func (s *SessionServiceImpl) CreateSession(ctx context.Context, session *entities.Session) error {
	return s.SessionRepo.Create(ctx, session)
}

func (s *SessionServiceImpl) GetAllSessions(ctx context.Context) ([]*entities.Session, error) {
	return s.SessionRepo.GetAll(ctx)
}

func (s *SessionServiceImpl) GetSessionByID(ctx context.Context, sessionID uint) (*entities.Session, error) {
	return s.SessionRepo.GetByID(ctx, sessionID)
}

func (s *SessionServiceImpl) UpdateSession(ctx context.Context, session *entities.Session, sessionID uint) error {
	return s.SessionRepo.Update(ctx, session, sessionID)
}
