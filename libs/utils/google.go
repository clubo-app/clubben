package utils

import (
	"context"
	"errors"

	"google.golang.org/api/idtoken"
)

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
}

type GoogleManager struct {
	clientId string
}

// The clientId is the GCP OAuth2 clientId. It can be accessed through the Credentials Service.
func NewGoogleManager(clientId string) GoogleManager {
	return GoogleManager{clientId: clientId}
}

func (g GoogleManager) ValidateGoogleToken(ctx context.Context, token string) (GoogleClaims, error) {
	p, err := idtoken.Validate(ctx, token, g.clientId)
	if err != nil {
		return GoogleClaims{}, err
	}

	if p.Issuer != "accounts.google.com" {
		return GoogleClaims{}, errors.New("iss is invalid")
	}

	claims := GoogleClaims{
		Email:         p.Claims["email"].(string),
		EmailVerified: p.Claims["email_verified"].(bool),
		FirstName:     p.Claims["given_name"].(string),
		LastName:      p.Claims["family_name"].(string),
	}

	return claims, nil
}
