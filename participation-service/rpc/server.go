package rpc

import (
	"log"
	"net"
	"strings"

	"github.com/clubo-app/clubben/participation-service/service"
	"github.com/clubo-app/packages/stream"
	"github.com/clubo-app/protobuf/participation"
	"google.golang.org/grpc"
)

type server struct {
	pi     service.Invite
	pp     service.Participant
	stream stream.Stream
	participation.UnimplementedParticipationServiceServer
}

func NewParticipationServer(pp service.Participant, pi service.Invite, stream stream.Stream) participation.ParticipationServiceServer {
	return &server{
		pi:     pi,
		pp:     pp,
		stream: stream,
	}
}

func Start(s participation.ParticipationServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	participation.RegisterParticipationServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
