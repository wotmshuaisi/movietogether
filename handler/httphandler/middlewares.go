package httphandler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// LoggingMiddleware ...
func (handlers *HTTPHandlers) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
		// Do stuff here
		handlers.Log.WithFields(logrus.Fields{
			"method":  r.Method,
			"version": r.Proto,
			"host":    r.RemoteAddr,
			"uri":     r.RequestURI,
			"ip":      r.RemoteAddr,
			"status":  w.Header().Get("status")}).Info()
	})
}
