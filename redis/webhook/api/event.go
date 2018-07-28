package api

import (
	"net/http"

	"github.com/go-redis/redis"
)

// NewEventHandler creates EventHandler instances
func NewEventHandler(r *redis.Client) http.Handler {
	return &EventHandler{}
}

// EventHandler is a HTTP request handler, sends incoming events to redis streams
type EventHandler struct{}

func (h *EventHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusAccepted)
	rw.Write([]byte("accepted"))
}
