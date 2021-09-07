package pgq

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/modern-go/reflect2"
)

type Pool struct {
	*pgxpool.Pool
}

func (p *Pool) QueryRow(ctx context.Context, sql string, row interface{}) {
	p.Pool.QueryRow(ctx, sql, )
}

type Row struct {
	pgx.Row
}

func (r *Row) Scan(dst interface{}) {
	rt := reflect2.TypeOfPtr(dst)
	num := rt.Elem().(reflect2.StructType).NumField()

	for i := 0; i < num; i++ {

	}
}