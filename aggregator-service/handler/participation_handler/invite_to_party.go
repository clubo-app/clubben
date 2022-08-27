package participationhandler

import (
	"time"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/durationpb"
)

type InviteRequest struct {
	ValidFor string `json:"valid_for"`
}

func (h participationHandler) InviteToParty(c *fiber.Ctx) error {
	req := new(InviteRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	pId := c.Params("pid")
	uId := c.Params("uid")
	user := middleware.ParseUser(c)

	duration, err := time.ParseDuration(req.ValidFor)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid valid_for parameter")
	}

	i, err := h.participationC.InviteToParty(c.Context(), &participation.InviteToPartyRequest{
		InviterId: user.Sub,
		UserId:    uId,
		PartyId:   pId,
		ValidFor:  durationpb.New(duration),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.
		PartyInviteToAgg(i).
		AddInviter(datastruct.AggregatedProfile{Id: i.InviterId}).
		AddUser(datastruct.AggregatedProfile{Id: i.UserId})

	return c.Status(fiber.StatusCreated).JSON(res)
}
