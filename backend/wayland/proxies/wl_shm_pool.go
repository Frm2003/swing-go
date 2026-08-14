package proxies

import (
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/structs"
)

type WlShmPool struct {
	objectId uint32
	send     structs.Sender
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

func (wl *WlShmPool) SetSender(send structs.Sender) {
	wl.send = send
}
