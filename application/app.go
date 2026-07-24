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
		runtime: runtime,
		windows: make([]Window, 0),
	}
}

func (a *App) NewWindow() *Window {
	driver, _ := a.runtime.NewWindow()

	window := NewWindow(driver)

	return window
}

func (a *App) Run() error {
	return nil
}
