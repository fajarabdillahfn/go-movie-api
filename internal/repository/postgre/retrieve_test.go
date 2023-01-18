package postgre

import (
	"context"
	"fmt"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func Test_repository_GetAll(t *testing.T) {
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Movie
		wantErr assert.ErrorAssertionFunc
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				conn: tt.fields.conn,
			}
			got, err := r.GetAll(tt.args.ctx)
			if !tt.wantErr(t, err, fmt.Sprintf("GetAll(%v)", tt.args.ctx)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetAll(%v)", tt.args.ctx)
		})
	}
}

func Test_repository_GetById(t *testing.T) {
	type fields struct {
		conn *gorm.DB
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
		wantErr assert.ErrorAssertionFunc
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				conn: tt.fields.conn,
			}
			got, err := r.GetById(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("GetById(%v, %v)", tt.args.ctx, tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetById(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}
