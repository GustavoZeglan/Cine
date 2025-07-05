package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type RoomRepository interface {
	Create(ctx context.Context, room *entities.Room) error
	Update(ctx context.Context, room *entities.Room, roomID uint) error
}
