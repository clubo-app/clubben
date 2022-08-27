package rpc

import (
	"log"
	"net"
	"strings"

	"github.com/clubo-app/clubben/story-service/service"
	sg "github.com/clubo-app/protobuf/story"
	"google.golang.org/grpc"
)

type storyServer struct {
	ss service.StoryService
	us service.UploadService
	sg.UnimplementedStoryServiceServer
}

func NewStoryServer(ss service.StoryService, us service.UploadService) sg.StoryServiceServer {
	return &storyServer{ss: ss, us: us}
}

func Start(s sg.StoryServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	sg.RegisterStoryServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
