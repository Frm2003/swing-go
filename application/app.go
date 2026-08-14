package application

type App struct {
	runtime Runtime
}

func NewApp(runtime Runtime) *App {
	if err := runtime.Bootstrap(); err != nil {
		panic(err)
	}

	return &App{
		runtime: runtime,
	}
}

func (a *App) NewWindow(width, height int) *Window {
	driver, err := a.runtime.NewWindow(width, height)

	if err != nil {
		panic(err)
	}

	return &Window{
		driver: driver,
	}
}
