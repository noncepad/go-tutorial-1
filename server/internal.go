package server

import (
	"context"
	"log"
)

type internal struct {
	ctx              context.Context
	errorC           chan<- error
	closeSignalCList []chan<- error
	config           *Configuration
	// put your protected state in here
}

func loopInternal(ctx context.Context, internalC <-chan func(*internal), config *Configuration) {

	var err error
	errorC := make(chan error, 1)
	doneC := ctx.Done()
	in := new(internal)
	in.ctx = ctx
	in.errorC = errorC
	in.closeSignalCList = make([]chan<- error, 0)

out:
	for {
		select {
		case <-doneC:
			break out
		case err = <-errorC:
			break out
		case req := <-internalC:
			req(in)
		}
	}
	in.finish(err)
}

func (in *internal) finish(err error) {
	log.Print("server exiting")
	log.Print(err)
	for i := 0; i < len(in.closeSignalCList); i++ {
		in.closeSignalCList[i] <- err
	}
}
