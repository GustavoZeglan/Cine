package oracle

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
	"gorm.io/gorm"
)

type ReservationRepositoryImpl struct {
	db *gorm.DB
}

func NewReservationRepository() repositories.ReservationRepository {
	return &ReservationRepositoryImpl{db: GetConnection()}
}

func (r *ReservationRepositoryImpl) Create(ctx context.Context, reservation *entities.Reservation) error {
	return r.db.WithContext(ctx).Create(reservation).Error
}

func (r *ReservationRepositoryImpl) GetByID(ctx context.Context, reservationID uint) (*entities.Reservation, error) {
	panic("unimplemented")
}

func (r *ReservationRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Reservation, error) {
	panic("unimplemented")
}

func (r *ReservationRepositoryImpl) Update(ctx context.Context, reservation *entities.Reservation, reservationID uint) error {
	reservation.ID = reservationID
	return r.db.WithContext(ctx).Save(reservation).Error
}
