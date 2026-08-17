package application

type WindowDriver interface {
	SetTitle(v string) error
	Show() error
}
