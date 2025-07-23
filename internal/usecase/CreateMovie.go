package usecase

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/dto"
	"github.com/GustavoZeglan/Cine/internal/repositories"
)

type CreateMovie struct {
	repo repositories.MovieRepository
}

func NewCreateMovie(repo repositories.MovieRepository) *CreateMovie {
	return &CreateMovie{repo: repo}
}

func (gm *CreateMovie) Execute(ctx context.Context, input dto.CreateMovieInput) (*Output, error) {
	movie := &entities.Movie{Title: input.Title}
	if err := gm.repo.Create(ctx, movie); err != nil {
		return nil, nil
	}
	return &Output{ID: movie.ID, Title: movie.Title}, nil
}

type Output struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
