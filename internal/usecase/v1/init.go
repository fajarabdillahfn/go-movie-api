package v1

import (
	rMovie "github.com/fajarabdillahfn/go-movie-api/internal/repository"
	uMovie "github.com/fajarabdillahfn/go-movie-api/internal/usecase"
)

type useCase struct {
	MovieRepo rMovie.Repository
}

func NewMovieUsecase(movieRepo rMovie.Repository) uMovie.UseCase {
	return &useCase{
		MovieRepo: movieRepo,
	}
}
