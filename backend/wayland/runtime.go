package wayland

import (
	"swing-go/backend/wayland/proxies"
	"swing-go/backend/wayland/structs"
)

type Runtime struct {
	loop *structs.EventLoop

	wlDIsplay  *proxies.WlDisplay
	WlRegistry *proxies.WlRegistry
	wlcallback *proxies.Wlcallback

	wlCompositor *proxies.WlCompositor
	wlShm        *proxies.WlShm
	XdgWmBase    *proxies.XdgWmBase
}

func NewRuntime() *Runtime {
	return &Runtime{
		loop: structs.NewEventLoop(),
	}
}

func (r *Runtime) Bootstrap() error {
	r.wlDIsplay = structs.CreateProxy(r.loop, proxies.NewWlDisplay)
	r.WlRegistry = structs.CreateProxy(r.loop, proxies.NewWlRegistry)

	if err := r.wlDIsplay.GetRegistry(r.WlRegistry.GetId()); err != nil {
		return err
	}

	if err := r.Sync(); err != nil {
		return err
	}

	r.wlCompositor = structs.CreateProxy(r.loop, proxies.NewWlCompositor)
	if err := r.WlRegistry.Bind(r.wlCompositor.GetId(), r.wlCompositor.GetInterfaceName()); err != nil {
		return err
	}

	r.wlShm = structs.CreateProxy(r.loop, proxies.NewWlShm)
	if err := r.WlRegistry.Bind(r.wlShm.GetId(), r.wlShm.GetInterfaceName()); err != nil {
		return err
	}

	r.XdgWmBase = structs.CreateProxy(r.loop, proxies.NewXdgWmBase)
	if err := r.WlRegistry.Bind(r.XdgWmBase.GetId(), r.XdgWmBase.GetInterfaceName()); err != nil {
		return err
	}

	return nil
}

func (r *Runtime) Sync() error {
	r.wlcallback = structs.CreateProxy(r.loop, proxies.NewWlcallback)

	if err := r.wlDIsplay.Sync(r.wlcallback.GetId()); err != nil {
		return err
	}

	r.wlcallback.Wait()

	return nil
}
