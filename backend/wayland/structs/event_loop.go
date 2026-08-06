package structs

import (
	"fmt"
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
)

type EventLoop struct {
	commands  chan iCommand
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
		commands:  make(chan iCommand),
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

		call(e, func() error {
			return e.dispatch(protocol.Decode(data))
		})
	}
}

func call[T any](e *EventLoop, f func() T) T {
	done := make(chan T)
	e.commands <- newCommand(f, done)
	return <-done
}

func (e *EventLoop) Send(data []byte) error {
	return call(e, func() error {
		return e.transport.Send(data)
	})
}

func (e *EventLoop) Run() {
	go e.receiveLoop()

	for {
		cmd := <-e.commands
		cmd.execute()
	}
}
