package postgre

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestNewMovieRepoPanic(t *testing.T) {
	assert.Panics(t, func() {
		NewMovieRepo(nil)
	})
}

func TestNewMovieRepoSuccess(t *testing.T) {
	assert.NotNil(t, NewMovieRepo(&gorm.DB{}))
}
