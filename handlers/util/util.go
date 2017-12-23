package util

import (
	"io"
	"net/http"
	"time"
)

// Run the provided function in a goroutine.
func Run(fn func()) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		go fn()
		w.WriteHeader(http.StatusOK)
		r.Body.Close()
	}
}

// Ping always returns HTTP 200.
func Ping() http.HandlerFunc {
	// TODO(negz): Check kubehook health?
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		r.Body.Close()
	}
}

// Content serves the supplied content at the supplied path.
func Content(c io.ReadSeeker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, "", time.Unix(0, 0), c)
		r.Body.Close()
	}
}
