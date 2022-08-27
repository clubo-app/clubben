package repository

import (
	"embed"
	"fmt"
	"net/url"

	pg "github.com/clubo-app/protobuf/profile"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func (p Profile) ToGRPCProfile() *pg.Profile {
	return &pg.Profile{
		Id:        p.ID,
		Username:  p.Username,
		Firstname: p.Firstname,
		Lastname:  p.Lastname.String,
		Avatar:    p.Avatar.String,
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
