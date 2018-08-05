package httphandler

import (
	"github.com/nats-io/go-nats"
)

// HTTPHandlers ...
type HTTPHandlers struct {
	NatsClient *nats.Conn
}
