package proxies

import (
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/request"
)

const (
	wlDisplaySync uint16 = iota
	wlDisplayGetRegistry
)

type WlDisplay struct {
	objectId uint32
	send     backend.Sender
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

func (wl *WlDisplay) SetSender(send backend.Sender) {
	wl.send = send
}

func (wl *WlDisplay) Sync(newId uint32) error {
	s := request.NewSerializer()

	data := protocol.Encode(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlDisplaySync,
		Payload:  s.Uint32(newId).Bytes(),
	})

	return wl.send(data)
}

func (wl *WlDisplay) GetRegistry(newId uint32) error {
	s := request.NewSerializer()

	data := protocol.Encode(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlDisplayGetRegistry,
		Payload:  s.Uint32(newId).Bytes(),
	})

	return wl.send(data)
}
