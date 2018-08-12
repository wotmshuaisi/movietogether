package httphandler

import (
	"github.com/nareix/joy4/av/pubsub"
	"github.com/sirupsen/logrus"
)

// HTTPHandlers ...
type HTTPHandlers struct {
	Log     *logrus.Logger
	Channel *pubsub.Queue
}
