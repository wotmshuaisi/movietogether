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
	router.HandleFunc(config.PREFIXURI+"/", httphandler.Index)
	router.HandleFunc(config.PREFIXURI+"/register", handlers.Register).Methods("POST")
	router.HandleFunc(config.PREFIXURI+"/checkuser", handlers.CheckUser).Methods("GET")

	router.HandleFunc(config.PREFIXURI+"/history", handlers.History).Methods("GET")
	router.HandleFunc(config.PREFIXURI+"/chat", handlers.Chat)
	// moive
	router.HandleFunc(config.PREFIXURI+"/movie", handlers.Movie)

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
	return &rtmpServer
}
