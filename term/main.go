// Example of registering and firing os signal handlers
package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type TermHook func()

var termHandlers []TermHook

func main() {
	AddOnTermHook(func() {
		log.Println("some resource shutdown func")
		<-time.After(time.Duration(1 * time.Second))
	})
	AddOnTermHook(func() {
		log.Println("another resource shutdown func")
		<-time.After(time.Duration(2 * time.Second))
	})

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)
	<-sigchan

	OnTerm()
}

func AddOnTermHook(h TermHook) {
	termHandlers = append(termHandlers, h)
}

func OnTerm() {
	timer := time.NewTimer(time.Duration(15 * time.Second))
	wg := sync.WaitGroup{}

	for _, f := range termHandlers {
		wg.Add(1)
		go func(th TermHook) {
			defer wg.Done()
			th()
		}(f)
	}

	doneChan := make(chan struct{})

	go func(group *sync.WaitGroup) {
		group.Wait()
		doneChan <- struct{}{}
	}(&wg)

	select {
	case <-timer.C:
		log.Fatalf("shutdonw timeout")
	case <-doneChan:
		log.Printf("shutdown ok")
	}
}
