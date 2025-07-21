package services

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
)

type MovieService interface {
	CreateMovie(ctx context.Context, movie *entities.Movie) error
	GetMovieByID(ctx context.Context, movieID uint) (*entities.Movie, error)
	GetAllMovies(ctx context.Context) ([]*entities.Movie, error)
	UpdateMovie(ctx context.Context, movie *entities.Movie, movieID uint) error
	DeleteMovie(ctx context.Context, movieID uint) error
}

type MovieServiceImpl struct {
	MovieRepo repositories.MovieRepository
}

func NewMovieService(movieRepo repositories.MovieRepository) MovieService {
	return &MovieServiceImpl{MovieRepo: movieRepo}
}

func (m *MovieServiceImpl) CreateMovie(ctx context.Context, movie *entities.Movie) error {
	return m.MovieRepo.Create(ctx, movie)
}

func (m *MovieServiceImpl) DeleteMovie(ctx context.Context, movieID uint) error {
	return m.MovieRepo.Delete(ctx, movieID)
}

func (m *MovieServiceImpl) GetAllMovies(ctx context.Context) ([]*entities.Movie, error) {
	return m.MovieRepo.GetAll(ctx)
}

func (m *MovieServiceImpl) GetMovieByID(ctx context.Context, movieID uint) (*entities.Movie, error) {
	return m.MovieRepo.GetByID(ctx, movieID)
}

func (m *MovieServiceImpl) UpdateMovie(ctx context.Context, movie *entities.Movie, movieID uint) error {
	return m.MovieRepo.Update(ctx, movie, movieID)
}
