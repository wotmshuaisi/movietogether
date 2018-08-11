package httphandler

import (
	"io"
	"net/http"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/format/flv"
)

// Movie return flv streaming
func (handlers *HTTPHandlers) Movie(w http.ResponseWriter, r *http.Request) {
	if handlers.Channel == nil {
		w.WriteHeader(204)
		return
	}
	w.Header().Set("Content-Type", "video/x-flv")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	flusher := w.(http.Flusher)
	flusher.Flush()

	muxer := flv.NewMuxerWriteFlusher(writeFlusher{
		httpflusher: flusher,
		Writer:      w,
	})

	cursor := handlers.Channel.Latest()

	avutil.CopyFile(muxer, cursor)
	return
}

type writeFlusher struct {
	httpflusher http.Flusher
	io.Writer
}

func (wf writeFlusher) Flush() error {
	wf.httpflusher.Flush()
	return nil
}
