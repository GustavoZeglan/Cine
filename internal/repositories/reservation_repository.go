package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type ReservationRepository interface {
	Create(ctx context.Context, reservation *entities.Reservation) error
	GetByID(ctx context.Context, reservationID uint) (*entities.Reservation, error)
	GetAll(ctx context.Context) ([]*entities.Reservation, error)
}
