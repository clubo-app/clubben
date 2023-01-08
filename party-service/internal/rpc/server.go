package rpc

import (
	"log"
	"net"
	"strings"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/party-service/internal/service"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"google.golang.org/grpc"
)

type partyServer struct {
	ps     service.PartyService
	stream stream.Stream
	pbparty.UnimplementedPartyServiceServer
}

func NewPartyServer(ps service.PartyService, stream stream.Stream) pbparty.PartyServiceServer {
	return &partyServer{
		ps:     ps,
		stream: stream,
	}
}

func Start(s pbparty.PartyServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	pbparty.RegisterPartyServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
