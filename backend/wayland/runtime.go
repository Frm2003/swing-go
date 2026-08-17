package wayland

import (
	"swing-go/application"
	"swing-go/backend/wayland/graphics"
	"swing-go/backend/wayland/infrastruct"
	"swing-go/backend/wayland/proxies"
)

type Runtime struct {
	dispatcher *infrastruct.Dispatcher

	wlDIsplay  *proxies.WlDisplay
	WlRegistry *proxies.WlRegistry

	wlCompositor *proxies.WlCompositor
	wlShm        *proxies.WlShm
	XdgWmBase    *proxies.XdgWmBase
}

func NewRuntime() *Runtime {
	dispatcher := infrastruct.NewDispatcher()

	go dispatcher.Run()

	return &Runtime{
		dispatcher: dispatcher,
	}
}

func (r *Runtime) Bootstrap() error {
	r.wlDIsplay = infrastruct.CreateProxy(r.dispatcher, proxies.NewWlDisplay)
	r.WlRegistry = infrastruct.CreateProxy(r.dispatcher, proxies.NewWlRegistry)

	if err := r.wlDIsplay.GetRegistry(r.WlRegistry.GetId()); err != nil {
		return err
	}

	if err := r.sync(); err != nil {
		return err
	}

	r.wlCompositor = infrastruct.CreateProxy(r.dispatcher, proxies.NewWlCompositor)
	if err := r.WlRegistry.Bind(r.wlCompositor.GetId(), r.wlCompositor.GetInterfaceName()); err != nil {
		return err
	}

	r.wlShm = infrastruct.CreateProxy(r.dispatcher, proxies.NewWlShm)
	if err := r.WlRegistry.Bind(r.wlShm.GetId(), r.wlShm.GetInterfaceName()); err != nil {
		return err
	}

	r.XdgWmBase = infrastruct.CreateProxy(r.dispatcher, proxies.NewXdgWmBase)
	if err := r.WlRegistry.Bind(r.XdgWmBase.GetId(), r.XdgWmBase.GetInterfaceName()); err != nil {
		return err
	}

	return nil
}

func (r *Runtime) sync() error {
	wlcallback := infrastruct.CreateProxy(r.dispatcher, proxies.NewWlcallback)

	if err := r.wlDIsplay.Sync(wlcallback.GetId()); err != nil {
		return err
	}

	wlcallback.Wait()

	return nil
}

func (r *Runtime) NewWindow(state *application.WindowState) (application.WindowDriver, error) {
	surface := infrastruct.CreateProxy(r.dispatcher, proxies.NewWlSurface)
	if err := r.wlCompositor.CreateSurface(surface.GetId()); err != nil {
		return nil, err
	}

	xdgSurface := infrastruct.CreateProxy(r.dispatcher, proxies.NewXdgSurface)
	if err := r.XdgWmBase.GetXdgSurface(xdgSurface.GetId(), surface.GetId()); err != nil {
		return nil, err
	}

	xdgToplevel := infrastruct.CreateProxy(r.dispatcher, proxies.NewXdgToplevel)
	if err := xdgSurface.GetToplevel(xdgToplevel.GetId()); err != nil {
		return nil, err
	}

	if err := surface.Commit(); err != nil {
		return nil, err
	}

	if err := r.sync(); err != nil {
		return nil, err
	}

	bufferManager, err := graphics.NewBufferManager(r.newShmPool, r.createBuffer, state.Width, state.Height)

	if err != nil {
		return nil, err
	}

	return &graphics.Driver{
		State:         state,
		BufferManager: bufferManager,
		Surface:       surface,
		XdgSurface:    xdgSurface,
		XdgToplevel:   xdgToplevel,
	}, nil
}

func (r *Runtime) newShmPool(fd, size int) (*proxies.WlShmPool, error) {
	whShmPool := infrastruct.CreateProxy(r.dispatcher, proxies.NewWlShmPool)

	if err := r.wlShm.CreatePool(whShmPool.GetId(), fd, size); err != nil {
		return nil, err
	}

	return whShmPool, nil
}

func (r *Runtime) createBuffer(d *graphics.Driver) (*graphics.Buffer, error) {
	wlBuffer := infrastruct.CreateProxy(r.dispatcher, proxies.NewWlBuffer)

	offset := len(d.BufferManager.Buffers) * d.BufferManager.Size

	if err := d.BufferManager.WlShmPool.CreateBuffer(
		wlBuffer.GetId(),
		int32(offset),
		int32(d.State.Width),
		int32(d.State.Height),
		int32(d.BufferManager.Stride),
		0,
	); err != nil {
		return nil, err
	}

	newBuffer := &graphics.Buffer{
		WlBuffer: wlBuffer,
		Offset:   offset,
		Busy:     false,
	}

	return newBuffer, nil
}
