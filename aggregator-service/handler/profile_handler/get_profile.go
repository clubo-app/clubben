package profilehandler

import (
	"sync"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *profileHandler) GetProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	p, err := h.profileClient.GetProfile(c.Context(), &pg.GetProfileRequest{
		Id: id,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	if p == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Profile not found")
	}

	res := datastruct.ProfileToAgg(p)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		relation := new(rg.FriendRelation)
		// if somebody wants the Profile of somebody else we also return the friendship status between them two
		if id != user.UserID {
			relation, _ = h.relationClient.GetFriendRelation(c.Context(), &rg.GetFriendRelationRequest{UserId: user.UserID, FriendId: id})
		}
		if relation != nil {
			fs := datastruct.ParseFriendShipStatus(user.UserID, relation)
			res.AddFs(fs)
		}
	}()

	go func() {
		defer wg.Done()
		friendCountRes, _ := h.relationClient.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: p.Id})
		if friendCountRes != nil {
			res.AddFCount(friendCountRes.FriendCount)
		}
	}()

	wg.Wait()

	return c.Status(fiber.StatusOK).JSON(res)
}
