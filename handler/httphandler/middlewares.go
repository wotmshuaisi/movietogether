package httphandler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// LoggingMiddleware ...
func (handlers *HTTPHandlers) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
		// Do stuff here
		handlers.Log.WithFields(logrus.Fields{
			"method": r.Method,
			"scheme": r.URL.Scheme,
			"host":   r.Host,
			"uri":    r.RequestURI,
			"ip":     r.RemoteAddr,
			"status": w.Header().Get("status")}).Infoln()
	})
}
