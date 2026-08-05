package wayland

import (
	"fmt"
	"swing-go/backend/wayland/proxies"
	"swing-go/backend/wayland/structs"
)

type Runtime struct {
	loop *structs.EventLoop

	wlDIsplay  *proxies.WlDisplay
	WlRegistry *proxies.WlRegistry
}

func NewRuntime() *Runtime {
	return &Runtime{
		loop: structs.NewEventLoop(),
	}
}

func (r *Runtime) Bootstrap() {
	r.loop.Do(func() error {
		fmt.Println("wl_display.getRegistry")

		r.wlDIsplay = structs.CreateProxy(r.loop, proxies.NewWlDisplay)
		r.WlRegistry = structs.CreateProxy(r.loop, proxies.NewWlRegistry)

		return nil
	})

	r.wlDIsplay.GetRegistry(r.WlRegistry.GetId())

	r.loop.Do(func() error {
		fmt.Println("bind_compositor")
		return nil
	})
}
