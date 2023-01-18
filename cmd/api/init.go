package api

import (
	"gorm.io/gorm"

	cDB "github.com/fajarabdillahfn/go-movie-api/common/db/postgre"

	httpMovie "github.com/fajarabdillahfn/go-movie-api/internal/delivery/http"
	rMovie "github.com/fajarabdillahfn/go-movie-api/internal/repository"
	uMovie "github.com/fajarabdillahfn/go-movie-api/internal/usecase"

	pgMovieRepo "github.com/fajarabdillahfn/go-movie-api/internal/repository/postgre"
	uMoveV1 "github.com/fajarabdillahfn/go-movie-api/internal/usecase/v1"
)

var (
	pgMovie           *gorm.DB
	movieRepo         rMovie.Repository
	movieUseCase      uMovie.UseCase
	movieHTTPDelivery *httpMovie.Delivery
)

func initialize() {
	pgMovie = cDB.OpenDB()

	movieRepo = pgMovieRepo.NewMovieRepo(pgMovie)
	movieUseCase = uMoveV1.NewMovieUsecase(movieRepo)
	movieHTTPDelivery = httpMovie.NewMovieHTTP(movieUseCase)
}
