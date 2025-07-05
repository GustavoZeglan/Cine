package request

type CreateRoomRequest struct {
	Name     string `json:"name" validate:"required,min=1"`
	Capacity int    `json:"capacity" validate:"required,gt=0"`
}

type UpdateRoomRequest struct {
	Name     string `json:"name" validate:"required,min=1"`
	Capacity int    `json:"capacity" validate:"required,gt=0"`
}
