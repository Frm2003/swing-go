package proxies

import (
	"swing-go/backend"
	"swing-go/backend/wayland/protocol"
)

// wl_surface requests
const (
	wlSurfaceDestroy uint16 = iota
	wlSurfaceAttach
	wlSurfaceDamage
	wlSurfaceFrame
	wlSurfaceSetOpaqueRegion
	wlSurfaceSetInputRegion
	wlSurfaceCommit
	wlSurfaceSetBufferTransform
	wlSurfaceSetBufferScale
	wlSurfaceDamageBuffer
	wlSurfaceOffset
)

// wl_surface events
const (
	wlSurfaceEnter uint16 = iota
	wlSurfaceLeave
)

type WlSurface struct {
	objectId uint32
	send     backend.Sender
}

func NewWlSurface(newId uint32) *WlSurface {
	return &WlSurface{
		objectId: newId,
	}
}

func (wl *WlSurface) Handle(message *protocol.Message) {

}

func (wl *WlSurface) GetId() uint32 {
	return wl.objectId
}

func (wl *WlSurface) SetSender(send backend.Sender) {
	wl.send = send
}

func (wl *WlSurface) Commit() error {
	data := protocol.Encode(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlSurfaceCommit,
		Payload:  nil,
	})

	return wl.send(data)
}
