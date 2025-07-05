package request

type CreateSeatRequest struct {
	RoomID uint `json:"room_id"`
	Row    int  `json:"row"`
	Column int  `json:"column"`
}

type UpdateSeatRequest struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}
