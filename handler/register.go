package handler

import (
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"github.com/wotmshuaisi/movietogether/handler/httphandler"
)

// RegisterHTTPHandlers ...
func RegisterHTTPHandlers(natsCon *nats.Conn) *mux.Router {
	// handler part
	handlers := httphandler.HTTPHandlers{
		NatsClient: natsCon,
	}
	// router part
	router := mux.NewRouter()

	router.HandleFunc("/chat", handlers.Chat)
	router.HandleFunc("/movie", handlers.Movie)
	router.HandleFunc("/publish", handlers.Publish)

	return router
}
