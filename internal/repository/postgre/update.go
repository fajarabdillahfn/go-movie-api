package postgre

import (
	"context"
	"errors"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
)

func (r *repository) Update(ctx context.Context, id uint, updateDetails map[string]interface{}) error {
	res := r.conn.WithContext(ctx).
		Model(&model.Movie{}).
		Where("id = ?", id).
		Updates(&updateDetails)

	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected != 1 {
		return errors.New("update failed")
	}

	return nil
}
