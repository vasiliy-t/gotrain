package api

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// NewEventHandler creates EventHandler instances
func NewEventHandler(r *redis.Client) http.Handler {
	return &EventHandler{
		redis: r,
	}
}

// EventHandler is a HTTP request handler, sends incoming events to redis streams
type EventHandler struct {
	redis *redis.Client
}

func (h *EventHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stream := vars["stream"]

	m := make(map[string]interface{})
	m["key1"] = "value1"

	if err := h.redis.XAdd(stream, "*", m).Err(); err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
	rw.Write([]byte("accepted"))
}
