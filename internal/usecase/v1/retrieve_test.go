package v1

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"github.com/fajarabdillahfn/go-movie-api/internal/repository"
	"reflect"
	"testing"
)

func Test_useCase_GetAllMovies(t *testing.T) {
	type fields struct {
		MovieRepo repository.Repository
	}
	type args struct {
		ctx context.Context
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Movie
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				MovieRepo: tt.fields.MovieRepo,
			}
			got, err := u.GetAllMovies(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMovies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllMovies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_GetMovieById(t *testing.T) {
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
		want    *model.Movie
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				MovieRepo: tt.fields.MovieRepo,
			}
			got, err := u.GetMovieById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMovieById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMovieById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
