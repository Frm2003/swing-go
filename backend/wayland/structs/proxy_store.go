package structs

import (
	"swing-go/backend/storage"
	"swing-go/backend/wayland/protocol"
)

type ProxyStore = storage.Store[uint32, Proxy]

type Factory[T Proxy] func(uint32) T

type Proxy interface {
	GetId() uint32
	Handle(*protocol.Message)
}

type SenderAware interface {
	SetSender(func([]byte) error)
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
	newId := e.store.NewId()
	obj := f(newId)

	if p, ok := Proxy(obj).(SenderAware); ok {
		p.SetSender(e.transport.Send)
	}

	e.store.Register(obj)
	return obj
}
