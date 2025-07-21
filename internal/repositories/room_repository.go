package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type RoomRepository interface {
	Create(ctx context.Context, room *entities.Room) error
	GetByID(ctx context.Context, roomID uint) (*entities.Room, error)
	GetAll(ctx context.Context) ([]*entities.Room, error)
	Update(ctx context.Context, room *entities.Room, roomID uint) error
	Delete(ctx context.Context, roomID uint) error
}
