package searchhandler

import (
	"strconv"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/search"
	"github.com/gofiber/fiber/v2"
)

func (h searchGatewayHandler) SearchParties(c *fiber.Ctx) error {
	q := c.Params("query")

	latStr := c.Query("lat")
	lat, _ := strconv.ParseFloat(latStr, 32)

	lonStr := c.Query("lon")
	lon, _ := strconv.ParseFloat(lonStr, 32)

	res, err := h.sc.SearchParties(c.Context(), &search.SearchPartiesRequest{
		Query: q,
		Lat:   float32(lat),
		Long:  float32(lon),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
