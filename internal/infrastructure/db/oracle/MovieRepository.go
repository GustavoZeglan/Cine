package oracle

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/repositories"
	"gorm.io/gorm"
)

type MovieRepositoryImpl struct {
	db *gorm.DB
}

func NewMovieRepository() repositories.MovieRepository {
	return &MovieRepositoryImpl{db: GetConnection()}
}

func (m *MovieRepositoryImpl) Create(ctx context.Context, movie *entities.Movie) error {
	return m.db.WithContext(ctx).Create(movie).Error
}

func (m *MovieRepositoryImpl) GetByID(ctx context.Context, movieID uint) (*entities.Movie, error) {
	var movie entities.Movie
	err := m.db.WithContext(ctx).First(&movie, movieID).Error
	return &movie, err
}

func (m *MovieRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Movie, error) {
	var movies []*entities.Movie
	err := m.db.WithContext(ctx).Find(&movies).Error
	return movies, err
}

func (m *MovieRepositoryImpl) Update(ctx context.Context, movie *entities.Movie, movieID uint) error {
	movie.ID = movieID
	return m.db.WithContext(ctx).Save(&movie).Error
}

func (m *MovieRepositoryImpl) Delete(ctx context.Context, movieID uint) error {
	return m.db.WithContext(ctx).Delete(&entities.Movie{}, movieID).Error
}
