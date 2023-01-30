package database

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

type Pool interface {
	Exec(sql squirrel.Sqlizer) error
}

type pool struct {
	p *pgx.ConnPool
}

func New(ctx context.Context) (Pool, error) {
	p, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "",
			Port:     0,
			Database: "",
			User:     "",
			Password: "",
		},
		MaxConnections: 0,
		AfterConnect:   nil,
	})
	if err != nil {
		return nil, fmt.Errorf("can't start database pool, err: %+v", err)
	}

	c, err := p.Acquire()
	if err != nil {
		return nil, fmt.Errorf("can't acquire connection from pool, err: %+v", err)
	}

	defer func() {
		_ = c.Close()
	}()

	return pool{p: p}, c.Ping(ctx)
}

func (s pool) Exec(sql squirrel.Sqlizer) error {
	q, a, err := sql.ToSql()
	if err != nil {
		return err
	}

	_, err = s.p.Exec(q, a)
	return err
}
