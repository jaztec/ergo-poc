package main

import (
	"context"
	"ergo.services/application/observer"
	"ergo.services/ergo"
	"ergo.services/ergo/gen"
	"ergo.services/ergo/lib"
	"flag"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaztec/ergo-poc/pkg/application"
)

var (
	databaseDSN string
)

func init() {
	flag.StringVar(&databaseDSN, "dsn", "", "database connection string")
}

func main() {
	var options gen.NodeOptions

	flag.Parse()

	if databaseDSN == "" {
		panic("DatabaseDSN is required")
	}

	conn, err := pgxpool.New(context.Background(), databaseDSN)
	if err != nil {
		panic(err)
	}

	options.Applications = []gen.ApplicationBehavior{
		observer.CreateApp(observer.Options{}),
		application.NewApp(conn),
	}

	options.Network.Cookie = lib.RandomString(16)

	node, err := ergo.StartNode("tasks@localhost", options)
	if err != nil {
		panic(err)
	}

	node.Wait()
}
