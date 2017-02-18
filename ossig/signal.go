// demo app to test os signals handling in docker container
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	schan := make(chan os.Signal)
	// docker stop sends SIGTERM, which we can catch and handle
	// docker kill sends SIGKILL, which we will never know
	// it allows to specify signal  docker kill -s TERM container_id
	// docker rm -f sends SIGKILL, does not allows to send any other
	signal.Notify(schan, syscall.SIGTERM)
	fmt.Println(<-schan)
	fmt.Println("Caught interrupt, exiting")
}
