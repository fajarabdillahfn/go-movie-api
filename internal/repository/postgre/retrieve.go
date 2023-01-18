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

func (r *repository) GetById(ctx context.Context, id uint) (*model.Movie, error) {
	var movie *model.Movie

	res := r.conn.WithContext(ctx).
		First(&movie, id)

	if res.Error != nil {
		return nil, res.Error
	}

	return movie, nil
}
