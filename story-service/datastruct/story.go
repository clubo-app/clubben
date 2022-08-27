package datastruct

import (
	"github.com/clubo-app/protobuf/story"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Story struct {
	Id            string   `db:"story_id"       validate:"required"`
	PartyId       string   `db:"party_id"       validate:"required"`
	UserId        string   `db:"user_id"        validate:"required"`
	Url           string   `db:"url"            validate:"required"`
	TaggedFriends []string `db:"tagged_friends"`
}

func (s Story) ToGRPCStory() *story.Story {
	story := story.Story{
		Id:            s.Id,
		PartyId:       s.PartyId,
		UserId:        s.UserId,
		Url:           s.Url,
		TaggedFriends: s.TaggedFriends,
	}

	c, err := ksuid.Parse(s.Id)
	if err == nil {
		story.CreatedAt = timestamppb.New(c.Time())
	}

	return &story
}
