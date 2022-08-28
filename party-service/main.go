package main

import (
	"log"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/party-service/config"
	"github.com/clubo-app/clubben/party-service/repository"
	"github.com/clubo-app/clubben/party-service/rpc"
	"github.com/clubo-app/clubben/party-service/service"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("Party Service")}
	stream, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.Close()

	d, err := repository.NewPartyRepository(c.POSTGRES_URL_PARTY_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	s := service.NewPartyService(d, stream)

	p := rpc.NewPartyServer(s, stream)
	rpc.Start(p, c.PORT)
}
