package proxies

import (
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/request"
)

const (
	xdgSurfaceDestroy uint16 = iota
	xdgSurfaceGetToplevel
	xdgSurfaceGetPopup
	xdgSurfaceSetWindowGeometry
	xdgSurfaceAckConfigure
)

// xdg_surface events
const (
	xdgSurfaceConfigure uint16 = iota
)

type XdgSurface struct {
	objectId uint32
	send     backend.Sender
}

func NewXdgSurface(newId uint32) *XdgSurface {
	return &XdgSurface{
		objectId: newId,
	}
}

func (xdg *XdgSurface) Handle(message *protocol.Message) {

}

func (xdg *XdgSurface) GetId() uint32 {
	return xdg.objectId
}

func (xdg *XdgSurface) SetSender(send backend.Sender) {
	xdg.send = send
}

func (xdg *XdgSurface) GetToplevel(newId uint32) error {
	s := request.NewSerializer()

	data := protocol.Encode(&protocol.Message{
		ObjectID: xdg.GetId(),
		OpCode:   xdgSurfaceGetToplevel,
		Payload:  s.Uint32(newId).Bytes(),
	})

	return xdg.send(data)
}
