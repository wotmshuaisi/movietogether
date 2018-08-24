package handler

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/movietogether/config"
	"github.com/wotmshuaisi/movietogether/handler/httphandler"
	"github.com/wotmshuaisi/movietogether/handler/rtmphandler"
	"github.com/wotmshuaisi/movietogether/model"
)

// RegisterHTTPHandlers ...
func RegisterHTTPHandlers(log *logrus.Logger, channel *pubsub.Queue, upgrader *websocket.Upgrader, msgqueue chan []byte, m model.MovietogetherDBInterface) *mux.Router {
	// handler part
	handlers := httphandler.HTTPHandlers{
		Log:       log,
		Channel:   channel,
		Upgrader:  upgrader,
		MsgQueue:  msgqueue,
		Model:     m,
		WsClients: map[string]*websocket.Conn{},
	}

	// router part
	router := mux.NewRouter()

	router.Use(handlers.LoggingMiddleware)
	router.Use(handlers.UserCheckMiddleware)
	// chat
	router.HandleFunc("/mt", httphandler.Index)
	router.HandleFunc("/mt/register", handlers.Register).Methods("POST")
	router.HandleFunc("/mt/checkuser", handlers.CheckUser).Methods("GET")

	router.HandleFunc("/mt/history", handlers.History).Methods("GET")
	router.HandleFunc("/mt/chat", handlers.Chat)
	// moive
	router.HandleFunc(config.FLVURL, handlers.Movie)

	// broadcast message
	go handlers.Broadcast()

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
	// rtmpServer.HandlePlay = handlers.Play
	return &rtmpServer
}
