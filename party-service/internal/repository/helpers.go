package repository

import (
	"embed"
	"fmt"
	"net/url"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
)

func (p Party) ToGRPCParty() *pbparty.Party {
	id, err := ksuid.Parse(p.ID)
	if err != nil {
		return nil
	}

	return &pbparty.Party{
		Id:              p.ID,
		UserId:          p.UserID,
		Title:           p.Title,
		Description:     p.Description.String,
		IsPublic:        p.IsPublic,
		MaxParticipants: p.MaxParticipants,
		Lat:             float32(p.Location.Lat()),
		Long:            float32(p.Location.Lon()),
		StreetAddress:   p.StreetAddress.String,
		PostalCode:      p.PostalCode.String,
		State:           p.State.String,
		Country:         p.Country.String,
		EntryDate:       timestamppb.New(p.EntryDate),
		CreatedAt:       timestamppb.New(id.Time()),
	}
}

const version = 1

//go:embed migrations/*.sql
var fs embed.FS

func validateSchema(url url.URL) error {
	url.Scheme = "pgx"
	urlf := fmt.Sprintf("%v%v", url.String(), "?sslmode=disable")

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, urlf)
	if err != nil {
		return err
	}

	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	defer m.Close()
	return nil
}
