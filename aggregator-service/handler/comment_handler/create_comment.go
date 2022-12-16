package commenthandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	cg "github.com/clubo-app/clubben/protobuf/comment"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h commentHandler) CreateComment(c *fiber.Ctx) error {
	req := new(cg.CreateCommentRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	pId := c.Params("id")
	user, useErr := firebaseauth.GetUser(c)
	if useErr != nil {
		return useErr
	}

	req.AuthorId = user.UserID
	req.PartyId = pId

	co, err := h.commentClient.CreateComment(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: co.AuthorId})

	ac := datastruct.AggregatedComment{
		Id:        co.Id,
		PartyId:   co.PartyId,
		Author:    profileRes,
		Body:      co.Body,
		CreatedAt: co.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(ac)
}
