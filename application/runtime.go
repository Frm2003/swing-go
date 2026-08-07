package application

type Runtime interface {
	Bootstrap() error
	NewWindow() (WindowDriver, error)
}
