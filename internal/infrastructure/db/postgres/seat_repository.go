package postgres

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
	"gorm.io/gorm"
)

type SeatRepositoryImpl struct {
	DB *gorm.DB
}

func NewSeatRepository(DB *gorm.DB) repositories.SeatRepository {
	return &SeatRepositoryImpl{DB}
}

func (s *SeatRepositoryImpl) Create(ctx context.Context, seat *entities.Seat) error {
	return s.DB.WithContext(ctx).Create(seat).Error
}

func (s *SeatRepositoryImpl) GetByID(ctx context.Context, seatID uint) (*entities.Seat, error) {
	var seat entities.Seat
	err := s.DB.WithContext(ctx).First(&seat, seatID).Error
	return &seat, err
}

func (s *SeatRepositoryImpl) Update(ctx context.Context, seat *entities.Seat, seatID uint) error {
	seat.ID = seatID
	return s.DB.WithContext(ctx).Save(seat).Error
}
