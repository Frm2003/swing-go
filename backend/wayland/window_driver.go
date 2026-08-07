package wayland

import "swing-go/backend/wayland/proxies"

type WindowDriver struct {
	surface     *proxies.WlSurface
	xdgSurface  *proxies.XdgSurface
	xdgToplevel *proxies.XdgToplevel
}

func (wd *WindowDriver) SetTitle(v string) error {
	if err := wd.xdgToplevel.SetAppID(v); err != nil {
		return err
	}

	if err := wd.xdgToplevel.SetTitle(v); err != nil {
		return err
	}

	return nil
}
