package response

type SeatResponse struct {
	ID     uint `json:"id"`
	RoomID uint `json:"room_id"`
	Row    int  `json:"row"`
	Column int  `json:"column"`
}
