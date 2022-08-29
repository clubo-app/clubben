package main

import (
	"log"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/profile-service/config"
	"github.com/clubo-app/clubben/profile-service/repository"
	"github.com/clubo-app/clubben/profile-service/rpc"
	"github.com/clubo-app/clubben/profile-service/service"
	"github.com/nats-io/nats.go"
)

func main() {
	c := config.LoadConfig()

	opts := []nats.Option{nats.Name("User Service")}
	stream, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.Close()

	r, err := repository.NewProfileRepository(c.POSTGRES_URL_PROFILE_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	up := service.NewUploadService(c.SPACES_ENDPOINT, c.SPACES_TOKEN)
	ps := service.NewProfileService(r)

	p := rpc.NewProfileServer(ps, up, stream)

	rpc.Start(p, c.PORT)
}
