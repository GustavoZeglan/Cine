package response

type AvailableSeat struct {
	SeatID uint   `json:"seat_id"`
	Row    string `json:"row"`
	Number int    `json:"number"`
}

type ListAvailableSeatsResponse struct {
	SessionID uint            `json:"session_id"`
	Seats     []AvailableSeat `json:"seats"`
}

type SessionItem struct {
	SessionID uint   `json:"session_id"`
	RoomID    uint   `json:"room_id"`
	StartAt   string `json:"start_at"`
}

type ListSessionsByMovieResponse struct {
	MovieID  uint          `json:"movie_id"`
	Sessions []SessionItem `json:"sessions"`
}
