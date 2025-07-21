package postgres

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
	"gorm.io/gorm"
)

type ReservationRepositoryImpl struct {
	DB *gorm.DB
}

func NewReservationRepository(DB *gorm.DB) repositories.ReservationRepository {
	return &ReservationRepositoryImpl{DB}
}

func (r *ReservationRepositoryImpl) Create(ctx context.Context, reservation *entities.Reservation) error {
	return r.DB.WithContext(ctx).Create(reservation).Error
}

func (r *ReservationRepositoryImpl) GetByID(ctx context.Context, reservationID uint) (*entities.Reservation, error) {
	panic("unimplemented")
}

func (r *ReservationRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Reservation, error) {
	panic("unimplemented")
}

func (r *ReservationRepositoryImpl) Update(ctx context.Context, reservation *entities.Reservation, reservationID uint) error {
	reservation.ID = reservationID
	return r.DB.WithContext(ctx).Save(reservation).Error
}
