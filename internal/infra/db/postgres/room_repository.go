package postgres

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/domain/repositories"
	"gorm.io/gorm"
)

type RoomRepo struct {
	DB *gorm.DB
}

func NewRoomRepository(DB *gorm.DB) repositories.RoomRepository {
	return &RoomRepo{DB}
}

func (r *RoomRepo) Create(ctx context.Context, room *entities.Room) error {
	return r.DB.WithContext(ctx).Create(room).Error
}

func (r *RoomRepo) Update(ctx context.Context, room *entities.Room, roomID uint) error {
	room.ID = roomID
	return r.DB.WithContext(ctx).Save(room).Error
}
