package dto

import (
	"github.com/clubo-app/clubben/auth-service/repository"
)

type Account struct {
	ID            string
	Email         string
	EmailVerified bool
	EmailCode     string
	PasswordHash  string
	Provider      repository.NullProvider
	Type          repository.Type
}
