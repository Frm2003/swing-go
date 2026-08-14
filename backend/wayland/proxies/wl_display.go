package proxies

import (
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/request"
	"swing-go/backend/wayland/structs"
)

const (
	wlDisplaySync uint16 = iota
	wlDisplayGetRegistry
)

type WlDisplay struct {
	objectId uint32
	send     structs.Sender
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

func (wl *WlDisplay) SetSender(send structs.Sender) {
	wl.send = send
}

func (wl *WlDisplay) Sync(newId uint32) error {
	s := request.NewSerializer()

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlDisplaySync,
		Payload:  s.Uint32(newId).Bytes(),
	})
}

func (wl *WlDisplay) GetRegistry(newId uint32) error {
	s := request.NewSerializer()

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlDisplayGetRegistry,
		Payload:  s.Uint32(newId).Bytes(),
	})
}
