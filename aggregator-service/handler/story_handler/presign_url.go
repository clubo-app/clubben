package storyhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h *storyHandler) PresignURL(c *fiber.Ctx) error {
	key := c.Params("key")

	res, err := h.storyClient.PresignURL(c.Context(), &story.PresignURLRequest{Key: key})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
