package wayland

import (
	"swing-go/application"
	"swing-go/backend/wayland/graphics"
	"swing-go/backend/wayland/infrastruct"
	"swing-go/backend/wayland/protocol"
	"swing-go/backend/wayland/proxies"
	"swing-go/ui"
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

func (r *Runtime) NewWindow(width, height int) (application.Driver, error) {
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

	renderer, err := r.NewRenderer(width, height)

	if err != nil {
		return nil, err
	}

	return graphics.NewDriver(
		renderer,
		surface,
		xdgSurface,
		xdgToplevel,
	), nil
}

func (r *Runtime) NewRenderer(width, height int) (*graphics.Renderer, error) {
	stride := width * 4
	size := stride * height

	memory, err := protocol.NewSharedMemory(size)

	if err != nil {
		return nil, err
	}

	wlShmPool := infrastruct.CreateProxy(r.dispatcher, proxies.NewWlShmPool)
	if err := r.wlShm.CreatePool(wlShmPool.GetId(), memory.Fd, memory.Size); err != nil {
		return nil, err
	}

	canva := &ui.Canvas{
		Height: height,
		Pixels: memory.Pixels,
		Stride: stride,
		Width:  width,
	}

	return graphics.NewRenderer(
		r.NewBuffer,
		canva,
		memory,
		wlShmPool,
	), nil
}

func (r *Runtime) NewBuffer(renderer *graphics.Renderer) (*proxies.WlBuffer, error) {
	wlBuffer := infrastruct.CreateProxy(r.dispatcher, proxies.NewWlBuffer)

	if err := renderer.WlShmPool.CreateBuffer(
		wlBuffer.GetId(),
		0,
		int32(renderer.Width()),
		int32(renderer.Height()),
		int32(renderer.Width())*4, // stride
		renderer.Format(),
	); err != nil {
		return nil, err
	}

	return wlBuffer, nil
}
