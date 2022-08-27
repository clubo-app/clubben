package rpc

import (
	"log"
	"net"
	"strings"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/profile-service/service"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	"google.golang.org/grpc"
)

type profileServer struct {
	ps service.ProfileService
	up service.UploadService

	stream stream.Stream

	pg.UnimplementedProfileServiceServer
}

func NewProfileServer(ps service.ProfileService, up service.UploadService, stream stream.Stream) pg.ProfileServiceServer {
	return &profileServer{ps: ps, up: up, stream: stream}
}

func Start(s pg.ProfileServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	pg.RegisterProfileServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
