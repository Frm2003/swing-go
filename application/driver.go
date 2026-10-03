package application

import "swing-go/ui"

type Driver interface {
	Draw(func(*ui.Canvas)) error
	SetTitle(v string) error
	Show() error
}
