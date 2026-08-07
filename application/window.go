package application

type Window struct {
	driver WindowDriver
}

func (w *Window) Show() {

}

func (w *Window) Close() {

}

func (w *Window) SetTitle(v string) error {
	return w.driver.SetTitle(v)
}

func (w *Window) SetSize(width, height int) {

}
