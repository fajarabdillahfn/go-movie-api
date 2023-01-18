package v1

import (
	"context"
	"encoding/json"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"time"
)

func (u *useCase) UpdateMovie(ctx context.Context, id uint, updateMovie *model.InputMovie) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	var updateDetails map[string]interface{}
	data, _ := json.Marshal(updateMovie)
	err := json.Unmarshal(data, &updateDetails)
	if err != nil {
		return err
	}

	return u.MovieRepo.Update(ctx, id, updateDetails)
}
