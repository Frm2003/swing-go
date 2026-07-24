package proxies

import "swing-go/backend/wayland/protocol"

type XdgWmBase struct {
	objectId uint32
	send     func(*protocol.Message) error
}

func NewXdgWmBase(newId uint32) *XdgWmBase {
	return &XdgWmBase{
		objectId: newId,
	}
}

func (wl *XdgWmBase) Handle(message *protocol.Message) {

}

func (wl *XdgWmBase) GetId() uint32 {
	return wl.objectId
}

func (wl *XdgWmBase) GetInterfaceName() string {
	return "xdg_wm_base"
}

func (wl *XdgWmBase) SetSender(send func(*protocol.Message) error) {
	wl.send = send
}
