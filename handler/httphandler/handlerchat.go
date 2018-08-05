package httphandler

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{}
)

// Chat websocket chat interface :  recv && transfer to nats
func (handlers *HTTPHandlers) Chat(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("status", "200")
	w.Write([]byte("no content"))
	return
}
