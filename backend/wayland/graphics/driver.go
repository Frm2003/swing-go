package graphics

import (
	"swing-go/backend/wayland/proxies"
)

type Driver struct {
	surface     *proxies.WlSurface
	xdgSurface  *proxies.XdgSurface
	xdgToplevel *proxies.XdgToplevel
}

func NewDriver(
	surface *proxies.WlSurface,
	xdgSurface *proxies.XdgSurface,
	xdgToplevel *proxies.XdgToplevel,
) *Driver {
	return &Driver{
		surface:     surface,
		xdgSurface:  xdgSurface,
		xdgToplevel: xdgToplevel,
	}
}

func (d *Driver) SetTitle(v string) error {
	if err := d.xdgToplevel.SetAppID(v); err != nil {
		return err
	}

	if err := d.xdgToplevel.SetTitle(v); err != nil {
		return err
	}

	return nil
}

func (d *Driver) Draw() error {
	// criar buffer

	// if err := d.Surface.Attach(buf.WlBuffer.GetId(), 0, 0); err != nil {
	// 	return err
	// }

	// if err := d.Surface.Damage(0, 0, d.State.Width, d.State.Height); err != nil {
	// 	return err
	// }

	return nil
}

func (d *Driver) Show() error {
	if err := d.xdgSurface.AckConfigure(); err != nil {
		return err
	}

	if err := d.surface.Commit(); err != nil {
		return err
	}

	return nil
}
