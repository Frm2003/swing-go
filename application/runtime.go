package application

type Runtime interface {
	Bootstrap() error
	NewWindow(width, height int) (WindowDriver, error)
}
