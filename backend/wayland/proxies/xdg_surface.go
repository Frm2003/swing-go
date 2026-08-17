package proxies

import (
	"swing-go/backend/wayland/infrastruct"
	"swing-go/backend/wayland/protocol"
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
	send     infrastruct.Sender

	serial uint32
}

func NewXdgSurface(newId uint32) *XdgSurface {
	return &XdgSurface{
		objectId: newId,
	}
}

func (xdg *XdgSurface) Handle(message *protocol.Message) {
	switch message.OpCode {
	case xdgSurfaceConfigure:
		xdg.handleConfigure(message.Payload)
	}
}

func (xdg *XdgSurface) handleConfigure(payload []byte) {
	d := protocol.NewDeSerializer(payload)
	xdg.serial = d.Uint32()
}

func (xdg *XdgSurface) GetId() uint32 {
	return xdg.objectId
}

func (xdg *XdgSurface) SetSender(send infrastruct.Sender) {
	xdg.send = send
}

func (xdg *XdgSurface) GetToplevel(newId uint32) error {
	s := protocol.NewSerializer()

	return xdg.send(&protocol.Message{
		ObjectID: xdg.GetId(),
		OpCode:   xdgSurfaceGetToplevel,
		Payload:  s.Uint32(newId).Bytes(),
	})
}

func (xdg *XdgSurface) AckConfigure() error {
	s := protocol.NewSerializer()

	return xdg.send(&protocol.Message{
		ObjectID: xdg.GetId(),
		OpCode:   xdgSurfaceAckConfigure,
		Payload:  s.Uint32(xdg.serial).Bytes(),
	})
}
