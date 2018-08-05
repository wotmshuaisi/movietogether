package handler

import (
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/movietogether/handler/httphandler"
)

// RegisterHTTPHandlers ...
func RegisterHTTPHandlers(natsCon *nats.Conn, log *logrus.Logger) *mux.Router {
	// handler part
	handlers := httphandler.HTTPHandlers{
		NatsClient: natsCon,
		Log:        log,
	}
	// router part
	router := mux.NewRouter()

	router.Use(handlers.LoggingMiddleware)
	router.HandleFunc("/chat", handlers.Chat)
	router.HandleFunc("/movie", handlers.Movie)
	router.HandleFunc("/publish", handlers.Publish)

	return router
}
