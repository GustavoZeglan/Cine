package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type SeatRepository interface {
	Create(ctx context.Context, seat *entities.Seat) error
	GetByID(ctx context.Context, seatID uint) (*entities.Seat, error)
	Update(ctx context.Context, seat *entities.Seat, seatID uint) error
}
