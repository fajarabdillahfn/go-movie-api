package v1

import (
	"context"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"github.com/fajarabdillahfn/go-movie-api/internal/repository"
	"reflect"
	"testing"
)

func Test_useCase_CreateMovie(t *testing.T) {
	type fields struct {
		MovieRepo repository.Repository
	}
	type args struct {
		ctx      context.Context
		newMovie *model.Movie
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    model.Movie
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				MovieRepo: tt.fields.MovieRepo,
			}
			got, err := u.CreateMovie(tt.args.ctx, tt.args.newMovie)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMovie() got = %v, want %v", got, tt.want)
			}
		})
	}
}
