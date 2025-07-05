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
func (r *RoomRepo) GetByID(ctx context.Context, roomID uint) (*entities.Room, error) {
	var room entities.Room
	err := r.DB.WithContext(ctx).First(&room, roomID).Error
	return &room, err
}

func (r *RoomRepo) GetAll(ctx context.Context) ([]*entities.Room, error) {
	var rooms []*entities.Room
	err := r.DB.WithContext(ctx).Find(&rooms).Error
	return rooms, err
}

func (r *RoomRepo) Update(ctx context.Context, room *entities.Room, roomID uint) error {
	room.ID = roomID
	return r.DB.WithContext(ctx).Save(room).Error
}

func (r *RoomRepo) Delete(ctx context.Context, roomID uint) error {
	return r.DB.WithContext(ctx).Delete(&entities.Room{}, roomID).Error
}
