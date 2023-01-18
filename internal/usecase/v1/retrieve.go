package v1

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"time"
)

func (u *useCase) GetAllMovies(ctx context.Context) ([]*model.Movie, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	return u.MovieRepo.GetAll(ctx)
}
