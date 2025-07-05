package repositories

import (
	"context"

	"github.com/GustavoZeglan/Cine/internal/domain/entities"
)

type MovieRepository interface {
	Create(ctx context.Context, movie *entities.Movie) error
	Update(ctx context.Context, movie *entities.Movie, movieID uint) error
}
