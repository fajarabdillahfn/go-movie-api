package usecase

//go:generate moq -out usecase_test.go . UseCase

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
)

type UseCase interface {
	CreateMovie(ctx context.Context, newMovie *model.Movie) (model.Movie, error)

	GetAllMovies(ctx context.Context) ([]*model.Movie, error)
	GetMovieById(ctx context.Context, id uint) (*model.Movie, error)

	UpdateMovie(ctx context.Context, id uint, updateMovie *model.InputMovie) error

	DeleteMovieById(ctx context.Context, id uint) error
}
