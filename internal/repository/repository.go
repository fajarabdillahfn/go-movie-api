package repository

//go:generate moq -out repository_test.go . Repository

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
)

type Repository interface {
	Create(ctx context.Context, newMovie *model.Movie) (model.Movie, error)

	GetAll(ctx context.Context) ([]*model.Movie, error)
	GetById(ctx context.Context, id uint) (*model.Movie, error)

	Update(ctx context.Context, id uint, updateDetails map[string]interface{}) error

	Delete(ctx context.Context, id uint) error
}
