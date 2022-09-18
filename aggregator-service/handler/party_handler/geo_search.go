package partyhandler

import (
	"strconv"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	pg "github.com/clubo-app/clubben/protobuf/party"
	"github.com/gofiber/fiber/v2"
)

func (h partyHandler) GeoSearch(c *fiber.Ctx) error {
	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	offsetStr := c.Query("offset")
	offset, _ := strconv.ParseInt(offsetStr, 10, 32)

	latStr := c.Params("lat")
	lat, err := strconv.ParseFloat(latStr, 32)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	lonStr := c.Params("lon")
	lon, err := strconv.ParseFloat(lonStr, 32)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	radiusStr := c.Query("radius")
	radius, _ := strconv.ParseInt(radiusStr, 10, 32)

	isPublicStr := c.Query("is_public")
	isPublic, _ := strconv.ParseBool(isPublicStr)

	parties, err := h.partyClient.GeoSearch(c.Context(), &pg.GeoSearchRequest{
		Lat:      float32(lat),
		Long:     float32(lon),
		Limit:    int32(limit),
		Offset:   int32(offset),
		Radius:   int32(radius),
		IsPublic: isPublic,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	aggP := make([]*datastruct.AggregatedParty, len(parties.Parties))
	for i, p := range parties.Parties {
		aggP[i] = datastruct.PartyToAgg(p)
	}

	return c.Status(fiber.StatusOK).JSON(aggP)
}
