package proxies

import (
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/request"
	"swing-go/backend/wayland/structs"
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
	send     structs.Sender
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

func (xdg *XdgToplevel) SetSender(send structs.Sender) {
	xdg.send = send
}

func (xdg *XdgToplevel) SetTitle(v string) error {
	s := request.NewSerializer()

	return xdg.send(&protocol.Message{
		ObjectID: xdg.GetId(),
		OpCode:   xdgToplevelSetTitle,
		Payload:  s.String(v).Bytes(),
	})
}

func (xdg *XdgToplevel) SetAppID(v string) error {
	s := request.NewSerializer()

	return xdg.send(&protocol.Message{
		ObjectID: xdg.GetId(),
		OpCode:   xdgToplevelSetAppID,
		Payload:  s.String(v).Bytes(),
	})
}
