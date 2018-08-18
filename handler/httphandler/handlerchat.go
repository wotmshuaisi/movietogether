package httphandler

import (
	"net/http"
)

// Chat interface :  recv message && broadcast
func (handlers *HTTPHandlers) Chat(w http.ResponseWriter, r *http.Request) {
	// wsCon, err := handlers.Upgrader.Upgrade(w, r, nil)
	return
}

// History get paged data
func (handlers *HTTPHandlers) History(w http.ResponseWriter, r *http.Request) {
	// wsCon, err :=
	return
}
