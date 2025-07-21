package oracle

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
	"gorm.io/gorm"
)

type RoomRepositoryImpl struct {
	db *gorm.DB
}

func NewRoomRepository() repositories.RoomRepository {
	return &RoomRepositoryImpl{db: GetConnection()}
}

func (r *RoomRepositoryImpl) Create(ctx context.Context, room *entities.Room) error {
	return r.db.WithContext(ctx).Create(room).Error
}
func (r *RoomRepositoryImpl) GetByID(ctx context.Context, roomID uint) (*entities.Room, error) {
	var room entities.Room
	err := r.db.WithContext(ctx).First(&room, roomID).Error
	return &room, err
}

func (r *RoomRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Room, error) {
	var rooms []*entities.Room
	err := r.db.WithContext(ctx).Find(&rooms).Error
	return rooms, err
}

func (r *RoomRepositoryImpl) Update(ctx context.Context, room *entities.Room, roomID uint) error {
	room.ID = roomID
	return r.db.WithContext(ctx).Save(room).Error
}

func (r *RoomRepositoryImpl) Delete(ctx context.Context, roomID uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Room{}, roomID).Error
}
