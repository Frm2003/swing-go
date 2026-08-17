package application

type Runtime interface {
	Bootstrap() error
	NewWindow(*WindowState) (WindowDriver, error)
}
