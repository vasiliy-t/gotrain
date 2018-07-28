package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis"
)

// QueryHandler is a net/http handler, returns stored objects with their headers
type QueryHandler struct {
	*redis.Client
}

func (c *QueryHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	b, err := c.Get(key).Bytes()
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
	var res ObjectContainer

	err = json.Unmarshal(b, &res)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	for _, m := range res.Meta {
		rw.Header().Add(m.Name, m.Value)
	}

	rw.Write(res.Value)
}
