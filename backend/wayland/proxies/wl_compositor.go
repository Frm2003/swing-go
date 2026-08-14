package proxies

import (
	"swing-go/backend/wayland/infrastruct"
	"swing-go/backend/wayland/protocol"
)

const (
	wlCompositorCreateSurface uint16 = iota
	wlCompositorCreateRegion
)

type WlCompositor struct {
	objectId uint32
	send     infrastruct.Sender
}

func NewWlCompositor(newId uint32) *WlCompositor {
	return &WlCompositor{
		objectId: newId,
	}
}

func (wl *WlCompositor) Handle(message *protocol.Message) {

}

func (wl *WlCompositor) GetId() uint32 {
	return wl.objectId
}

func (wl *WlCompositor) GetInterfaceName() string {
	return "wl_compositor"
}

func (wl *WlCompositor) SetSender(send infrastruct.Sender) {
	wl.send = send
}

func (wl *WlCompositor) CreateSurface(newId uint32) error {
	s := protocol.NewSerializer()

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlCompositorCreateSurface,
		Payload:  s.Uint32(newId).Bytes(),
	})
}
