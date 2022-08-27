package consumer

import (
	"sync"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/protobuf/events"
)

type consumer struct {
	stream  *stream.Stream
	profile profileConsumer
}

func NewConsumer(stream *stream.Stream, profile profileConsumer) consumer {
	return consumer{
		stream:  stream,
		profile: profile,
	}
}

func (c consumer) Start() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go c.stream.SubscribeToEvent("search.profile.created", events.ProfileCreated{}, c.profile.ProfileCreated)
	go c.stream.SubscribeToEvent("search.profile.updated", events.ProfileUpdated{}, c.profile.ProfileUpdate)
	go c.stream.SubscribeToEvent("search.user.deleted", events.UserDeleted{}, c.profile.ProfileDeleted)

	wg.Wait()
}
