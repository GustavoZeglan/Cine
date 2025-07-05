package services

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/domain/repositories"
)

type SeatService interface {
	CreateSeat(ctx context.Context, seat *entities.Seat) error
	GetSeatByID(ctx context.Context, seatID uint) (*entities.Seat, error)
	UpdateSeat(ctx context.Context, seat *entities.Seat, seatID uint) error
}

type SeatServiceImpl struct {
	SeatRepo repositories.SeatRepository
}

func NewSeatService(seatRepo repositories.SeatRepository) SeatService {
	return &SeatServiceImpl{SeatRepo: seatRepo}
}

func (s *SeatServiceImpl) CreateSeat(ctx context.Context, seat *entities.Seat) error {
	return s.SeatRepo.Create(ctx, seat)
}

func (s *SeatServiceImpl) GetSeatByID(ctx context.Context, seatID uint) (*entities.Seat, error) {
	return s.SeatRepo.GetByID(ctx, seatID)
}

func (s *SeatServiceImpl) UpdateSeat(ctx context.Context, seat *entities.Seat, seatID uint) error {
	return s.SeatRepo.Update(ctx, seat, seatID)
}
