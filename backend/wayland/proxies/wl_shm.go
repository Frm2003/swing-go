package proxies

import (
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
)

type WlShm struct {
	objectId uint32
	send     backend.Sender
}

func NewWlShm(newId uint32) *WlShm {
	return &WlShm{
		objectId: newId,
	}
}

func (wl *WlShm) Handle(message *protocol.Message) {

}

func (wl *WlShm) GetId() uint32 {
	return wl.objectId
}

func (wl *WlShm) GetInterfaceName() string {
	return "wl_shm"
}

func (wl *WlShm) SetSender(send backend.Sender) {
	wl.send = send
}
