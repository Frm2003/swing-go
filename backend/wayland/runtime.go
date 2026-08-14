package wayland

import (
	"swing-go/application"
	"swing-go/backend/wayland/proxies"
	"swing-go/backend/wayland/structs"
)

type Runtime struct {
	dispatcher *structs.Dispatcher

	wlDIsplay  *proxies.WlDisplay
	WlRegistry *proxies.WlRegistry
	wlcallback *proxies.Wlcallback

	wlCompositor *proxies.WlCompositor
	wlShm        *proxies.WlShm
	XdgWmBase    *proxies.XdgWmBase
}

func NewRuntime() *Runtime {
	dispatcher := structs.NewDispatcher()

	go dispatcher.Run()

	return &Runtime{
		dispatcher: dispatcher,
	}
}

func (r *Runtime) Bootstrap() error {
	r.wlDIsplay = structs.CreateProxy(r.dispatcher, proxies.NewWlDisplay)
	r.WlRegistry = structs.CreateProxy(r.dispatcher, proxies.NewWlRegistry)

	if err := r.wlDIsplay.GetRegistry(r.WlRegistry.GetId()); err != nil {
		return err
	}

	if err := r.sync(); err != nil {
		return err
	}

	r.wlCompositor = structs.CreateProxy(r.dispatcher, proxies.NewWlCompositor)
	if err := r.WlRegistry.Bind(r.wlCompositor.GetId(), r.wlCompositor.GetInterfaceName()); err != nil {
		return err
	}

	r.wlShm = structs.CreateProxy(r.dispatcher, proxies.NewWlShm)
	if err := r.WlRegistry.Bind(r.wlShm.GetId(), r.wlShm.GetInterfaceName()); err != nil {
		return err
	}

	r.XdgWmBase = structs.CreateProxy(r.dispatcher, proxies.NewXdgWmBase)
	if err := r.WlRegistry.Bind(r.XdgWmBase.GetId(), r.XdgWmBase.GetInterfaceName()); err != nil {
		return err
	}

	return nil
}

func (r *Runtime) sync() error {
	r.wlcallback = structs.CreateProxy(r.dispatcher, proxies.NewWlcallback)

	if err := r.wlDIsplay.Sync(r.wlcallback.GetId()); err != nil {
		return err
	}

	r.wlcallback.Wait()

	return nil
}

func (r *Runtime) NewWindow(width, height int) (application.WindowDriver, error) {
	surface := structs.CreateProxy(r.dispatcher, proxies.NewWlSurface)
	if err := r.wlCompositor.CreateSurface(surface.GetId()); err != nil {
		return nil, err
	}

	xdgSurface := structs.CreateProxy(r.dispatcher, proxies.NewXdgSurface)
	if err := r.XdgWmBase.GetXdgSurface(xdgSurface.GetId(), surface.GetId()); err != nil {
		return nil, err
	}

	xdgToplevel := structs.CreateProxy(r.dispatcher, proxies.NewXdgToplevel)
	if err := xdgSurface.GetToplevel(xdgToplevel.GetId()); err != nil {
		return nil, err
	}

	if err := surface.Commit(); err != nil {
		return nil, err
	}

	bm, err := r.createBuffer(width * height * 4)

	if err != nil {
		bm.Close()
		return nil, err
	}

	return &WindowDriver{bm, surface, xdgSurface, xdgToplevel}, nil
}

func (r *Runtime) createBuffer(size int) (*BufferManager, error) {
	wlShmPool := structs.CreateProxy(r.dispatcher, proxies.NewWlShmPool)
	bm, err := NewBufferManager(wlShmPool, size)

	if err != nil {
		return nil, err
	}

	if err := r.wlShm.CreatePool(wlShmPool.GetId(), bm.Fd, size); err != nil {
		return nil, err
	}

	return bm, nil
}
