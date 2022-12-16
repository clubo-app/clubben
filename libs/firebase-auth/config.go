package firebaseauth

import (
	"errors"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
)

type FirebaseUser struct {
	UserID        string
	Email         string
	EmailVerified bool
}

type errorResponse struct {
	Message string `json:"message"`
}

const contextKey = "user"

type Config struct {
	FirebaseApp  *firebase.App
	ErrorHandler fiber.ErrorHandler

	// Skip Email Check.
	// Optional. Default: false
	CheckEmailVerified bool

	// Call next even on auth error.
	// Optional. Default: false
	AuthOptional bool
}

var (
	missingErrMsg    = errors.New("Missing or malformed ID Token")
	invalidToken     = errors.New("Invalid or expired ID Token")
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
				return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Message: err.Error()})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse{Message: err.Error()})
		}
	}

	return cfg
}
