package structs

import (
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
)

type Command struct {
	run  func() error
	done chan error
}

type EventLoop struct {
	commands  chan *Command
	dispatch  *Dispatcher
	store     *ProxyStore
	transport *backend.Transport
}

func NewEventLoop() *EventLoop {
	transport := backend.NewTransport(
		protocol.Connect,
		protocol.Frame,
	)

	store := NewProxyStore()

	loop := &EventLoop{
		commands:  make(chan *Command),
		dispatch:  NewDispatcher(store.Get),
		store:     store,
		transport: transport,
	}

	go loop.Run()

	return loop
}

func (e *EventLoop) Do(run func() error) error {
	done := make(chan error, 1)

	e.commands <- &Command{
		run:  run,
		done: done,
	}

	return <-done
}

func (e *EventLoop) receiveLoop() error {
	for {
		data, err := e.transport.Receive()

		if err != nil {
			return err
		}

		e.Do(func() error {
			return e.dispatch.dispatch(protocol.Decode(data))
		})
	}
}

func (e *EventLoop) Send(data []byte) {
	e.Do(func() error {
		return e.transport.Send(data)
	})
}

func (e *EventLoop) Run() {
	go e.receiveLoop()

	for {
		cmd := <-e.commands
		cmd.done <- cmd.run()
	}
}
