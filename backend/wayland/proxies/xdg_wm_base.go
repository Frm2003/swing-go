package proxies

import (
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/proxies/request"
)

type XdgWmBase struct {
	objectId uint32
	send     func(*protocol.Message) error
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

func (xdg *XdgWmBase) SetSender(send func(*protocol.Message) error) {
	xdg.send = send
}

func (xdg *XdgWmBase) GetXdgSurface(newId uint32) error {
	s := request.NewSerializer()

	return xdg.send(&protocol.Message{
		ObjectID: xdg.GetId(),
		OpCode:   0,
		Payload:  s.Uint32(newId).Bytes(),
	})
}
