package structs

import (
	"fmt"
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
)

type Command struct {
	run  func() error
	done chan error
}

type EventLoop struct {
	commands  chan *Command
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
		store:     store,
		transport: transport,
	}

	go loop.Run()

	return loop
}

func (e *EventLoop) dispatch(message *protocol.Message) error {
	fmt.Println("Recebido: ", message)

	proxy, ok := e.store.Get(message.ObjectID)

	if !ok {
		return fmt.Errorf("Proxy with id %d not found!", message.ObjectID)
	}

	proxy.Handle(message)

	return nil
}

func (e *EventLoop) receiveLoop() error {
	for {
		data, err := e.transport.Receive()

		if err != nil {
			return err
		}

		e.Do(func() error {
			return e.dispatch(protocol.Decode(data))
		})
	}
}

func (e *EventLoop) Do(run func() error) error {
	done := make(chan error, 1)

	e.commands <- &Command{
		run:  run,
		done: done,
	}

	return <-done
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
