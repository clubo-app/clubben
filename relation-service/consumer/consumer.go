package consumer

import (
	"context"
	"log"
	"sync"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/protobuf/events"
	"github.com/clubo-app/clubben/relation-service/service"
)

type consumer struct {
	stream stream.Stream
	fs     service.FriendRelation
	ps     service.FavoriteParty
}

func New(stream stream.Stream, fs service.FriendRelation, ps service.FavoriteParty) consumer {
	return consumer{stream: stream, fs: fs, ps: ps}
}

func (c consumer) Start() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	go c.stream.PushSubscribe("relation.friend.created.count", events.FriendCreated{}, c.FriendCreated)
	go c.stream.PushSubscribe("relation.friend.removed.count", events.FriendRemoved{}, c.FriendRemoved)

	wg.Wait()
}

func (c consumer) FriendCreated(e *events.FriendCreated) {
	err := c.fs.IncreaseFriendCount(context.Background(), e.UserId)

	if err != nil {
		log.Println("Error increasing Count: ", err)
	}
}

func (c consumer) FriendRemoved(e *events.FriendCreated) {
	err := c.fs.DecreaseFriendCount(context.Background(), e.UserId)

	if err != nil {
		log.Println("Error decreasing Count: ", err)
	}
}
