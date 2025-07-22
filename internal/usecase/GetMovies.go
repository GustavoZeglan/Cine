package usecase

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
)

type GetMovies struct {
	repo repositories.MovieRepository
}

func NewGetMovies(repo repositories.MovieRepository) *GetMovies {
	return &GetMovies{repo: repo}
}

func (gm *GetMovies) Execute(ctx context.Context) ([]*entities.Movie, error) {
	return gm.repo.GetAll(ctx)
}
