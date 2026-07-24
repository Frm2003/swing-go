package wayland

import "swing-go/backend/wayland/proxies"

type Window struct {
	wlSurface *proxies.WlSurface
}

func (w *Window) SetTitle(v string) error {
	return nil
}

func (w *Window) Show() error {
	return nil
}

func (w *Window) Hide() error {
	return nil
}

func (w *Window) Destroy() error {
	return nil
}
