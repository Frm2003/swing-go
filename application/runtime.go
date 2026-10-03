package application

type Runtime interface {
	Bootstrap() error
	NewWindow(int, int) (Driver, error)
}
