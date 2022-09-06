package consumer

import (
	"context"
	"log"

	"github.com/clubo-app/clubben/protobuf/events"
	"github.com/clubo-app/clubben/search-service/datastruct"
	"github.com/clubo-app/clubben/search-service/repository"
)

type profileConsumer struct {
	repo *repository.ProfileRepository
}

func NewProfileConsumer(repo *repository.ProfileRepository) profileConsumer {
	return profileConsumer{
		repo: repo,
	}
}

func (c profileConsumer) ProfileCreated(p *events.ProfileCreated) {
	if p == nil {
		return
	}
	err := c.repo.PutProfile(context.Background(), datastruct.Profile{
		Id:        p.Profile.Id,
		Username:  p.Profile.Username,
		Firstname: p.Profile.Firstname,
		Lastname:  p.Profile.Lastname,
	})
	if err != nil {
		log.Println("Error inserting Profile: ", err)
	}
	log.Printf("Profile Created %+v", p)
}

func (c profileConsumer) ProfileUpdate(p *events.ProfileUpdated) {
	if p == nil {
		return
	}
	err := c.repo.UpdateProfile(context.Background(), datastruct.Profile{
		Id:        p.Profile.Id,
		Username:  p.Profile.Username,
		Firstname: p.Profile.Firstname,
		Lastname:  p.Profile.Lastname,
	})
	if err != nil {
		log.Println("Error updating Profile: ", err)
	}
}

func (c profileConsumer) ProfileDeleted(p *events.UserDeleted) {
	if p == nil {
		return
	}
	err := c.repo.RemoveProfile(context.Background(), p.Id)
	if err != nil {
		log.Println("Error removing Profile: ", err)
	}
}
