package proxies

import (
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/structs"
)

type WlShm struct {
	objectId uint32
	send     structs.Sender
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

func (wl *WlShm) SetSender(send structs.Sender) {
	wl.send = send
}
