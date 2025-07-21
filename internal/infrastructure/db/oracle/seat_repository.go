package oracle

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/domain/repositories"
	"gorm.io/gorm"
)

type SeatRepositoryImpl struct {
	db *gorm.DB
}

func NewSeatRepository() repositories.SeatRepository {
	return &SeatRepositoryImpl{db: GetConnection()}
}

func (s *SeatRepositoryImpl) Create(ctx context.Context, seat *entities.Seat) error {
	return s.db.WithContext(ctx).Create(seat).Error
}

func (s *SeatRepositoryImpl) GetByID(ctx context.Context, seatID uint) (*entities.Seat, error) {
	var seat entities.Seat
	err := s.db.WithContext(ctx).First(&seat, seatID).Error
	return &seat, err
}

func (s *SeatRepositoryImpl) Update(ctx context.Context, seat *entities.Seat, seatID uint) error {
	seat.ID = seatID
	return s.db.WithContext(ctx).Save(seat).Error
}
