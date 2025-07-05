package response

type ReservationResponse struct {
	ID        uint `json:"id"`
	SessionID uint `json:"session_id"`
	SeatID    uint `json:"seat_id"`
}
