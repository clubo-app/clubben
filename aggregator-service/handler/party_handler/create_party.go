package partyhandler

import (
	"time"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreatePartyReq struct {
	Title           string    `json:"title"`
	Description     string    `json:"description,omitempty"`
	Lat             float32   `json:"lat"`
	Long            float32   `json:"long"`
	IsPublic        bool      `json:"is_public"`
	MusicGenre      string    `json:"music_genre"`
	MaxParticipants int32     `json:"max_participants"`
	StreetAddress   string    `json:"street_address"`
	PostalCode      string    `json:"postal_code"`
	State           string    `json:"state"`
	Country         string    `json:"country"`
	EntryDate       time.Time `json:"entry_date"`
}

func (h partyGatewayHandler) CreateParty(c *fiber.Ctx) error {
	req := new(CreatePartyReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user := middleware.ParseUser(c)

	p, err := h.pc.CreateParty(c.Context(), &party.CreatePartyRequest{
		RequesterId:     user.Sub,
		Title:           req.Title,
		Description:     req.Description,
		Lat:             req.Lat,
		Long:            req.Long,
		IsPublic:        req.IsPublic,
		MusicGenre:      req.MusicGenre,
		MaxParticipants: req.MaxParticipants,
		StreetAddress:   req.StreetAddress,
		PostalCode:      req.PostalCode,
		State:           req.State,
		Country:         req.Country,
		EntryDate:       timestamppb.New(req.EntryDate),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.prf.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})

	res := datastruct.PartyToAgg(p).AddCreator(datastruct.ProfileToAgg(profileRes))
	return c.Status(fiber.StatusCreated).JSON(res)
}
