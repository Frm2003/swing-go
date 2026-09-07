package application

import "swing-go/ui"

type Driver interface {
	Draw(func(*ui.Canvas)) error
	SetTitle(v string) error
	Show() error
}

type Window struct {
	driver Driver
	state  *WindowState
}

func (w *Window) Close() {

}

func (w *Window) Draw(root *ui.Root) error {
	root.Layout()
	return w.driver.Draw(root.Draw)
}

func (w *Window) SetSize(width, height int) {

}

func (w *Window) SetTitle(v string) error {
	return w.driver.SetTitle(v)
}

func (w *Window) Show() error {
	return w.driver.Show()
}
