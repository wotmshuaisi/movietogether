package main

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/wotmshuaisi/movietogether/test/handler"
	"net/http"
)

func main() {
	// set address string getway
	listenAddr := flag.String("addr", ":8080", "service address")
	// get args
	flag.Parse()
	// listen websocket
	initWebService(*listenAddr)
}

func initWebService(address string) {
	// router
	router := chi.NewRouter()
	router.Get("/", handler.IndexHandler)
	router.HandleFunc("/ws", handler.WebsocketHandler)
	http.ListenAndServe(address, router)
}
