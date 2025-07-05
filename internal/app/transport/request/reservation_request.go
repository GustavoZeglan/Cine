package request

type CreateReservationRequest struct {
	SessionID uint `json:"session_id" validate:"required,gt=0"`
	SeatID    uint `json:"seat_id" validate:"required,gt=0"`
}
