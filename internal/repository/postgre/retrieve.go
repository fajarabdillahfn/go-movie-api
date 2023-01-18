package postgre

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
)

func (r *repository) GetAll(ctx context.Context) ([]*model.Movie, error) {
	var movies []*model.Movie

	res := r.conn.WithContext(ctx).
		Find(&movies)

	if res.Error != nil {
		return nil, res.Error
	}

	return movies, nil
}
