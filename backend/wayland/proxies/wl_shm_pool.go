package proxies

import (
	"swing-go/backend/wayland/infrastruct"
	"swing-go/backend/wayland/protocol"
)

// wl_shm_pool requests
const (
	wlShmPoolCreateBuffer uint16 = iota // 0
	wlShmPoolDestroy                    // 1
	wlShmPoolResize                     // 2
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

func (wl *WlShmPool) CreateBuffer(newId uint32, offset, width, height, stride int32, format uint32) error {
	s := protocol.NewSerializer()

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlShmPoolCreateBuffer,
		Payload: s.Uint32(newId).
			Uint32(uint32(offset)).
			Uint32(uint32(width)).
			Uint32(uint32(height)).
			Uint32(uint32(stride)).
			Uint32(format).
			Bytes(),
	})
}
