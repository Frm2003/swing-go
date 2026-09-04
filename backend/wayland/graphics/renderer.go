package graphics

import (
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/proxies"
	"swing-go/ui"
)

type Renderer struct {
	canva        *ui.Canvas
	sharedMemory *protocol.SharedMemory
	wlShmPool    *proxies.WlShmPool
}

func NewRenderer(
	canva *ui.Canvas,
	sharedMemory *protocol.SharedMemory,
	wlShmPool *proxies.WlShmPool,
) *Renderer {
	return &Renderer{
		canva:        canva,
		sharedMemory: sharedMemory,
		wlShmPool:    wlShmPool,
	}
}
