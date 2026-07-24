package proxies

import (
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/proxies/request"
	"swing-go/backend/wayland/proxies/response"
)

// wl_registry requests
const (
	wlRegistryBind uint16 = iota
)

// wl_registry events
const (
	wlRegistryGlobal uint16 = iota
	wlRegistryGlobalRemove
)

type WlRegistry struct {
	objectId uint32
	send     func(*protocol.Message) error

	globalStore *globalStore
}

type globalStore struct {
	byName  map[uint32]*global
	byIface map[string][]*global
}

type global struct {
	name    uint32
	iface   string
	version uint32
}

func NewWlRegistry(newId uint32) *WlRegistry {
	globalStore := &globalStore{
		byName:  make(map[uint32]*global),
		byIface: make(map[string][]*global),
	}

	return &WlRegistry{
		objectId:    newId,
		globalStore: globalStore,
	}
}

func (wl *WlRegistry) Handle(message *protocol.Message) {
	switch message.OpCode {
	case wlRegistryGlobal:
		wl.registryGlobal(message.Payload)
	case wlRegistryGlobalRemove:
	}
}

func (wl *WlRegistry) registryGlobal(payload []byte) {
	d := response.NewDeSerializer(payload)

	name := d.Uint32()
	iface := d.String()

	g := &global{
		name:    name,
		iface:   iface,
		version: d.Uint32(),
	}

	wl.globalStore.byName[name] = g
	wl.globalStore.byIface[iface] = append(wl.globalStore.byIface[iface], g)
}

func (wl *WlRegistry) GetId() uint32 {
	return wl.objectId
}

func (wl *WlRegistry) SetSender(send func(*protocol.Message) error) {
	wl.send = send
}

func (w *WlRegistry) Bind(newObjectId uint32, iface string) error {
	g := w.globalStore.byIface[iface][0]

	s := request.NewSerializer()

	message := &protocol.Message{
		ObjectID: w.GetId(),
		OpCode:   wlRegistryBind,
		Payload: s.Uint32(g.name).
			String(g.iface).
			Uint32(g.version).
			Uint32(newObjectId).
			Bytes(),
	}

	return w.send(message)
}
