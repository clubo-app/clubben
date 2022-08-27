package main

import (
	"log"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/search-service/config"
	"github.com/clubo-app/clubben/search-service/consumer"
	"github.com/clubo-app/clubben/search-service/repository"
	"github.com/jonashiltl/govespa"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("Search Service")}
	stream, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.Close()

	vespa := govespa.NewClient(govespa.NewClientParams{
		BaseUrl: c.VESPA_URL,
	})

	pRepo := repository.NewProfileRepository(vespa)

	profileCon := consumer.NewProfileConsumer(&pRepo)
	con := consumer.NewConsumer(&stream, profileCon)

	con.Start()
}
