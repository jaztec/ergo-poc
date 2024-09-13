package application

import (
	"ergo.services/ergo/gen"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaztec/ergo-poc/pkg/db"
	"github.com/jaztec/ergo-poc/pkg/web"
)

type App struct {
	pool *pgxpool.Pool
}

func (a App) Load(_ gen.Node, _ ...any) (gen.ApplicationSpec, error) {
	return gen.ApplicationSpec{
		Name:        "web",
		Description: "Web application",
		Mode:        gen.ApplicationModeTransient,
		Group: []gen.ApplicationMemberSpec{{
			Name:    "pool",
			Factory: web.NewWebPool,
		}, {
			Name:    "database",
			Factory: db.NewPoolFac(a.pool),
		}},
	}, nil
}

func (a App) Start(_ gen.ApplicationMode) {
}

func (a App) Terminate(_ error) {
}

func NewApp(pool *pgxpool.Pool) gen.ApplicationBehavior {
	return &App{pool: pool}
}
