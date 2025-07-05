package postgres

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/domain/repositories"
	"gorm.io/gorm"
)

type ReservationRepo struct {
	DB *gorm.DB
}

func NewReservationRepository(DB *gorm.DB) repositories.ReservationRepository {
	return &ReservationRepo{DB}
}

func (r *ReservationRepo) Create(ctx context.Context, reservation *entities.Reservation) error {
	return r.DB.WithContext(ctx).Create(reservation).Error
}

func (r *ReservationRepo) Update(ctx context.Context, reservation *entities.Reservation, reservationID uint) error {
	reservation.ID = reservationID
	return r.DB.WithContext(ctx).Save(reservation).Error
}
