package services

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
)

type RoomService interface {
	CreateRoom(ctx context.Context, room *entities.Room) error
	GetRoomByID(ctx context.Context, roomID uint) (*entities.Room, error)
	GetAllRooms(ctx context.Context) ([]*entities.Room, error)
	UpdateRoom(ctx context.Context, room *entities.Room, roomID uint) error
	DeleteRoom(ctx context.Context, roomID uint) error
}

type RoomServiceImpl struct {
	RoomRepo repositories.RoomRepository
}

func NewRoomService(roomRepo repositories.RoomRepository) RoomService {
	return &RoomServiceImpl{RoomRepo: roomRepo}
}

func (r *RoomServiceImpl) CreateRoom(ctx context.Context, room *entities.Room) error {
	return r.RoomRepo.Create(ctx, room)
}

func (r *RoomServiceImpl) GetAllRooms(ctx context.Context) ([]*entities.Room, error) {
	return r.RoomRepo.GetAll(ctx)
}

func (r *RoomServiceImpl) GetRoomByID(ctx context.Context, roomID uint) (*entities.Room, error) {
	return r.RoomRepo.GetByID(ctx, roomID)
}

func (r *RoomServiceImpl) UpdateRoom(ctx context.Context, room *entities.Room, roomID uint) error {
	return r.RoomRepo.Update(ctx, room, roomID)
}

func (r *RoomServiceImpl) DeleteRoom(ctx context.Context, roomID uint) error {
	return r.RoomRepo.Delete(ctx, roomID)
}
