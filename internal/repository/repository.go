package repository

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
)

type Repository interface {
	Create(ctx context.Context, newMovie *model.Movie) (model.Movie, error)

	GetAll(ctx context.Context) ([]*model.Movie, error)
	GetById(ctx context.Context, id uint) (*model.Movie, error)
}
