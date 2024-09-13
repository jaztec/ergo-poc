package db

import (
	"ergo.services/ergo/act"
	"ergo.services/ergo/gen"
	"github.com/jackc/pgx/v5/pgxpool"
	db_gen "github.com/jaztec/ergo-poc/pkg/db/gen"
)

type Pool struct {
	act.Pool

	queries *db_gen.Queries
}

func (p *Pool) Init(args ...any) (act.PoolOptions, error) {
	var poolOptions act.PoolOptions

	poolOptions.WorkerFactory = NewWorker
	poolOptions.WorkerArgs = []any{p.queries}

	return poolOptions, nil
}

func NewPoolFac(conn *pgxpool.Pool) func() gen.ProcessBehavior {
	queries := db_gen.New(conn)
	return func() gen.ProcessBehavior {
		return &Pool{
			queries: queries,
		}
	}
}
