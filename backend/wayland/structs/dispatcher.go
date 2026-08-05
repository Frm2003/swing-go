package structs

import (
	"fmt"
	"swing-go/backend/wayland/protocol"
)

type getFunc func(uint32) (Proxy, bool)

type Dispatcher struct {
	get getFunc
}

func NewDispatcher(g getFunc) *Dispatcher {
	return &Dispatcher{
		get: g,
	}
}

func (d *Dispatcher) dispatch(message *protocol.Message) error {
	fmt.Println("Recebido: ", message)

	proxy, ok := d.get(message.ObjectID)

	if !ok {
		return fmt.Errorf("Proxy with id %d not found!", message.ObjectID)
	}

	proxy.Handle(message)

	return nil
}
