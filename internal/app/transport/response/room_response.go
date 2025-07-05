package response

type RoomResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}
