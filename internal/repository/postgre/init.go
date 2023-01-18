package postgre

import (
	rMovie "github.com/fajarabdillahfn/go-movie-api/internal/repository"
	"gorm.io/gorm"
	"log"
)

type repository struct {
	conn *gorm.DB
}

func NewMovieRepo(db *gorm.DB) rMovie.Repository {
	if db == nil {
		log.Panic("missing database connection")
	}

	repo := &repository{
		conn: db,
	}

	return repo
}
