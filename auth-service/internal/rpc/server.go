package rpc

import (
	"log"
	"net"
	"strings"

	"github.com/clubo-app/clubben/auth-service/internal/service"
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"google.golang.org/grpc"
)

type authServer struct {
	accountService service.AccountService

	pbauth.UnimplementedAuthServiceServer
}

func NewAuthServer(ac service.AccountService) pbauth.AuthServiceServer {
	return &authServer{
		accountService: ac,
	}
}

func Start(s pbauth.AuthServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	pbauth.RegisterAuthServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
