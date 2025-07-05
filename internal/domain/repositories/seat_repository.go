package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type SeatRepository interface {
	Create(ctx context.Context, seat *entities.Movie) error
	Update(ctx context.Context, seat *entities.Movie, seatID uint) error
}
