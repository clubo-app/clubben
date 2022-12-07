package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/clubo-app/clubben/auth-service/config"
	"github.com/clubo-app/clubben/auth-service/internal/repository"
	"github.com/clubo-app/clubben/auth-service/internal/rpc"
	"github.com/clubo-app/clubben/auth-service/internal/service"
	"google.golang.org/api/option"
)

func main() {
	c := config.LoadConfig()

	opt := option.WithCredentialsFile(c.GOOGLE_APPLICATION_CREDENTIALS)
	firebase, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := firebase.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewFirebaseRepository(auth)

	as := service.NewAccountService(repo)

	s := rpc.NewAuthServer(as)

	rpc.Start(s, c.PORT)
}
