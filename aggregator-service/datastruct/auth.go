package datastruct

import (
	ag "github.com/clubo-app/clubben/protobuf/auth"
)

type AggregatedAccount struct {
	Id      string            `json:"id"`
	Profile AggregatedProfile `json:"profile"`
	Email   string            `json:"email"`
}

type LoginResponse struct {
	Tokens  ag.TokenResponse  `json:"tokens"`
	Account AggregatedAccount `json:"account"`
}
