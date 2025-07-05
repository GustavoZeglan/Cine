package request

type CreateRoomRequest struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type UpdateRoomRequest struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}
