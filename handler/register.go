package handler

import (
	"github.com/gorilla/mux"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/movietogether/config"
	"github.com/wotmshuaisi/movietogether/handler/httphandler"
	"github.com/wotmshuaisi/movietogether/handler/rtmphandler"
)

// RegisterHTTPHandlers ...
func RegisterHTTPHandlers(log *logrus.Logger, channel *pubsub.Queue) *mux.Router {
	// handler part
	handlers := httphandler.HTTPHandlers{
		Log:     log,
		Channel: channel,
	}
	// router part
	router := mux.NewRouter()

	router.Use(handlers.LoggingMiddleware)
	// chat
	router.HandleFunc("/chat", handlers.Chat)
	router.HandleFunc("/msg", handlers.Msg)
	// moive
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
