// Go 1.8 introduces http.Server.Shutdown() method
// for handling graceful server shutdown
// This demo app illustrates basic usage
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(time.Duration(5 * time.Second)):
			fmt.Fprintf(rw, "qwerty")
		}
	})
	s := &http.Server{
		Addr:        "0.0.0.0:9000",
		Handler:     mux,
		ReadTimeout: time.Duration(10 * time.Second),
		IdleTimeout: time.Duration(10 * time.Second),
	}
	// immediately exits when Shutdown is called
	go s.ListenAndServe()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)
	log.Println(<-sigchan)
	s.Shutdown(context.Background())
}
