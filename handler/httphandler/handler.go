package httphandler

import (
	"io"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/movietogether/model"
)

// HTTPHandlers ...
type HTTPHandlers struct {
	Log      *logrus.Logger
	Channel  *pubsub.Queue
	Upgrader *websocket.Upgrader
	MsgQueue chan []byte
	Model    model.MovietogetherDBInterface
}

// others
type writeFlusher struct {
	httpflusher http.Flusher
	io.Writer
}

func (wf writeFlusher) Flush() error {
	wf.httpflusher.Flush()
	return nil
}

// context key
type ckey int

const (
	namekey ckey = iota
)
