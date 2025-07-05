package request

type CreateReservationRequest struct {
	SessionID uint `json:"session_id"`
	SeatID    uint `json:"seat_id"`
}
