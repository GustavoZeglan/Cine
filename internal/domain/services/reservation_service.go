package services

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/domain/repositories"
)

type ReservationService interface {
	CreateReservation(ctx context.Context, reservation *entities.Reservation) error
	GetReservationByID(ctx context.Context, reservationID uint) (*entities.Reservation, error)
	GetAllReservations(ctx context.Context) ([]*entities.Reservation, error)
	UpdateReservation(ctx context.Context, reservation *entities.Reservation, reservationID uint) error
}

type ReservationServiceImpl struct {
	ReservationRepo repositories.ReservationRepository
}

func NewReservationService(reservationRepo repositories.ReservationRepository) ReservationService {
	return &ReservationServiceImpl{ReservationRepo: reservationRepo}
}

func (r *ReservationServiceImpl) CreateReservation(ctx context.Context, reservation *entities.Reservation) error {
	return r.ReservationRepo.Create(ctx, reservation)
}

func (r *ReservationServiceImpl) GetAllReservations(ctx context.Context) ([]*entities.Reservation, error) {
	return r.ReservationRepo.GetAll(ctx)
}

func (r *ReservationServiceImpl) GetReservationByID(ctx context.Context, reservationID uint) (*entities.Reservation, error) {
	return r.ReservationRepo.GetByID(ctx, reservationID)
}

func (r *ReservationServiceImpl) UpdateReservation(ctx context.Context, reservation *entities.Reservation, reservationID uint) error {
	return r.ReservationRepo.Update(ctx, reservation, reservationID)
}
