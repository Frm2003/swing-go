package proxies

import (
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/request"
)

// xdg_wm_base requests
const (
	xdgWmBaseDestroy uint16 = iota
	xdgWmBaseCreatePositioner
	xdgWmBaseGetXdgSurface
	xdgWmBasePong
)

// xdg_wm_base events
const (
	xdgWmBasePing uint16 = iota
)

type XdgWmBase struct {
	objectId uint32
	send     backend.Sender
}

func NewXdgWmBase(newId uint32) *XdgWmBase {
	return &XdgWmBase{
		objectId: newId,
	}
}

func (xdg *XdgWmBase) Handle(message *protocol.Message) {

}

func (xdg *XdgWmBase) GetId() uint32 {
	return xdg.objectId
}

func (xdg *XdgWmBase) GetInterfaceName() string {
	return "xdg_wm_base"
}

func (xdg *XdgWmBase) SetSender(send backend.Sender) {
	xdg.send = send
}

func (xdg *XdgWmBase) GetXdgSurface(newId, surfaceId uint32) error {
	s := request.NewSerializer()

	data := protocol.Encode(&protocol.Message{
		ObjectID: xdg.GetId(),
		OpCode:   xdgWmBaseGetXdgSurface,
		Payload:  s.Uint32(newId).Uint32(surfaceId).Bytes(),
	})

	return xdg.send(data)
}
