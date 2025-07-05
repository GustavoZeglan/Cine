package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type ReservationRepository interface {
	Create(ctx context.Context, reservation *entities.Movie) error
	Update(ctx context.Context, reservation *entities.Movie, reservationID uint) error
}
