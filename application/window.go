package application

type Window struct {
	state  *WindowState
	driver WindowDriver
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
