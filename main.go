// Package main ...
package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/podanypepa/graceful-shutdown/pkg/rest"
	"github.com/rs/zerolog/log"
)

const (
	port = 8765
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Info().Msg("exit signal recieved")
		cancel()
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		rest.Start(ctx, port)
	}()

	wg.Wait()
}
