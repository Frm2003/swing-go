package application

type WindowDriver interface {
	SetTitle(string) error
	Show() error
	Hide() error
	Destroy() error
}

type Window struct {
	driver WindowDriver
}

func NewWindow() *Window {
	return &Window{}
}

func (w *Window) SetTitle(v string) {
	w.driver.SetTitle(v)
}

func (w *Window) SetRoot(v Widget) {

}

func (w *Window) Show() {

}
