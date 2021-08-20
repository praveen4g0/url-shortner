package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// repo <- service -> serializer  -> http

func main() {
	errs := make(chan error, 1)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
	os.Exit(0)
}
