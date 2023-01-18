package v1

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"time"
)

func (u *useCase) CreateMovie(ctx context.Context, newMovie *model.Movie) (model.Movie, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	return u.MovieRepo.Create(ctx, newMovie)
}
