package rpc

import (
	"log"
	"net"
	"strings"

	"github.com/clubo-app/clubben/protobuf/search"
	"github.com/clubo-app/clubben/search-service/repository"
	"google.golang.org/grpc"
)

type searchServer struct {
	profile repository.ProfileRepository
	party   repository.PartyRepository

	search.UnimplementedSearchServiceServer
}

func NewSearchServer(profile repository.ProfileRepository, party repository.PartyRepository) search.SearchServiceServer {
	return &searchServer{profile: profile, party: party}
}

func Start(s search.SearchServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	search.RegisterSearchServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
