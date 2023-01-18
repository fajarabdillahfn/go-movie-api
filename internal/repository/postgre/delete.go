package postgre

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
)

func (r *repository) Delete(ctx context.Context, id uint) error {
	res := r.conn.WithContext(ctx).
		Delete(&model.Movie{}, id)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
