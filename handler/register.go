package handler

import (
	"github.com/gorilla/mux"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/movietogether/config"
	"github.com/wotmshuaisi/movietogether/handler/httphandler"
	"github.com/wotmshuaisi/movietogether/handler/rtmphandler"
)

// RegisterHTTPHandlers ...
func RegisterHTTPHandlers(natsCon *nats.Conn, log *logrus.Logger, channel *pubsub.Queue) *mux.Router {
	// handler part
	handlers := httphandler.HTTPHandlers{
		NatsClient: natsCon,
		Log:        log,
		Channel:    channel,
	}
	// router part
	router := mux.NewRouter()

	router.Use(handlers.LoggingMiddleware)
	router.HandleFunc("/chat", handlers.Chat)
	router.HandleFunc(config.FLVURL, handlers.Movie)

	return router
}

// RegisterRTMPHandlers ...
func RegisterRTMPHandlers(channel *pubsub.Queue) *rtmp.Server {
	rtmpServer := rtmp.Server{}
	handlers := rtmphandler.RTMPHandler{
		Channel: channel,
	}
	// handle func part
	rtmpServer.HandlePublish = handlers.Publish
	rtmpServer.HandlePlay = handlers.Play
	return &rtmpServer
}
