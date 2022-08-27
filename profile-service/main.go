package main

import (
	"log"

	"github.com/clubo-app/clubben/profile-service/config"
	"github.com/clubo-app/clubben/profile-service/repository"
	"github.com/clubo-app/clubben/profile-service/rpc"
	"github.com/clubo-app/clubben/profile-service/service"
	"github.com/clubo-app/packages/stream"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("User Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()
	stream := stream.New(nc)

	r, err := repository.NewProfileRepository(c.DB_USER, c.DB_PW, c.DB_NAME, c.DB_HOST, c.DB_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	up := service.NewUploadService(c.SPACES_ENDPOINT, c.SPACES_TOKEN)
	ps := service.NewProfileService(r)

	p := rpc.NewProfileServer(ps, up, stream)

	rpc.Start(p, c.PORT)
}
