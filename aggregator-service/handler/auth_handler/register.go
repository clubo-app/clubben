package authhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	Username    string `json:"username,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	GoogleToken string `json:"google_token"`
	AppleToken  string `json:"apple_token"`
}

func (h authHandler) Register(c *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	nameTaken, err := h.profileClient.UsernameTaken(c.Context(), &profile.UsernameTakenRequest{
		Username: req.Username,
	})
	if err != nil || nameTaken.Taken {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Username already taken")
	}

	a, err := h.authClient.RegisterUser(c.Context(), &ag.RegisterUserRequest{
		Email:       req.Email,
		Password:    req.Password,
		GoogleToken: req.GoogleToken,
		AppleToken:  req.AppleToken,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	p, err := h.profileClient.CreateProfile(c.Context(), &profile.CreateProfileRequest{
		Id:        a.Account.Id,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.LoginResponse{
		Account: datastruct.AggregatedAccount{
			Id:      a.Account.Id,
			Profile: datastruct.ProfileToAgg(p),
			Email:   a.Account.Email,
		},
		Tokens: *a.Tokens,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
