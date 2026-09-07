package graphics

import (
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/proxies"
	"swing-go/ui"
)

type AllocBuffer func(r *Renderer) (*proxies.WlBuffer, error)

type Renderer struct {
	allocBuffer AllocBuffer

	canva        *ui.Canvas
	sharedMemory *protocol.SharedMemory
	WlShmPool    *proxies.WlShmPool
}

func NewRenderer(
	allocBuffer AllocBuffer,
	canva *ui.Canvas,
	sharedMemory *protocol.SharedMemory,
	wlShmPool *proxies.WlShmPool,
) *Renderer {
	return &Renderer{
		allocBuffer:  allocBuffer,
		canva:        canva,
		sharedMemory: sharedMemory,
		WlShmPool:    wlShmPool,
	}
}

func (r *Renderer) Format() uint32 {
	return 0
}

func (r *Renderer) Height() int {
	return r.canva.Height
}

func (r *Renderer) Width() int {
	return r.canva.Width
}

func (r *Renderer) AcquireBuffer() (*proxies.WlBuffer, error) {
	return r.allocBuffer(r)
}
