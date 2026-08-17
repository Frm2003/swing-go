package graphics

import (
	"swing-go/application"
	"swing-go/backend/wayland/proxies"
)

type Driver struct {
	BufferManager *BufferManager
	State         *application.WindowState

	Surface     *proxies.WlSurface
	XdgSurface  *proxies.XdgSurface
	XdgToplevel *proxies.XdgToplevel
}

func (wd *Driver) SetTitle(v string) error {
	if err := wd.XdgToplevel.SetAppID(v); err != nil {
		return err
	}

	if err := wd.XdgToplevel.SetTitle(v); err != nil {
		return err
	}

	return nil
}

func (wd *Driver) Show() error {
	if err := wd.XdgSurface.AckConfigure(); err != nil {
		return err
	}

	buf, err := wd.BufferManager.CreateBuffer(wd)

	if err != nil {
		return err
	}

	for y := 0; y < wd.State.Height; y++ {
		for x := 0; x < wd.State.Width; x++ {
			offset := y*wd.BufferManager.Stride + x*4

			wd.BufferManager.Pixels[offset+0] = 0x00 // B
			wd.BufferManager.Pixels[offset+1] = 0x00 // G
			wd.BufferManager.Pixels[offset+2] = 0x00 // R
			wd.BufferManager.Pixels[offset+3] = 0xFF // A
		}
	}

	if err := wd.Surface.Attach(buf.WlBuffer.GetId(), 0, 0); err != nil {
		return err
	}

	if err := wd.Surface.Damage(0, 0, wd.State.Width, wd.State.Height); err != nil {
		return err
	}

	if err := wd.Surface.Commit(); err != nil {
		return err
	}

	return nil
}
