package graphics

import (
	"swing-go/backend/wayland/proxies"
	"swing-go/ui"
)

type Driver struct {
	renderer *Renderer

	surface     *proxies.WlSurface
	xdgSurface  *proxies.XdgSurface
	xdgToplevel *proxies.XdgToplevel
}

func NewDriver(
	renderer *Renderer,
	surface *proxies.WlSurface,
	xdgSurface *proxies.XdgSurface,
	xdgToplevel *proxies.XdgToplevel,
) *Driver {
	return &Driver{
		renderer:    renderer,
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

func (d *Driver) Draw(render func(*ui.Canvas)) error {
	buffer, err := d.renderer.AcquireBuffer()

	if err != nil {
		return err
	}

	render(d.renderer.canva)

	if err := d.surface.Attach(buffer.GetId(), 0, 0); err != nil {
		return err
	}

	if err := d.surface.Damage(0, 0, d.renderer.Width(), d.renderer.Height()); err != nil {
		return err
	}

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
