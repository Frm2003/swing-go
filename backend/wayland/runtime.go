package wayland

import (
	"swing-go/backend/wayland/proxies"
	"swing-go/backend/wayland/structs"
)

type Runtime struct {
	loop *structs.EventLoop

	wlDIsplay  *proxies.WlDisplay
	WlRegistry *proxies.WlRegistry
	wlCallback *proxies.WlCallback

	wlCompositor *proxies.WlCompositor
	wlShm        *proxies.WlShm
	XdgWmBase    *proxies.XdgWmBase
}

func NewRuntime() *Runtime {
	return &Runtime{
		loop: structs.NewEventLoop(),
	}
}

func (r *Runtime) Bootstrap() {
	r.loop.Do(func() error {
		r.wlDIsplay = structs.CreateProxy(r.loop, proxies.NewWlDisplay)
		r.WlRegistry = structs.CreateProxy(r.loop, proxies.NewWlRegistry)

		return nil
	})

	r.wlDIsplay.GetRegistry(r.WlRegistry.GetId())

	r.Sync()

	r.loop.Do(func() error {
		r.wlCompositor = structs.CreateProxy(r.loop, proxies.NewWlCompositor)
		return nil
	})

	r.loop.Do(func() error {
		r.wlShm = structs.CreateProxy(r.loop, proxies.NewWlShm)
		return nil
	})

	r.loop.Do(func() error {
		r.XdgWmBase = structs.CreateProxy(r.loop, proxies.NewXdgWmBase)
		return nil
	})

	r.WlRegistry.Bind(r.wlCompositor.GetId(), r.wlCompositor.GetInterfaceName())
	r.WlRegistry.Bind(r.wlShm.GetId(), r.wlShm.GetInterfaceName())
	r.WlRegistry.Bind(r.XdgWmBase.GetId(), r.XdgWmBase.GetInterfaceName())
}

func (r *Runtime) Sync() error {
	r.loop.Do(func() error {
		r.wlCallback = structs.CreateProxy(r.loop, proxies.NewWlCallback)
		return nil
	})

	if err := r.wlDIsplay.Sync(r.wlCallback.GetId()); err != nil {
		return err
	}

	r.wlCallback.Wait()

	return nil
}
