package storyhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h storyHandler) DeleteStory(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	sId := c.Params("id")

	res, err := h.storyClient.DeleteStory(c.Context(), &story.DeleteStoryRequest{RequesterId: user.Sub, StoryId: sId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
