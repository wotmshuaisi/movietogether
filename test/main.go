package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/wotmshuaisi/movietogether/test/handler"
)

func main() {
	// set address string getway
	listenAddr := flag.String("addr", ":8080", "service address")
	// get args
	flag.Parse()
	// listen websocket
	initWebService(*listenAddr)

	// // go lang channel example
	// channel := make(chan []byte, 2048)
	// go func() {
	// 	for {
	// 		channel <- []byte("ping")
	// 		time.Sleep(time.Second * 3)
	// 	}
	// }()
	// for {
	// 	print(<-channel)
	// 	time.Sleep(time.Second * 1)
	// }
}

func initWebService(address string) {
	// router
	router := chi.NewRouter()
	router.Get("/", handler.IndexHandler)
	router.HandleFunc("/ws", handler.WebsocketHandler)
	http.ListenAndServe(address, router)
}
