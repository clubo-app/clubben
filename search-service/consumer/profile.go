package consumer

import (
	"context"
	"log"

	"github.com/clubo-app/clubben/protobuf/events"
	"github.com/clubo-app/clubben/search-service/datastruct"
	"github.com/clubo-app/clubben/search-service/repository"
)

type profileConsumer struct {
	userRepo *repository.ProfileRepository
}

func NewProfileConsumer(repo *repository.ProfileRepository) profileConsumer {
	return profileConsumer{
		userRepo: repo,
	}
}

func (c profileConsumer) ProfileCreated(p *events.ProfileCreated) {
	if p == nil {
		return
	}
	err := c.userRepo.PutProfile(context.Background(), datastruct.Profile{
		Id:        p.Profile.Id,
		Username:  p.Profile.Username,
		Firstname: p.Profile.Firstname,
		Lastname:  p.Profile.Lastname,
	})
	if err != nil {
		log.Println("Error inserting Profile: ", err)
	}
}

func (c profileConsumer) ProfileUpdate(p *events.ProfileUpdated) {
	if p == nil {
		return
	}
	err := c.userRepo.UpdateProfile(context.Background(), datastruct.Profile{
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
	err := c.userRepo.RemoveProfile(context.Background(), p.Id)
	if err != nil {
		log.Println("Error removing Profile: ", err)
	}
}
