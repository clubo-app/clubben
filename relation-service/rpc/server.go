package rpc

import (
	"log"
	"net"
	"strings"

	"github.com/clubo-app/clubben/libs/stream"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/clubo-app/clubben/relation-service/service"
	"google.golang.org/grpc"
)

type relationServer struct {
	fs     service.FriendRelation
	fp     service.FavoriteParty
	stream stream.Stream
	rg.UnimplementedRelationServiceServer
}

func NewRelationServer(fs service.FriendRelation, fp service.FavoriteParty, stream stream.Stream) rg.RelationServiceServer {
	return &relationServer{
		fs:     fs,
		fp:     fp,
		stream: stream,
	}
}

func Start(s rg.RelationServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	rg.RegisterRelationServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
