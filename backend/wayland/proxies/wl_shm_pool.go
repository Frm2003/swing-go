package proxies

import (
	"swing-go/backend/wayland/infrastruct"
	"swing-go/backend/wayland/protocol"
)

type WlShmPool struct {
	objectId uint32
	send     infrastruct.Sender
}

func NewWlShmPool(newId uint32) *WlShmPool {
	return &WlShmPool{
		objectId: newId,
	}
}

func (wl *WlShmPool) Handle(message *protocol.Message) {

}

func (wl *WlShmPool) GetId() uint32 {
	return wl.objectId
}

func (wl *WlShmPool) SetSender(send infrastruct.Sender) {
	wl.send = send
}
