package consumer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/party-service/internal/repository"
	"github.com/clubo-app/clubben/protobuf/events"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Consumer struct {
	stream    *stream.Stream
	partyRepo *repository.PartyRepository
}

func New(stream *stream.Stream, repo *repository.PartyRepository) Consumer {
	return Consumer{stream: stream, partyRepo: repo}
}

func (c *Consumer) Start() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()

		sub, err := c.stream.PullSubscribe(events.PartyFavorited{}, "party-service")
		if err != nil {
			log.Println("Failed to subscribe to PartyFavorited Event: ", err)
			return
		}
		c.partyFavorited(sub)
	}()

	go func() {
		defer wg.Done()

		sub, err := c.stream.PullSubscribe(events.PartyUnfavorited{}, "party-service")
		if err != nil {
			log.Println("Failed to subscribe to PartyUnfavorited Event: ", err)
			return
		}
		c.partyUnfavorited(sub)
	}()

	go func() {
		defer wg.Done()

		sub, err := c.stream.PushSubscribe(
			events.PartyJoined{},
			"party-service.party.joined.count",
			"party-service",
		)
		if err != nil {
			log.Println("Failed to subscribe to PartyJoined Event: ", err)
			return
		}

		c.partyJoined(sub)
	}()

	go func() {
		defer wg.Done()

		sub, err := c.stream.PushSubscribe(
			events.PartyLeft{},
			"party-service.party.left.count",
			"party-service",
		)
		if err != nil {
			log.Println("Failed to subscribe to PartyLeft Event: ", err)
			return
		}

		c.partyLeft(sub)
	}()

	wg.Wait()

	log.Println("All Consumers unexpectedly stopped")
}

func (c *Consumer) partyFavorited(sub *nats.Subscription) {
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
		batch.Queue("UPDATE parties SET favorite_count = favorite_count + $1 WHERE id = $2", count, partyId)
	}

	b := c.partyRepo.Pool.SendBatch(context.Background(), batch)
	bRes, err := b.Exec()
	if err != nil {
		log.Println("Error executing Batch for increasing Party Favorite Count: ", err)
		return
	}
	log.Printf("%v Rows were effected by Increate Party Favorite Count in Batch", bRes.RowsAffected())
}

func (c *Consumer) partyUnfavorited(sub *nats.Subscription) {
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
		batch.Queue("UPDATE parties SET favorite_count = favorite_count - $1 WHERE id = $2", count, partyId)
	}

	b := c.partyRepo.Pool.SendBatch(context.Background(), batch)
	bRes, err := b.Exec()
	if err != nil {
		log.Println("Error executing Batch for decreasing Party Favorite Count: ", err)
		return
	}
	log.Printf("%v Rows were effected by Decreasing Party Favorite Count in Batch", bRes.RowsAffected())
}

func (c *Consumer) partyJoined(sub *nats.Subscription) {
	msg, err := sub.NextMsg(1 * time.Second)
	if err != nil {
		log.Print("PartyJoined nextMsg: ")
	}

	e := &events.PartyJoined{}
	err = proto.Unmarshal(msg.Data, e)
	if err != nil {
		log.Println("PartyJoined Unmarshal: ", err)
	}

	if e.PartyId == "" {
		log.Printf("PartyJoined invalid event: %+v", &e)
		return
	}
	c.partyRepo.IncreaseParticipantsCount(context.Background(), repository.IncreaseParticipantsCountParams{
		ParticipantsCount: 1,
		ID:                e.PartyId,
	})

	msg.Ack()
}

func (c *Consumer) partyLeft(sub *nats.Subscription) {
	msg, err := sub.NextMsg(1 * time.Second)
	if err != nil {
		log.Print("PartyLeft nextMsg: ")
	}

	e := &events.PartyLeft{}
	err = proto.Unmarshal(msg.Data, e)
	if err != nil {
		log.Println("PartyLeft Unmarshal: ", err)
	}

	if e.PartyId == "" {
		log.Printf("PartyLeft invalid event: %+v", &e)
		return
	}
	c.partyRepo.DecreaseParticipantsCount(context.Background(), repository.DecreaseParticipantsCountParams{
		ParticipantsCount: 1,
		ID:                e.PartyId,
	})

	msg.Ack()
}
