package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type ReservationRepository interface {
	Create(ctx context.Context, reservation *entities.Reservation) error
	Update(ctx context.Context, reservation *entities.Reservation, reservationID uint) error
}
