package main

import (
	"context"
	"log"
	"time"

	svr "github.com/noncepad/go-tutorial-1/server"
)

func main() {
	log.Print("hello world!")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server, err := svr.Create(ctx, &svr.Configuration{Version: 1})
	if err != nil {
		panic("server failed to start")
	}
	log.Print("server started")
	signalC := server.CloseSignal()
	err = <-signalC
	if err != nil {
		panic(err)
	}
	log.Print("server exited")
}
