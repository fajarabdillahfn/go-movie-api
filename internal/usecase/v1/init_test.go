package v1

import (
	rMovie "github.com/fajarabdillahfn/go-movie-api/internal/repository/postgre"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestNewMovieUsecase(t *testing.T) {
	usecaseMock := NewMovieUsecase(rMovie.NewMovieRepo(&gorm.DB{}))
	assert.NotNil(t, usecaseMock)

}
