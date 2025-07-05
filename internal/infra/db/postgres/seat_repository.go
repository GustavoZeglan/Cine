package postgres

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/domain/repositories"
	"gorm.io/gorm"
)

type SeatRepo struct {
	DB *gorm.DB
}

func NewSeatRepository(DB *gorm.DB) repositories.SeatRepository {
	return &SeatRepo{DB}
}

func (s *SeatRepo) Create(ctx context.Context, seat *entities.Movie) error {
	return s.DB.WithContext(ctx).Create(seat).Error
}

func (s *SeatRepo) Update(ctx context.Context, seat *entities.Movie, seatID uint) error {
	seat.ID = seatID
	return s.DB.WithContext(ctx).Save(seat).Error
}
