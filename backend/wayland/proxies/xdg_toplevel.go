package proxies

import (
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/request"
)

// xdg_toplevel requests
const (
	xdgToplevelDestroy uint16 = iota
	xdgToplevelSetParent
	xdgToplevelSetTitle
	xdgToplevelSetAppID
	xdgToplevelShowWindowMenu
	xdgToplevelMove
	xdgToplevelResize
	xdgToplevelSetMaxSize
	xdgToplevelSetMinSize
	xdgToplevelSetMaximized
	xdgToplevelUnsetMaximized
	xdgToplevelSetFullscreen
	xdgToplevelUnsetFullscreen
	xdgToplevelSetMinimized
)

// xdg_toplevel events
const (
	xdgToplevelConfigure uint16 = iota
	xdgToplevelClose
	xdgToplevelConfigureBounds
	xdgToplevelWmCapabilities
)

type XdgToplevel struct {
	objectId uint32
	send     backend.Sender
}

func NewXdgToplevel(newId uint32) *XdgToplevel {
	return &XdgToplevel{
		objectId: newId,
	}
}

func (xdg *XdgToplevel) Handle(message *protocol.Message) {

}

func (xdg *XdgToplevel) GetId() uint32 {
	return xdg.objectId
}

func (xdg *XdgToplevel) SetSender(send backend.Sender) {
	xdg.send = send
}

func (xdg *XdgToplevel) SetTitle(v string) error {
	s := request.NewSerializer()

	data := protocol.Encode(&protocol.Message{
		ObjectID: xdg.GetId(),
		OpCode:   xdgToplevelSetTitle,
		Payload:  s.String(v).Bytes(),
	})

	return xdg.send(data)
}

func (xdg *XdgToplevel) SetAppID(v string) error {
	s := request.NewSerializer()

	data := protocol.Encode(&protocol.Message{
		ObjectID: xdg.GetId(),
		OpCode:   xdgToplevelSetAppID,
		Payload:  s.String(v).Bytes(),
	})

	return xdg.send(data)
}
