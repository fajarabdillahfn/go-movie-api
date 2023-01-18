package postgre

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
)

func (r *repository) Create(ctx context.Context, newMovie *model.Movie) (model.Movie, error) {
	res := r.conn.
		WithContext(ctx).
		Create(&newMovie)

	if res.Error != nil {
		return model.Movie{}, res.Error
	}

	return *newMovie, nil
}
