package authhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
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

	a, err := h.authClient.Register(c.Context(), &pbauth.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
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
		Account: datastruct.AccountToAgg(a.Account).AddProfile(datastruct.ProfileToAgg(p)),
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
