package consumer

import (
	"context"
	"log"

	"github.com/clubo-app/clubben/libs/stream"
	scyllacdc "github.com/scylladb/scylla-cdc-go"
)

type PartyParticipantsConsumer struct {
	Id        int
	TableName string
	Stream    stream.Stream
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
