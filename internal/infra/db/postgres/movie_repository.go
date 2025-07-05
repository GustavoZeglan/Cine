package postgres

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/domain/repositories"
	"gorm.io/gorm"
)

type MovieRepo struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) repositories.MovieRepository {
	return &MovieRepo{db}
}

func (m *MovieRepo) Create(ctx context.Context, movie *entities.Movie) error {
	return m.db.WithContext(ctx).Create(movie).Error
}

func (m *MovieRepo) Update(ctx context.Context, movie *entities.Movie, movieID uint) error {
	movie.ID = movieID
	return m.db.WithContext(ctx).Save(&movie).Error
}
