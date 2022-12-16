package profilehandler

import (
	"sync"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type UpdateRequest struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

func (h *profileHandler) UpdateUser(c *fiber.Ctx) error {
	req := new(UpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	a := new(pbauth.Account)
	p := new(pg.Profile)
	var friendCount int

	var wg sync.WaitGroup
	wg.Add(3)

	var pErr error
	go func() {
		defer wg.Done()

		if req.Email != "" || req.Password != "" {
			a, pErr = h.authClient.UpdateAccount(c.Context(), &pbauth.UpdateAccountRequest{
				Id:       user.UserID,
				Email:    req.Email,
				Password: req.Password,
			})
		}
	}()

	var aErr error
	go func() {
		defer wg.Done()

		if req.Username != "" || req.Firstname != "" || req.Lastname != "" || req.Avatar != "" {
			p, aErr = h.profileClient.UpdateProfile(c.Context(), &pg.UpdateProfileRequest{
				Id:        user.UserID,
				Username:  req.Username,
				Firstname: req.Firstname,
				Lastname:  req.Lastname,
				Avatar:    req.Avatar,
			})
		}
	}()

	go func() {
		defer wg.Done()
		friendCountRes, _ := h.relationClient.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: user.UserID})
		if friendCountRes != nil {
			friendCount = int(friendCountRes.FriendCount)
		}
	}()

	wg.Wait()

	if pErr != nil {
		return utils.ToHTTPError(pErr)
	}
	if aErr != nil {
		return utils.ToHTTPError(aErr)
	}

	res := datastruct.AggregatedAccount{
		Id:      p.Id,
		Profile: datastruct.ProfileToAgg(p).AddFCount(uint32(friendCount)),
		Email:   a.Email,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
