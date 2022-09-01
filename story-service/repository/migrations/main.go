package main

import (
	"context"
	"log"
	"strings"

	"github.com/clubo-app/clubben/libs/cqlx"
	"github.com/clubo-app/clubben/story-service/config"
	"github.com/clubo-app/clubben/story-service/repository/migrations/cql"
	"github.com/scylladb/gocqlx/v2/migrate"
)

func main() {
	c := config.LoadConfig()
	ctx := context.Background()

	h := strings.Split(c.CQL_HOSTS, ",")
	manager := cqlx.NewManager(c.CQL_KEYSPACE, h)

	if err := manager.CreateKeyspace(c.CQL_KEYSPACE, 1, "SimpleStrategy"); err != nil {
		log.Fatalln(err)
	}

	session, err := manager.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	if err := migrate.FromFS(ctx, session, cql.Files); err != nil {
		log.Fatal("Migrate: ", err)
	}
}
