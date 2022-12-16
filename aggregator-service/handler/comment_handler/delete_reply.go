package commenthandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/comment"
	"github.com/gofiber/fiber/v2"
)

func (h *commentHandler) DeleteReply(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	cId := c.Params("id")
	rId := c.Params("rId")

	res, err := h.commentClient.DeleteReply(c.Context(), &comment.DeleteReplyRequest{AuthorId: user.UserID, CommentId: cId, ReplyId: rId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
