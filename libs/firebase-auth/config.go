package firebaseauth

import (
	"errors"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

type FirebaseUser struct {
	UserID        string
	Email         string
	EmailVerified bool
	ProviderId    string
}

type Config struct {
	FirebaseApp  *firebase.App
	ErrorHandler fiber.ErrorHandler

	// Context key to store user information from the token into context.
	// Optional. Default: "user".
	ContextKey string

	// Skip Email Check.
	// Optional. Default: false
	CheckEmailVerified bool

	// Call next even on auth error.
	// Optional. Default: false
	AuthOptional bool
}

var (
	missingErrMsg    = errors.New("Missing or malformed ID Token")
	invalidToken     = errors.New("Invalid or expiret ID Token")
	emailNotVerified = errors.New("Email is not verified")
)

func makeCfg(config []Config) (cfg Config) {
	if len(config) > 0 {
		cfg = config[0]
	}

	// Check Mandatory FirebaseApp is provided
	if cfg.FirebaseApp == nil {
		fmt.Println("****************************************************************")
		fmt.Println("firebaseauth :: Error PLEASE PASS Firebase App in Config")
		fmt.Println("*****************************************************************")
	}

	if cfg.ErrorHandler == nil {
		cfg.ErrorHandler = func(c *fiber.Ctx, err error) error {
			if err.Error() != invalidToken.Error() {
				return c.Status(fiber.StatusBadRequest).SendString(err.Error())
			}
			return c.Status(fiber.StatusUnauthorized).SendString(invalidToken.Error())
		}
	}

	if cfg.ContextKey == "" {
		cfg.ContextKey = "user"
	}

	return cfg
}
