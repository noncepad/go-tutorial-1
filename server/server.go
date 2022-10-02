package server

import "context"

type Server struct {
	ctx       context.Context
	internalC chan<- func(*internal)
}

type Configuration struct {
	Version int
}

func Create(ctx context.Context, config *Configuration) (Server, error) {
	internalC := make(chan func(*internal), 10)

	go loopInternal(ctx, internalC, config)
	e1 := Server{
		ctx:       ctx,
		internalC: internalC,
	}

	return e1, nil
}

// this function tells me when the server object has exited/died
func (e1 Server) CloseSignal() <-chan error {
	signalC := make(chan error, 1)
	e1.internalC <- func(in *internal) {
		in.closeSignalCList = append(in.closeSignalCList, signalC)
	}
	return signalC
}
