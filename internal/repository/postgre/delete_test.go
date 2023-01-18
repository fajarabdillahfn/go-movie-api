package postgre

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func Test_repository_Delete(t *testing.T) {
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
		wantErr assert.ErrorAssertionFunc
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				conn: tt.fields.conn,
			}
			tt.wantErr(t, r.Delete(tt.args.ctx, tt.args.id), fmt.Sprintf("Delete(%v, %v)", tt.args.ctx, tt.args.id))
		})
	}
}
