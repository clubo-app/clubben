package main

import (
	"log"
	"net/http"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/search-service/config"
	"github.com/clubo-app/clubben/search-service/consumer"
	"github.com/clubo-app/clubben/search-service/repository"
	"github.com/jonashiltl/govespa"
	"github.com/nats-io/nats.go"
)

func main() {
	c := config.LoadConfig()

	opts := []nats.Option{nats.Name("Search Service")}
	stream, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.Close()

	vespa := govespa.NewClient(govespa.NewClientParams{
		BaseUrl:    c.VESPA_URL,
		HttpClient: newHttp(),
	})

	pRepo := repository.NewProfileRepository(vespa)

	profileCon := consumer.NewProfileConsumer(&pRepo)
	con := consumer.NewConsumer(&stream, profileCon)

	con.Start()
}

func newHttp() *http.Client {
	return http.DefaultClient
}
