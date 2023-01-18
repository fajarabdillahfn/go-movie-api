package repository

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
)

type Repository interface {
	Create(ctx context.Context, newMovie *model.Movie) (model.Movie, error)
}
