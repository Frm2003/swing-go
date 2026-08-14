package proxies

import (
	"swing-go/backend/wayland/infrastruct"
	"swing-go/backend/wayland/protocol"
)

// wl_shm requests
const (
	wlShmCreatePool uint16 = iota
	wlShmRelease
)

// wl_shm events
const (
	wlShmFormat uint16 = iota
)

type WlShm struct {
	objectId uint32
	send     infrastruct.Sender
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

func (wl *WlShm) SetSender(send infrastruct.Sender) {
	wl.send = send
}

func (wl *WlShm) CreatePool(newId uint32, fd, size int) error {
	s := protocol.NewSerializer()

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlShmCreatePool,
		Payload:  s.Uint32(newId).Uint32(uint32(size)).Bytes(),
		Fds:      []int{fd},
	})
}
