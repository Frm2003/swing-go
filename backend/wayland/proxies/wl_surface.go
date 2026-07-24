package proxies

import "swing-go/backend/wayland/protocol"

type WlSurface struct {
	objectId uint32
	send     func(*protocol.Message) error
}

func NewWlSurface(newId uint32) *WlSurface {
	return &WlSurface{
		objectId: newId,
	}
}

func (wl *WlSurface) Handle(message *protocol.Message) {

}

func (wl *WlSurface) GetId() uint32 {
	return wl.objectId
}

func (wl *WlSurface) SetSender(send func(*protocol.Message) error) {
	wl.send = send
}
