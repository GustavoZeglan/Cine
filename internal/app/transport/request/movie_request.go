package request

type CreateMovieRequest struct {
	Title string `json:"title"`
}

type UpdateMovieRequest struct {
	Title string `json:"title"`
}
