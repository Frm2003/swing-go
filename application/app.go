package application

type App struct {
	windows []Window
	runtime Runtime
}

func New(runtime Runtime) *App {
	if err := runtime.Bootstrap(); err != nil {
		panic(err)
	}

	return &App{
		windows: make([]Window, 0),
	}
}

func (a *App) NewWindow() *Window {
	window := NewWindow()
	a.windows = append(a.windows, *window)
	return window
}

func (a *App) Run() error {
	return nil
}
