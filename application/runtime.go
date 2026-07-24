package application

type Runtime interface {
	Bootstrap() error
	Run() error
	CreateDriver() WindowDriver
}
