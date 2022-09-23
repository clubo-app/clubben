package partyhandler

import (
	"sync"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/clubo-app/clubben/protobuf/profile"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UpdatePartyReq struct {
	Title           string                `json:"title"`
	Description     string                `json:"description,omitempty"`
	Lat             float32               `json:"lat"`
	Lon             float32               `json:"lon"`
	MusicGenre      string                `json:"music_genre"`
	MaxParticipants int32                 `json:"max_participants"`
	StreetAddress   string                `json:"street_address"`
	PostalCode      string                `json:"postal_code"`
	State           string                `json:"state"`
	Country         string                `json:"country"`
	EntryDate       timestamppb.Timestamp `json:"entry_date"`
}

func (h partyHandler) UpdateParty(c *fiber.Ctx) error {
	req := new(UpdatePartyReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user := middleware.ParseUser(c)
	pId := c.Params("id")
	if pId == "" {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Party Id is required")
	}

	p, err := h.partyClient.UpdateParty(c.Context(), &party.UpdatePartyRequest{
		PartyId:       pId,
		RequesterId:   user.Sub,
		Title:         req.Title,
		Description:   req.Description,
		Lat:           req.Lat,
		Long:          req.Lon,
		MusicGenre:    req.MusicGenre,
		StreetAddress: req.StreetAddress,
		PostalCode:    req.PostalCode,
		State:         req.State,
		Country:       req.Country,
		EntryDate:     &req.EntryDate,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	party := datastruct.PartyToAgg(p)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		profileRes, _ := h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})
		party.AddCreator(datastruct.ProfileToAgg(profileRes))
	}()

	go func() {
		storyRes, _ := h.storyClient.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: p.Id})
		if storyRes != nil {
			party.AddStory(storyRes.Stories)
		}
	}()

	wg.Wait()

	return c.Status(fiber.StatusOK).JSON(party)
}
