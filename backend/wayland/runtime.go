package wayland

import (
	"swing-go/application"
	"swing-go/backend/wayland/proxies"
)

type Runtime struct {
	dispatcher *dispatcher
	proxyStore *proxyStore

	wlDisplay  *proxies.WlDisplay
	wlRegistry *proxies.WlRegistry

	wlCompositor *proxies.WlCompositor
	wlShm        *proxies.WlShm
	xdgWmBase    *proxies.XdgWmBase
}

func NewRuntime() *Runtime {
	proxyStore := NewProxyStore()

	return &Runtime{
		proxyStore: proxyStore,
		dispatcher: NewDispatcher(proxyStore),
	}
}

func (r *Runtime) Bootstrap() error {
	r.wlDisplay = createProxy(r, proxies.NewWlDisplay)
	r.wlRegistry = createProxy(r, proxies.NewWlRegistry)

	go r.dispatcher.Run()

	if err := r.wlDisplay.GetRegistry(r.wlRegistry.GetId()); err != nil {
		return err
	}

	r.sync()

	r.wlCompositor = createProxy(r, proxies.NewWlCompositor)
	if err := r.wlRegistry.Bind(r.wlCompositor.GetId(), r.wlCompositor.GetInterfaceName()); err != nil {
		return err
	}

	r.wlShm = createProxy(r, proxies.NewWlShm)
	if err := r.wlRegistry.Bind(r.wlShm.GetId(), r.wlShm.GetInterfaceName()); err != nil {
		return err
	}

	r.xdgWmBase = createProxy(r, proxies.NewXdgWmBase)
	if err := r.wlRegistry.Bind(r.xdgWmBase.GetId(), r.xdgWmBase.GetInterfaceName()); err != nil {
		return err
	}

	return nil
}

func (r *Runtime) sync() error {
	wlCallback := createProxy(r, proxies.NewWlCallback)

	if err := r.wlDisplay.Sync(wlCallback.GetId()); err != nil {
		return err
	}

	wlCallback.Wait()

	return nil
}

func (r *Runtime) Run() error {
	return nil
}

func (r *Runtime) NewWindow() (application.WindowDriver, error) {
	wlSurface := createProxy(r, proxies.NewWlSurface)
	if err := r.wlCompositor.CreateSurface(wlSurface.GetId()); err != nil {
		return nil, err
	}

	return &Window{}, nil
}

func createProxy[T Proxy](r *Runtime, f Factory[T]) T {
	newId := r.proxyStore.GenerateNewId()
	obj := f(newId)

	if p, ok := Proxy(obj).(SenderAware); ok {
		p.SetSender(r.dispatcher.Send)
	}

	r.proxyStore.Register(obj)
	return obj
}
