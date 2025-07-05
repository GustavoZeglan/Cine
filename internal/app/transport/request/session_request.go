package request

type CreateSessionRequest struct {
	MovieID uint   `json:"movie_id"`
	RoomID  uint   `json:"room_id"`
	StartAt string `json:"start_at"`
}

type UpdateSessionRequest struct {
	MovieID uint   `json:"movie_id"`
	RoomID  uint   `json:"room_id"`
	StartAt string `json:"start_at"`
}
