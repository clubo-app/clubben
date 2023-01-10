package stream

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Stream struct {
	nc *nats.Conn
	js nats.JetStreamContext
}

func new(nc *nats.Conn, js nats.JetStreamContext) Stream {
	return Stream{nc: nc, js: js}
}

func (s Stream) Close() {
	s.nc.Close()
}

func Connect(cluster string, opts []nats.Option) (Stream, error) {
	opts = setupConnOptions(opts)

	nc, err := nats.Connect(cluster, opts...)
	if err != nil {
		return Stream{}, err
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Println("Error creating JestStream Context: ", err)
	}

	log.Println("Connected to Nats Server at ", nc.ConnectedUrl())

	return new(nc, js), nil
}

func (s Stream) PublishEvent(event proto.Message) (*nats.PubAck, error) {
	msg, err := proto.Marshal(event)
	if err != nil {
		return nil, err
	}

	e := eventFromProtobufMessage(event)

	s.makeSureStreamExists(e.streamName, e.subject)
	return s.js.Publish(e.subject, msg)
}

// PullSubscribe crates a Pull based Consumer.
func (s Stream) PullSubscribe(event any, queue string, opts ...nats.SubOpt) (*nats.Subscription, error) {
	e := eventFromProtobufMessage(event)

	return s.js.PullSubscribe(e.subject, queue, append([]nats.SubOpt{nats.BindStream(e.streamName)}, opts...)...)
}

// PushSubscribe creates a push-based Consumer.
// When specifying a queue the messages will be distributed when using multiple Consumer on the same Subject.
func (s Stream) PushSubscribe(event any, queue string, opts ...nats.SubOpt) (*nats.Subscription, error) {
	e := eventFromProtobufMessage(event)

	opt := append([]nats.SubOpt{nats.BindStream(e.streamName)}, opts...)

	if queue != "" {
		return s.js.QueueSubscribeSync(e.subject, queue, opt...)
	}
	return s.js.SubscribeSync(e.subject, opt...)
}

func (s Stream) makeSureStreamExists(name string, subject string) {
	stream, _ := s.js.StreamInfo(name)

	if stream == nil {
		log.Printf("Creating stream: %s\n", name)

		_, err := s.js.AddStream(&nats.StreamConfig{
			Name:     name,
			Subjects: []string{name + ".*"},
		})
		if err != nil {
			log.Fatalf("Err creating stream: %v\n", err)
		}
	}
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
