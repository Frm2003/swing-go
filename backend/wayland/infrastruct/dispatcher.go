package infrastruct

import (
	"fmt"
	"swing-go/backend/wayland/protocol"
)

type Sender func(*protocol.Message) error

type Dispatcher struct {
	store     *ProxyStore
	transport *protocol.Transport
	eventLoop *EventLoop
}

func NewDispatcher() *Dispatcher {
	transport := protocol.NewTransport()

	return &Dispatcher{
		eventLoop: NewEventLoop(),
		store:     NewProxyStore(),
		transport: transport,
	}
}

func (d *Dispatcher) dispatch(message *protocol.Message) error {
	fmt.Println("Recebido: ", message)

	proxy, ok := d.store.Get(message.ObjectID)

	if !ok {
		return fmt.Errorf("Proxy with id %d not found!", message.ObjectID)
	}

	proxy.Handle(message)

	return nil
}

func (d *Dispatcher) Send(message *protocol.Message) error {
	fmt.Println("Enviado: ", message)

	return call(d.eventLoop, func() error {
		return d.transport.Send(message)
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
		msg, err := d.transport.Receive()

		if err != nil {
			return err
		}

		err = call(d.eventLoop, func() error {
			return d.dispatch(msg)
		})

		if err != nil {
			return err
		}
	}
}
