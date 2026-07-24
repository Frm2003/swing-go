package wayland

import (
	"fmt"
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
)

type dispatcher struct {
	channelMsg chan *protocol.Message
	proxyStore *proxyStore
	transport  *backend.Transport
}

func NewDispatcher(proxyStore *proxyStore) *dispatcher {
	transport := backend.NewTransport(
		protocol.Connect,
		protocol.Frame,
	)

	return &dispatcher{
		channelMsg: make(chan *protocol.Message),
		proxyStore: proxyStore,
		transport:  transport,
	}
}

func (d *dispatcher) dispatch(message *protocol.Message) error {
	fmt.Println("Recebido: ", message)

	proxy, ok := d.proxyStore.proxies[message.ObjectID]

	if !ok {
		return fmt.Errorf("target_id %d not available", message.ObjectID)
	}

	proxy.Handle(message)

	return nil
}

func (d *dispatcher) receive() error {
	data, err := d.transport.Receive()

	if err != nil {
		return err
	}

	d.channelMsg <- protocol.Decode(data)

	return nil
}

func (d *dispatcher) Send(message *protocol.Message) error {
	fmt.Println("Enviado: ", message)
	return d.transport.Send(protocol.Encode(message))
}

func (d *dispatcher) Run() error {
	channelErr := make(chan error, 1)

	go func() {
		for {
			err := d.receive()

			if err != nil {
				channelErr <- err
				return
			}
		}
	}()

	for {
		select {
		case msg := <-d.channelMsg:
			if err := d.dispatch(msg); err != nil {
				channelErr <- err
			}
		case err := <-channelErr:
			if err != nil {
				return err
			}
		}
	}
}
