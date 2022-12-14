package main

import (
	"strings"

	"github.com/clubo-app/clubben/libs/cqlx"
	"github.com/scylladb/gocqlx/v2"
)

func newCluster(keyspace, hosts string) (*gocqlx.Session, error) {
	h := strings.Split(hosts, ",")

	manager := cqlx.NewManager(keyspace, h)

	if err := manager.CreateKeyspace(keyspace, 1, "SimpleStrategy"); err != nil {
		return nil, err
	}

	session, err := manager.Connect()
	if err != nil {
		return nil, err
	}
	return &session, nil
}
