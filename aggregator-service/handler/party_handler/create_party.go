package partyhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreatePartyReq struct {
	Title           string                `json:"title"`
	Description     string                `json:"description,omitempty"`
	Lat             float32               `json:"lat"`
	Lon             float32               `json:"lon"`
	IsPublic        bool                  `json:"is_public"`
	MusicGenre      string                `json:"music_genre"`
	MaxParticipants int32                 `json:"max_participants"`
	StreetAddress   string                `json:"street_address"`
	PostalCode      string                `json:"postal_code"`
	State           string                `json:"state"`
	Country         string                `json:"country"`
	EntryDate       timestamppb.Timestamp `json:"entry_date"`
}

func (h *partyHandler) CreateParty(c *fiber.Ctx) error {
	req := new(CreatePartyReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	p, err := h.partyClient.CreateParty(c.Context(), &pbparty.CreatePartyRequest{
		RequesterId:     user.UserID,
		Title:           req.Title,
		Description:     req.Description,
		Lat:             req.Lat,
		Long:            req.Lon,
		IsPublic:        req.IsPublic,
		MusicGenre:      req.MusicGenre,
		MaxParticipants: req.MaxParticipants,
		StreetAddress:   req.StreetAddress,
		PostalCode:      req.PostalCode,
		State:           req.State,
		Country:         req.Country,
		EntryDate:       &req.EntryDate,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})

	res := datastruct.PartyToAgg(p).AddCreator(datastruct.ProfileToAgg(profileRes))
	return c.Status(fiber.StatusCreated).JSON(res)
}
