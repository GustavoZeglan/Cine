package response

type SessionResponse struct {
	ID      uint   `json:"id"`
	MovieID uint   `json:"movie_id"`
	RoomID  uint   `json:"room_id"`
	StartAt string `json:"start_at"`
}
