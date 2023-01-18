package http

import uMovie "github.com/fajarabdillahfn/go-movie-api/internal/usecase"

type Delivery struct {
	movieUC uMovie.UseCase
}

func NewMovieHTTP(movieUC uMovie.UseCase) *Delivery {
	return &Delivery{
		movieUC: movieUC,
	}
}
