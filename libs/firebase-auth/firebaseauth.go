package firebaseauth

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cfg := makeCfg(config)

		IDToken := c.Get(fiber.HeaderAuthorization)
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
		if err != nil {
			return cfg.ErrorHandler(c, invalidToken)
		}

		if cfg.CheckEmailVerified && !token.Claims["email_verified"].(bool) {
			return cfg.ErrorHandler(c, errors.New("Email not verified"))
		}

		if token != nil {
			c.Locals(cfg.ContextKey, FirebaseUser{
				UserID:        token.Claims["user_id"].(string),
				Email:         token.Claims["email"].(string),
				EmailVerified: token.Claims["email_verified"].(bool),
				ProviderId:    token.Claims["provider_id"].(string),
			})

			return c.Next()
		}
		return cfg.ErrorHandler(c, invalidToken)
	}
}
