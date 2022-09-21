package profilehandler

import (
	"sync"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h profileHandler) GetProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	user := middleware.ParseUser(c)

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
		if id != user.Sub {
			relation, _ = h.relationClient.GetFriendRelation(c.Context(), &rg.GetFriendRelationRequest{UserId: user.Sub, FriendId: id})
		}
		if relation != nil {
			fs := datastruct.ParseFriendShipStatus(user.Sub, relation)
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
