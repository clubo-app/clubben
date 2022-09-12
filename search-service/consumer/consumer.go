package consumer

import (
	"sync"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/protobuf/events"
)

type consumer struct {
	stream  *stream.Stream
	profile profileConsumer
	party   partyConsumer
}

func NewConsumer(stream *stream.Stream, profile profileConsumer, party partyConsumer) consumer {
	return consumer{
		stream:  stream,
		profile: profile,
		party:   party,
	}
}

func (c consumer) Start() {
	wg := sync.WaitGroup{}
	wg.Add(9)

	go c.stream.SubscribeToEvent("search.profile.created", events.ProfileCreated{}, c.profile.ProfileCreated)
	go c.stream.SubscribeToEvent("search.profile.updated", events.ProfileUpdated{}, c.profile.ProfileUpdate)
	go c.stream.SubscribeToEvent("search.user.deleted", events.UserDeleted{}, c.profile.ProfileDeleted)
	go c.stream.SubscribeToEvent("search.friend.created", events.FriendCreated{}, c.profile.FriendCreated)
	go c.stream.SubscribeToEvent("search.friend.removed", events.FriendRemoved{}, c.profile.FriendRemoved)

	go c.stream.SubscribeToEvent("search.party.created", events.PartyCreated{}, c.party.PartyCreated)
	go c.stream.SubscribeToEvent("search.party.updated", events.PartyUpdated{}, c.party.PartyUpdated)
	go c.stream.SubscribeToEvent("search.party.favorited", events.PartyFavorited{}, c.party.PartyFavorited)
	go c.stream.SubscribeToEvent("search.party.unfavorited", events.PartyUnfavorited{}, c.party.PartyUnfavorited)

	wg.Wait()
}
