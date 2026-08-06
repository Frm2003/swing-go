package structs

import (
	"swing-go/backend"
	"swing-go/backend/storage"
	"swing-go/backend/wayland/protocol"
)

type ProxyStore = storage.Store[uint32, Proxy]

type Factory[T Proxy] func(uint32) T

type Proxy interface {
	GetId() uint32
	Handle(*protocol.Message)
}

type senderAware interface {
	SetSender(backend.Sender)
}

type allocator struct {
	currObjId uint32
}

func (a *allocator) NewId() uint32 {
	actual := a.currObjId
	a.currObjId++
	return actual
}

func NewProxyStore() *ProxyStore {
	return storage.NewStore[uint32, Proxy](&allocator{1})
}

func CreateProxy[T Proxy](e *EventLoop, f Factory[T]) T {
	return call(e, func() T {
		newID := e.store.NewId()
		obj := f(newID)

		if p, ok := any(obj).(senderAware); ok {
			p.SetSender(e.Send)
		}

		e.store.Register(obj)
		return obj
	})
}
