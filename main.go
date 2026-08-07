package main

import (
	"swing-go/application"
	"swing-go/backend/wayland"
)

func main() {
	runtime := wayland.NewRuntime()

	app := application.NewApp(runtime)
	window := app.NewWindow()

	window.SetTitle("new_window")

	select {}
}
