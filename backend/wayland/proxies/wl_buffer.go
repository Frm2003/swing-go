package proxies

import (
	"swing-go/backend/wayland/infrastruct"
	"swing-go/backend/wayland/protocol"
)

type WlBuffer struct {
	objectId uint32
	send     infrastruct.Sender
}

func NewWlBuffer(newId uint32) *WlBuffer {
	return &WlBuffer{
		objectId: newId,
	}
}

func (wl *WlBuffer) Handle(message *protocol.Message) {

}

func (wl *WlBuffer) GetId() uint32 {
	return wl.objectId
}

func (wl *WlBuffer) SetSender(send infrastruct.Sender) {
	wl.send = send
}
