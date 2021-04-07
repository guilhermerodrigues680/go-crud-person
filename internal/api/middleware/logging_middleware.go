package middleware

import (
	"log"
	"net/http"
)

type LoggingMiddleware struct{}

func NewLoggingMiddleware() LoggingMiddleware {
	return LoggingMiddleware{}
}

func (LoggingMiddleware) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s from IP %s\n", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
