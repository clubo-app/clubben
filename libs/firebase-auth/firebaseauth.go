package firebaseauth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cfg := makeCfg(config)

		bearerToken := c.Get(fiber.HeaderAuthorization)
		IDToken := strings.TrimPrefix(bearerToken, "Bearer ")

		// Validate token
		if len(IDToken) == 0 {
			return cfg.ErrorHandler(c, missingErrMsg)
		}

		if cfg.FirebaseApp == nil {
			fmt.Println("****************************************************************")
			fmt.Println("firebaseauth :: Error PLEASE PASS Firebase App in Config")
			fmt.Println("*****************************************************************")
			return cfg.ErrorHandler(c, errors.New("Missing Firebase App Object"))
		}

		client, err := cfg.FirebaseApp.Auth(context.Background())
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		// Verify IDToken
		token, err := client.VerifyIDToken(context.Background(), IDToken)
		if err != nil && !cfg.AuthOptional {
			println("Error verifyIDToken: ", err.Error())
			return cfg.ErrorHandler(c, invalidToken)
		}

		if token == nil {
			if cfg.AuthOptional {
				return c.Next()
			} else {
				println("Token is nil")
				return cfg.ErrorHandler(c, invalidToken)
			}
		}

		if cfg.CheckEmailVerified && !cfg.AuthOptional && !token.Claims["email_verified"].(bool) {
			return cfg.ErrorHandler(c, errors.New("Email not verified"))
		}

		c.Locals(contextKey, FirebaseUser{
			UserID:        token.Claims["user_id"].(string),
			Email:         token.Claims["email"].(string),
			EmailVerified: token.Claims["email_verified"].(bool),
		})

		return c.Next()
	}
}

func GetUser(c *fiber.Ctx) (FirebaseUser, *fiber.Error) {
	user, ok := c.Locals(contextKey).(FirebaseUser)
	if !ok {
		return FirebaseUser{}, fiber.NewError(fiber.ErrBadRequest.Code, "User not avilable in context")
	}
	return user, nil
}
