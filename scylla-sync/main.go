package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/scylla-sync/config"
	"github.com/clubo-app/clubben/scylla-sync/consumer"
	"github.com/nats-io/nats.go"
	scyllacdc "github.com/scylladb/scylla-cdc-go"
)

func main() {
	c := config.LoadConfig()

	opts := []nats.Option{nats.Name("Scylla-Sync Service")}
	stream, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.Close()

	sess, err := newCluster(c.CQL_KEYSPACE, c.CQL_HOSTS)
	if err != nil {
		log.Fatalln(err)
	}
	defer sess.Close()

	logger := log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	factory := newFactory(logger, &stream)
	// progressManager, err := scyllacdc.NewTableBackedProgressManager(sess.Session, "cdc_progress", "scylla_sync")
	// if err != nil {
	// log.Println("Error creating progress Manager", err)
	// }

	cfg := &scyllacdc.ReaderConfig{
		Session: sess.Session,
		TableNames: []string{
			"relation." + consumer.FRIEND_RELATIONS_TABLE,
			"relation." + consumer.FAVORITE_PARTIES_TABLE,
			"participants." + consumer.PARTY_PARTICIPANTS_TABLE,
		},
		// ProgressManager:       progressManager,
		ChangeConsumerFactory: factory,
		Logger:                logger,
		Advanced: scyllacdc.AdvancedReaderConfig{
			PostEmptyQueryDelay:    time.Second * 1,
			PostNonEmptyQueryDelay: time.Second * 1,
			PostFailedQueryDelay:   time.Second * 1,
			ConfidenceWindowSize:   time.Second * 1,
			QueryTimeWindowSize:    time.Second * 10,
			ChangeAgeLimit:         time.Second * 10,
		},
	}

	reader, err := scyllacdc.NewReader(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	signalC := make(chan os.Signal)
	go func() {
		<-signalC
		reader.Stop()

		<-signalC
		os.Exit(1)
	}()
	signal.Notify(signalC, os.Interrupt)

	if err := reader.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
