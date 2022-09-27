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
	stream    stream.Stream
	partyRepo *repository.PartyRepository
}

func New(stream stream.Stream, repo *repository.PartyRepository) consumer {
	return consumer{stream: stream, partyRepo: repo}
}

func (c consumer) Start() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	favoritedSub, err := c.stream.PullSubscribe(events.PartyFavorited{})
	if err != nil {
		log.Fatalln("Failed connect to PartyFavorited Event: ", err)
	}
	go func() {
		defer wg.Done()
		c.partyFavorited(favoritedSub)
	}()

	unfavoritedSub, err := c.stream.PullSubscribe(events.PartyUnfavorited{})
	if err != nil {
		log.Fatalln("Failed connect to PartyUnfavorited Event: ", err)
	}
	go func() {
		defer wg.Done()
		c.partyUnfavorited(unfavoritedSub)
	}()

	wg.Wait()
}

func (c consumer) partyFavorited(sub *nats.Subscription) {
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

func (c consumer) partyUnfavorited(sub *nats.Subscription) {
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
