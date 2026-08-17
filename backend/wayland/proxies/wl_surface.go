package proxies

import (
	"swing-go/backend/wayland/infrastruct"
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
	send     infrastruct.Sender
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

func (wl *WlSurface) SetSender(send infrastruct.Sender) {
	wl.send = send
}

func (wl *WlSurface) Attach(newId, x, y uint32) error {
	s := protocol.NewSerializer()

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlSurfaceAttach,
		Payload: s.Uint32(newId).
			Uint32(x).
			Uint32(y).
			Bytes(),
	})
}

func (wl *WlSurface) Damage(x, y, width, height int) error {
	s := protocol.NewSerializer()

	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlSurfaceDamage,
		Payload: s.Uint32(uint32(x)).
			Uint32(uint32(y)).
			Uint32(uint32(width)).
			Uint32(uint32(height)).
			Bytes(),
	})
}

func (wl *WlSurface) Commit() error {
	return wl.send(&protocol.Message{
		ObjectID: wl.GetId(),
		OpCode:   wlSurfaceCommit,
		Payload:  nil,
	})
}
