package application

type Runtime interface {
	Bootstrap() error
	Run() error
	NewWindow() (WindowDriver, error)
}
