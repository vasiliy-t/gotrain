package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"encoding/json"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const META_PREFIX = "X-"

type MutationHandler struct {
	*redis.Client
}

func (c *MutationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Infof("HEADERS: %+v", extractMeta(r.Header))

	container := ObjectContainer{}
	container.Meta = extractMeta(r.Header)
	container.Value = value

	toStore, err := json.Marshal(container)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	c.Set(key, toStore, 0)
	rw.Write([]byte("hello"))
}

func extractMeta(headers http.Header) []Meta {
	res := []Meta{}
	for name := range headers {
		if strings.HasPrefix(name, META_PREFIX) {
			res = append(res, Meta{Name: strings.TrimPrefix(name, META_PREFIX), Value: headers.Get(name)})
		}
	}

	return res
}
