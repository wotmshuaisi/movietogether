package httphandler

import (
	"context"
	"errors"
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

// UserCheckMiddleware ...
func (handlers *HTTPHandlers) UserCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			jsonerrorreturn(err, "400", w)
		}
		c, err := handlers.Model.ClientGet(cookie.Value)
		if err != nil {
			jsonerrorreturn(err, "200", w)
		}
		if c == nil || len(c.Token) <= 0 {
			jsonerrorreturn(errors.New("not allowed"), "403", w)
		}
		ctx := context.WithValue(r.Context(), namekey, c.Name)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
