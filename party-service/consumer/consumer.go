package consumer

import (
	"sync"

	"github.com/clubo-app/clubben/libs/stream"
)

type consumer struct {
	stream stream.Stream
}

func New(stream stream.Stream) consumer {
	return consumer{stream: stream}
}

func (c consumer) Start() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	wg.Wait()
}
