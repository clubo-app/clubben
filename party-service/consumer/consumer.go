package consumer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/party-service/repository"
	"github.com/clubo-app/clubben/protobuf/events"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type consumer struct {
	stream    *stream.Stream
	partyRepo *repository.PartyRepository
}

func New(stream *stream.Stream, repo *repository.PartyRepository) consumer {
	return consumer{stream: stream, partyRepo: repo}
}

func (c *consumer) Start() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		sub, err := c.stream.PullSubscribe(events.PartyFavorited{})
		if err != nil {
			log.Fatalln("Failed to subscribe to PartyFavorited Event: ", err)
		}
		defer wg.Done()
		c.partyFavorited(sub)
	}()

	go func() {
		sub, err := c.stream.PullSubscribe(events.PartyUnfavorited{})
		if err != nil {
			log.Fatalln("Failed to subscribe to PartyUnfavorited Event: ", err)
		}
		defer wg.Done()
		c.partyUnfavorited(sub)
	}()

	go func() {
		defer wg.Done()
		_, err := c.stream.PushSubscribe("party-service.participants.joined.count", events.PartyJoined{}, c.partyJoined)
		if err != nil {
			log.Fatalln("Failed to subscribe to PartyJoined Event: ", err)
		}
	}()

	go func() {
		defer wg.Done()
		_, err := c.stream.PushSubscribe("party-service.participants.left.count", events.PartyJoined{}, c.partyLeft)
		if err != nil {
			log.Fatalln("Failed to subscribe to PartyLeft Event: ", err)
		}
	}()

	wg.Wait()

	log.Println("All Consumers unexpectedly stopped")
}

func (c *consumer) partyFavorited(sub *nats.Subscription) {
	msgs, err := sub.Fetch(100, nats.MaxWait(time.Second*30))
	if err != nil {
		log.Println("Error fetching PartyFavorited Events: ", err)
	}

	batch := &pgx.Batch{}

	aggregatedCount := make(map[string]int, len(msgs))

	// here we aggregate the counts for every party.
	// we could theoretically save a few DB ops by adding counts for the same Party instead of sending them seperataly.
	for _, msg := range msgs {
		event := &events.PartyFavorited{}
		err := proto.Unmarshal(msg.Data, event)
		if err != nil {
			log.Println(err)
			continue
		}

		aggregatedCount[event.PartyId] += 1
		msg.Ack()
	}

	for partyId, count := range aggregatedCount {
		batch.Queue("UPDATE parties SET participants_count = participants_count + $1 WHERE id = $2", count, partyId)
	}

	b := c.partyRepo.Pool.SendBatch(context.Background(), batch)
	bRes, err := b.Exec()
	if err != nil {
		log.Println("Error executing Batch for increasing Party Favorite Count: ", err)
		return
	}
	log.Printf("%v Rows were effected by Increate Party Favorite Count in Batch", bRes.RowsAffected())
}

func (c *consumer) partyUnfavorited(sub *nats.Subscription) {
	msgs, err := sub.Fetch(100, nats.MaxWait(time.Second*30))
	if err != nil {
		log.Println("Error fetching PartyUnfavorited Events: ", err)
	}

	batch := &pgx.Batch{}

	aggregatedCount := make(map[string]int, len(msgs))

	// here we aggregate the counts for every party.
	// we could theoretically save a few DB ops by adding counts for the same Party instead of sending them seperataly.
	for _, msg := range msgs {
		event := &events.PartyUnfavorited{}
		err := proto.Unmarshal(msg.Data, event)
		if err != nil {
			log.Println(err)
			continue
		}

		aggregatedCount[event.PartyId] += 1
		msg.Ack()
	}

	for partyId, count := range aggregatedCount {
		batch.Queue("UPDATE parties SET participants_count = participants_count - $1 WHERE id = $2", count, partyId)
	}

	b := c.partyRepo.Pool.SendBatch(context.Background(), batch)
	bRes, err := b.Exec()
	if err != nil {
		log.Println("Error executing Batch for decreasing Party Favorite Count: ", err)
		return
	}
	log.Printf("%v Rows were effected by Decreasing Party Favorite Count in Batch", bRes.RowsAffected())
}

func (c *consumer) partyJoined(msg *nats.Msg) {
	e := &events.PartyJoined{}
	err := proto.Unmarshal(msg.Data, e)
	if err != nil {
		log.Println("Error unmarshaling PartyJoined: ", err)
	}

	if e.PartyId == "" {
		log.Printf("Invalid PartyJoined Event %+v", &e)
		return
	}
	c.partyRepo.IncreaseParticipantsCount(context.Background(), repository.IncreaseParticipantsCountParams{
		ParticipantsCount: 1,
		ID:                e.PartyId,
	})
}

func (c *consumer) partyLeft(msg *nats.Msg) {
	e := &events.PartyLeft{}
	err := proto.Unmarshal(msg.Data, e)
	if err != nil {
		log.Println("Error unmarshaling PartyLeft: ", err)
	}

	if e.PartyId == "" {
		log.Printf("Invalid PartyLeft Event %+v", &e)
		return
	}
	c.partyRepo.DecreaseParticipantsCount(context.Background(), repository.DecreaseParticipantsCountParams{
		ParticipantsCount: 1,
		ID:                e.PartyId,
	})
}
