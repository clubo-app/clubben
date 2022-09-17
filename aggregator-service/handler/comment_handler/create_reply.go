package commenthandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	cg "github.com/clubo-app/clubben/protobuf/comment"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h commentHandler) CreateReply(c *fiber.Ctx) error {
	req := new(cg.CreateReplyRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	cId := c.Params("id")
	user := middleware.ParseUser(c)

	req.AuthorId = user.Sub
	req.CommentId = cId

	r, err := h.commentClient.CreateReply(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: r.AuthorId})

	ar := datastruct.AggregatedReply{
		Id:        r.Id,
		CommentId: r.CommentId,
		Author:    profileRes,
		Body:      r.Body,
		CreatedAt: r.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(ar)
}
