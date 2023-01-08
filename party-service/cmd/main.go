package main

import (
	"log"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/party-service/config"
	"github.com/clubo-app/clubben/party-service/internal/consumer"
	"github.com/clubo-app/clubben/party-service/internal/repository"
	"github.com/clubo-app/clubben/party-service/internal/rpc"
	"github.com/clubo-app/clubben/party-service/internal/service"
	"github.com/nats-io/nats.go"
)

func main() {
	c := config.LoadConfig()

	opts := []nats.Option{nats.Name("Party Service")}
	stream, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.Close()

	repo, err := repository.NewPartyRepository(c.POSTGRES_URL_PARTY_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()

	s := service.NewPartyService(repo, &stream)

	con := consumer.New(&stream, repo)
	go con.Start()

	p := rpc.NewPartyServer(s, stream)
	rpc.Start(p, c.PORT)
}
