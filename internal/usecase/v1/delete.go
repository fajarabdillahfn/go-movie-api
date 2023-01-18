package v1

import (
	"context"
	"time"
)

func (u *useCase) DeleteMovieById(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	return u.MovieRepo.Delete(ctx, id)
}
