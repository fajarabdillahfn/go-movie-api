package postgre

import (
	"context"
	"fmt"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func Test_repository_Create(t *testing.T) {
	type fields struct {
		conn *gorm.DB
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
		wantErr assert.ErrorAssertionFunc
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				conn: tt.fields.conn,
			}
			got, err := r.Create(tt.args.ctx, tt.args.newMovie)
			if !tt.wantErr(t, err, fmt.Sprintf("Create(%v, %v)", tt.args.ctx, tt.args.newMovie)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Create(%v, %v)", tt.args.ctx, tt.args.newMovie)
		})
	}
}
