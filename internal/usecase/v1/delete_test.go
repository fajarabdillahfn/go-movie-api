package v1

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/repository"
	"testing"
)

func Test_useCase_DeleteMovieById(t *testing.T) {
	type fields struct {
		MovieRepo repository.Repository
	}
	type args struct {
		ctx context.Context
		id  uint
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
			if err := u.DeleteMovieById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMovieById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
