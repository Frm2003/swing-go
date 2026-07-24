package proxies

import (
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/proxies/request"
)

const (
	wlCompositorCreateSurface uint16 = iota
	wlCompositorCreateRegion
)

type WlCompositor struct {
	objectId uint32
	send     func(*protocol.Message) error
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

func (wl *WlCompositor) SetSender(send func(*protocol.Message) error) {
	wl.send = send
}

func (wl *WlCompositor) CreateSurface(newId uint32) error {
	s := request.NewSerializer()

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlCompositorCreateSurface,
		Payload:  s.Uint32(newId).Bytes(),
	})
}
