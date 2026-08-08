package structs

import (
	"fmt"
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
)

type Dispatcher struct {
	store     *ProxyStore
	transport *backend.Transport
	eventLoop *EventLoop
}

func NewDispatcher() *Dispatcher {
	transport := backend.NewTransport(
		protocol.Connect,
		protocol.Frame,
	)

	return &Dispatcher{
		eventLoop: NewEventLoop(),
		store:     NewProxyStore(),
		transport: transport,
	}
}

func (d *Dispatcher) dispatch(message *protocol.Message) error {
	proxy, ok := d.store.Get(message.ObjectID)

	if !ok {
		return fmt.Errorf("Proxy with id %d not found!", message.ObjectID)
	}

	proxy.Handle(message)

	return nil
}

func (d *Dispatcher) Send(data []byte) error {
	return call(d.eventLoop, func() error {
		return d.transport.Send(data)
	})
}

func CreateProxy[T Proxy](d *Dispatcher, f Factory[T]) T {
	return call(d.eventLoop, func() T {
		newID := d.store.NewId()
		obj := f(newID)

		if p, ok := any(obj).(senderAware); ok {
			p.SetSender(d.Send)
		}

		d.store.Register(obj)
		return obj
	})
}

func (d *Dispatcher) Run() error {
	for {
		data, err := d.transport.Receive()

		if err != nil {
			return err
		}

		err = call(d.eventLoop, func() error {
			return d.dispatch(protocol.Decode(data))
		})

		if err != nil {
			return err
		}
	}
}
