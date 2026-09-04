package application

type Driver interface {
	SetTitle(v string) error
	Show() error
}

type Window struct {
	driver Driver
	state  *WindowState
}

func (w *Window) Show() error {
	return w.driver.Show()
}

func (w *Window) Close() {

}

func (w *Window) SetTitle(v string) error {
	return w.driver.SetTitle(v)
}

func (w *Window) SetSize(width, height int) {

}
