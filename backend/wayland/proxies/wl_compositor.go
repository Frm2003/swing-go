package proxies

import "swing-go/backend/wayland/protocol"

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
