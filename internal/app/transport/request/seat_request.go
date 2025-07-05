package request

type CreateSeatRequest struct {
	RoomID uint `json:"room_id" validate:"required,gt=0"`
	Row    int  `json:"row" validate:"required,gt=0"`
	Column int  `json:"column" validate:"required,gt=0"`
}

type UpdateSeatRequest struct {
	Row    int `json:"row" validate:"required,gt=0"`
	Column int `json:"column" validate:"required,gt=0"`
}
