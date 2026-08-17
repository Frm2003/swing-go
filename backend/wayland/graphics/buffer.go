package graphics

import "swing-go/backend/wayland/proxies"

type Buffer struct {
	WlBuffer *proxies.WlBuffer
	Offset   int
	Busy     bool
}
