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
	return
}
