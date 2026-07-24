package proxies

import (
	"encoding/binary"
	"swing-go/backend/wayland/protocol"
)

const (
	wlDisplaySync uint16 = iota
	wlDisplayGetRegistry
)

type WlDisplay struct {
	objectId uint32
	send     func(*protocol.Message) error
}

func NewWlDisplay(newId uint32) *WlDisplay {
	return &WlDisplay{
		objectId: newId,
	}
}

func (wl *WlDisplay) Handle(message *protocol.Message) {

}

func (wl *WlDisplay) GetId() uint32 {
	return wl.objectId
}

func (wl *WlDisplay) SetSender(send func(*protocol.Message) error) {
	wl.send = send
}

func (wl *WlDisplay) Sync(newId uint32) error {
	payload := make([]byte, 4)

	binary.LittleEndian.PutUint32(payload[:4], newId)

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlDisplaySync,
		Payload:  payload,
	})
}

func (wl *WlDisplay) GetRegistry(newId uint32) error {
	payload := make([]byte, 4)

	binary.LittleEndian.PutUint32(payload[:4], newId)

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlDisplayGetRegistry,
		Payload:  payload,
	})
}
