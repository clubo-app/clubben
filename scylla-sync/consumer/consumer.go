package consumer

import (
	"context"
	"log"

	scyllacdc "github.com/scylladb/scylla-cdc-go"
)

type Consumer struct {
	Id                    int
	TableName             string
	Reporter              *scyllacdc.PeriodicProgressReporter
	FriendConsumer        *FriendRelationConsumer
	FavoritePartyConsumer *FavoritePartiesConsumer
	ParticipantsConsumer  *PartyParticipantsConsumer
}

func (c *Consumer) End() error {
	_ = c.Reporter.SaveAndStop(context.Background())
	return nil
}

func (c *Consumer) Consume(ctx context.Context, ch scyllacdc.Change) error {
	log.Println(ch)
	if c.TableName == FRIEND_RELATIONS_TABLE {
		c.FriendConsumer.Consume(ctx, ch)
	} else if c.TableName == FAVORITE_PARTIES_TABLE {
		c.FavoritePartyConsumer.Consume(ctx, ch)
	} else if c.TableName == PARTY_PARTICIPANTS_TABLE {
		c.ParticipantsConsumer.Consume(ctx, ch)
	}
	return nil
}
