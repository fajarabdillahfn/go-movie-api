package v1

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"github.com/fajarabdillahfn/go-movie-api/internal/repository"
	"testing"
)

func Test_useCase_UpdateMovie(t *testing.T) {
	type fields struct {
		MovieRepo repository.Repository
	}
	type args struct {
		ctx         context.Context
		id          uint
		updateMovie *model.InputMovie
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				MovieRepo: tt.fields.MovieRepo,
			}
			if err := u.UpdateMovie(tt.args.ctx, tt.args.id, tt.args.updateMovie); (err != nil) != tt.wantErr {
				t.Errorf("UpdateMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
