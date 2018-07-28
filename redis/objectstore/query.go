package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis"
)

type QueryHandler struct {
	*redis.Client
}

func (c *QueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	b, err := c.Get(key).Bytes()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var res ObjectContainer

	err = json.Unmarshal(b, &res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(res.Value)
}
