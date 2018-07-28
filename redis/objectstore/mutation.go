package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"encoding/json"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type MutationHandler struct {
	*redis.Client
}

func (c *MutationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	container := ObjectContainer{}
	container.Meta = extractMeta(r.Header)
	container.Value = value

	toStore, err := json.Marshal(container)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	c.Set(key, toStore, 0)
	w.Write([]byte("hello"))
}

func extractMeta(headers http.Header) []Meta {
	res := []Meta{}
	for name := range headers {
		log.Infof("Name is: %s", name)
		if strings.HasPrefix("X-", name) {
			res = append(res, Meta{Name: name, Value: headers.Get(name)})
		}
	}

	return res
}
