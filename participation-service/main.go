package main

import (
	"log"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/participation-service/config"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/participation-service/rpc"
	"github.com/clubo-app/clubben/participation-service/service"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/go-playground/validator"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("Relation Service")}
	stream, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.Close()

	cqlx, err := repository.NewDB(c.CQL_KEYSPACE, c.CQL_HOSTS)
	if err != nil {
		log.Fatal(err)
	}
	defer cqlx.Close()

	pc, err := party.NewClient(c.PARTY_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to party service: %v", err)
	}

	dao := repository.NewDAO(cqlx)
	val := validator.New()

	pr := dao.NewParticipantRepository(val)
	ir := dao.NewInviteRepository(val)

	ps := service.NewParticipantService(pr, pc)
	is := service.NewInviteService(ir, pc, ps)

	r := rpc.NewParticipationServer(ps, is, stream)
	rpc.Start(r, c.PORT)
}
