package proxies

import (
	"swing-go/backend/wayland/protocol"
)

type XdgSurface struct {
	objectId uint32
	send     func(*protocol.Message) error
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

func (xdg *XdgSurface) GetInterfaceName() string {
	return "xdg_wm_base"
}

func (xdg *XdgSurface) SetSender(send func(*protocol.Message) error) {
	xdg.send = send
}
