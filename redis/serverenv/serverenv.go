package serverenv

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type ShutdownFunc func()

var registeredShutdownFuncs []ShutdownFunc

func RegisterShutdownFunc(h ShutdownFunc) {
	registeredShutdownFuncs = append(registeredShutdownFuncs, h)
}

func Shutdown(timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	wg := sync.WaitGroup{}

	for _, f := range registeredShutdownFuncs {
		wg.Add(1)
		go func(th ShutdownFunc) {
			defer wg.Done()
			th()
		}(f)
	}

	doneChan := make(chan struct{})

	go func() {
		wg.Wait()
		doneChan <- struct{}{}
	}()

	select {
	case <-timer.C:
		return errors.New("shutdown timeout")
	case <-doneChan:
		return nil
	}
}

func LoopUntilShutdown(shutdownTimeout time.Duration) error {
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)
	<-sigchan

	return Shutdown(shutdownTimeout)
}

