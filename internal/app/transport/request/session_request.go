package request

type CreateSessionRequest struct {
	MovieID uint   `json:"movie_id" validate:"required,gt=0"`
	RoomID  uint   `json:"room_id" validate:"required,gt=0"`
	StartAt string `json:"start_at" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

type UpdateSessionRequest struct {
	MovieID uint   `json:"movie_id" validate:"required,gt=0"`
	RoomID  uint   `json:"room_id" validate:"required,gt=0"`
	StartAt string `json:"start_at" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}
