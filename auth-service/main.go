package main

import (
	"log"

	"github.com/clubo-app/clubben/auth-service/config"
	"github.com/clubo-app/clubben/auth-service/repository"
	"github.com/clubo-app/clubben/auth-service/rpc"
	"github.com/clubo-app/clubben/auth-service/service"
)

func main() {
	c := config.LoadConfig()

	r, err := repository.NewAccountRepository(c.POSTGRES_URL_AUTH_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	t := service.NewTokenManager(c.TOKEN_SECRET)
	goog := service.NewGoogleManager(c.GOOGLE_CLIENTID)
	pw := service.NewPasswordManager()

	as := service.NewAccountService(r)

	s := rpc.NewAuthServer(t, pw, goog, as)

	rpc.Start(s, c.PORT)
}
