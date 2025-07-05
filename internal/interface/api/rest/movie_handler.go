package rest

import "github.com/GustavoZeglan/Cine/internal/domain/services"

type MovieHandler struct {
	MovieService services.MovieService
}

func NewMovieHandler(movieService services.MovieService) *MovieHandler {
	return &MovieHandler{
		MovieService: movieService,
	}
}
