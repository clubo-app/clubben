package consumer

import (
	"context"
	"log"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/protobuf/events"
	scyllacdc "github.com/scylladb/scylla-cdc-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PartyParticipantsConsumer struct {
	Id        int
	TableName string
	Stream    *stream.Stream
	Reporter  *scyllacdc.PeriodicProgressReporter
}

func (c *PartyParticipantsConsumer) End() error {
	_ = c.Reporter.SaveAndStop(context.Background())
	return nil
}

func (c *PartyParticipantsConsumer) Consume(ctx context.Context, ch scyllacdc.Change) error {
	for _, change := range ch.Delta {
		switch change.GetOperation() {
		case scyllacdc.Update:
			_ = c.processUpdateOrInsert(ctx, change)
		case scyllacdc.Insert:
			_ = c.processUpdateOrInsert(ctx, change)
		case scyllacdc.RowDelete:
			_ = c.processDelete(ctx, change)
		default:
			log.Println("unsupported operation: " + change.GetOperation().String())
		}
	}
	c.Reporter.Update(ch.Time)
	return nil
}

func (c *PartyParticipantsConsumer) processUpdateOrInsert(ctx context.Context, change *scyllacdc.ChangeRow) error {
	log.Println("Processing Participation Update or Insert")
	userId, _ := change.GetValue("user_id")
	partyId, _ := change.GetValue("party_id")
	requested, _ := change.GetValue("requested")
	requestedAt, _ := change.GetValue("requested_at")
	joinedAt, _ := change.GetValue("joined_at")

	log.Printf("Requested: %v", requested)
	log.Printf("Requested Pointer Value: %v", &requested)

	uId := ParseString(userId)
	pId := ParseString(partyId)
	rAt := ParseTimestamp(requestedAt)
	jAt := ParseTimestamp(joinedAt)

	if requested == true {
		e := events.PartyRequested{
			UserId:      uId,
			PartyId:     pId,
			RequestedAt: rAt,
		}
		err := c.Stream.PublishEvent(&e)
		if err != nil {
			log.Println("Error publishing PartyRequested: ", err)
		}
	} else {
		e := events.PartyJoined{
			UserId:   uId,
			PartyId:  pId,
			JoinedAt: jAt,
		}
		err := c.Stream.PublishEvent(&e)
		if err != nil {
			log.Println("Error publishing PartyJoined: ", err)
		}
	}
	return nil
}

func (c *PartyParticipantsConsumer) processDelete(ctx context.Context, change *scyllacdc.ChangeRow) error {
	log.Println("Processing Participation Delete")
	userId, _ := change.GetValue("user_id")
	partyId, _ := change.GetValue("party_id")

	uId := ParseString(userId)
	pId := ParseString(partyId)

	e := events.PartyLeft{
		UserId:  uId,
		PartyId: pId,
		LeftAt:  timestamppb.Now(),
	}
	err := c.Stream.PublishEvent(&e)
	if err != nil {
		log.Println("Error publishing PartyLeft: ", err)
	}
	return nil
}
