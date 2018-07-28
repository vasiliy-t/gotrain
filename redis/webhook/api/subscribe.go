package api

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// subscribeHandler registers subscribers to webhooks by paths
type subscribeHandler struct {
	redis *redis.Client
}

func (h *subscribeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stream := vars["stream"]

	msgs, err := h.redis.XRange(stream, "-", "+").Result()

	if err != nil {
		log.Errorf("Subscriber: failed to range stream %s", err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Info("Subscriber: messages %+v", msgs)

	for _, msg := range msgs {
		log.Infof("Subscriber: from stream %+v", msg.Values)
	}

	rw.WriteHeader(http.StatusAccepted)
	rw.Write([]byte("subscribed"))
}

// NewSubscribeHandler creates event subscriber instances
func NewSubscribeHandler(r *redis.Client) http.Handler {
	return &subscribeHandler{redis: r}
}
