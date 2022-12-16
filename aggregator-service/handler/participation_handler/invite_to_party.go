package participationhandler

import (
	"time"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/durationpb"
)

func (h *participationHandler) InviteToParty(c *fiber.Ctx) error {
	pId := c.Params("pid")
	uId := c.Params("uid")
	validFor := c.Query("validFor")
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	duration, err := time.ParseDuration(validFor)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid valid_for parameter")
	}

	i, err := h.participationClient.InviteToParty(c.Context(), &participation.InviteToPartyRequest{
		InviterId: user.UserID,
		UserId:    uId,
		PartyId:   pId,
		ValidFor:  durationpb.New(duration),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.
		PartyInviteToAgg(i).
		AddInviter(&datastruct.AggregatedProfile{Id: i.InviterId}).
		AddProfile(&datastruct.AggregatedProfile{Id: i.UserId}).
		AddParty(&datastruct.AggregatedParty{Id: pId})

	return c.Status(fiber.StatusCreated).JSON(res)
}
