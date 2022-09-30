package consumer

import (
	"context"
	"log"
	"sync"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/protobuf/events"
	"github.com/clubo-app/clubben/relation-service/service"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type consumer struct {
	stream *stream.Stream
	fs     service.FriendRelation
	ps     service.FavoriteParty
}

func New(stream *stream.Stream, fs service.FriendRelation, ps service.FavoriteParty) consumer {
	return consumer{stream: stream, fs: fs, ps: ps}
}

func (c *consumer) Start() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		c.stream.PushSubscribe("relation.friend.created.count", events.FriendCreated{}, c.friendCreated)
	}()
	go c.stream.PushSubscribe("relation.friend.removed.count", events.FriendRemoved{}, c.friendRemoved)

	wg.Wait()
}

func (c *consumer) friendCreated(msg *nats.Msg) {
	e := events.FriendCreated{}
	err := proto.Unmarshal(msg.Data, &e)
	if err != nil {
		log.Println("Error unmarshaling FriendCreated: ", err)
	}

	err = c.fs.IncreaseFriendCount(context.Background(), e.UserId)
	if err != nil {
		log.Println("Error increasing Count: ", err)
	}
}

func (c *consumer) friendRemoved(msg *nats.Msg) {
	e := events.FriendRemoved{}
	err := proto.Unmarshal(msg.Data, &e)
	if err != nil {
		log.Println("Error unmarshaling FriendRemoved: ", err)
	}

	err = c.fs.DecreaseFriendCount(context.Background(), e.UserId)
	if err != nil {
		log.Println("Error decreasing Count: ", err)
	}
}
