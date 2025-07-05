package request

type CreateMovieRequest struct {
	Title string `json:"title" validate:"required,min=1"`
}

type UpdateMovieRequest struct {
	Title string `json:"title" validate:"required,min=1"`
}
