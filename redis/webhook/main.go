package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliy-t/gotrain/redis/serverenv"
	"github.com/vasiliy-t/gotrain/redis/webhook/api"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Redis: failed to ping, %s", err)
	}

	log.Infof("Redis: connected %s", pong)

	eventHandler := api.NewEventHandler(client)
	subscribeHandler := api.NewSubscribeHandler(client)

	m := mux.NewRouter()
	m.Handle("/api/event/{stream}", eventHandler)
	m.Handle("/api/subscribe/{stream}", subscribeHandler)

	httpServer := &http.Server{
		Addr:         "0.0.0.0:9000",
		Handler:      m,
		ReadTimeout:  time.Duration(1 * time.Second),
		WriteTimeout: time.Duration(1 * time.Second),
	}

	log.Infof("Ready to count")

	go httpServer.ListenAndServe()

	serverenv.RegisterShutdownFunc(func() {
		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Errorf("HTTP server: failed to shutdown %s", err)
		}
		log.Info("HTTP server: shutdown successfully")
	})

	serverenv.LoopUntilShutdown(5 * time.Second)
}
