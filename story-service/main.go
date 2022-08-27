package main

import (
	"log"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/story-service/config"
	"github.com/clubo-app/clubben/story-service/repository"
	rpc "github.com/clubo-app/clubben/story-service/rpc"
	"github.com/clubo-app/clubben/story-service/service"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("Story Service")}
	stream, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.Close()

	cqlx, err := repository.NewDB(c.CQL_KEYSPACE, c.CQL_HOSTS)
	if err != nil {
		log.Fatal(err)
	}

	dao := repository.NewDAO(cqlx)
	val := validator.New()

	us := service.NewUploadService(c.SPACES_KEY, c.SPACES_ENDPOINT, c.SPACES_KEY)

	st := rpc.NewStoryServer(dao.NewStoryRepository(val), us)
	rpc.Start(st, c.PORT)
}
