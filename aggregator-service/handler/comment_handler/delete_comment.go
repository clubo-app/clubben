package commenthandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/comment"
	"github.com/gofiber/fiber/v2"
)

func (h commentHandler) DeleteComment(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	cId := c.Params("id")
	pId := c.Params("pId")

	res, err := h.commentClient.DeleteComment(c.Context(), &comment.DeleteCommentRequest{AuthorId: user.UserID, PartyId: pId, CommentId: cId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
