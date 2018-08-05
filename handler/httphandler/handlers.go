package httphandler

import (
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
)

// HTTPHandlers ...
type HTTPHandlers struct {
	NatsClient *nats.Conn
	Log        *logrus.Logger
}
