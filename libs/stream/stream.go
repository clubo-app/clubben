package stream

import (
	"log"
	"reflect"
	"strings"
	"time"
	"unicode"

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

	sub := EventToSubject(event)

	return s.js.Publish(sub, msg)
}

// PullSubscribe crates a Pull based Consumer
func (s Stream) PullSubscribe(event any, opts ...nats.SubOpt) (*nats.Subscription, error) {
	sub := EventToSubject(event)

	return s.js.PullSubscribe(sub, "pull-consumer", opts...)
}

// PushSubscribe creates a push-based Consumer.
// When specifying a queue the messages will be distributed when using multiple Consumer on the same Subject.
func (s Stream) PushSubscribe(queue string, event any, handler nats.MsgHandler) (*nats.Subscription, error) {
	sub := EventToSubject(event)
	log.Printf("Subject is: %v", sub)

	if queue != "" {
		return s.js.QueueSubscribe(sub, queue, handler)
	}
	return s.js.Subscribe(sub, handler)
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

func EventToSubject(event any) string {
	t := reflect.TypeOf(event)

	str := strings.ReplaceAll(t.String(), "*", "")

	lower := camelcaseStringToDotString(str)

	s := strings.Split(lower, ".")

	// if type is events.ImportantType, remove events prefix from string
	if len(s) > 1 {
		return camelcaseStringToDotString(s[len(s)-1])
	}

	return camelcaseStringToDotString(t.String())
}

func camelcaseStringToDotString(camelcase string) string {
	var b strings.Builder

	for i, c := range camelcase {
		if unicode.IsUpper(c) {
			if i != 0 {
				b.WriteString(".")
			}
			b.WriteRune(unicode.ToLower(c))
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}
