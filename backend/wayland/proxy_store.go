package wayland

import "swing-go/backend/wayland/protocol"

type Factory[T Proxy] func(uint32) T

type Proxy interface {
	Handle(*protocol.Message)
	GetId() uint32
}

type SenderAware interface {
	Proxy
	SetSender(func(*protocol.Message) error)
}

type proxyStore struct {
	currObjectId uint32
	proxies      map[uint32]Proxy
}

func NewProxyStore() *proxyStore {
	return &proxyStore{
		currObjectId: 1,
		proxies:      make(map[uint32]Proxy),
	}
}

func (ps *proxyStore) Register(p Proxy) {
	ps.proxies[p.GetId()] = p
}

func (ps *proxyStore) GenerateNewId() uint32 {
	actual := ps.currObjectId
	ps.currObjectId++
	return actual
}
