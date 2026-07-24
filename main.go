package main

import (
	"swing-go/application"
	"swing-go/backend/wayland"
)

func main() {
	runtime := wayland.NewRuntime()

	application.New(runtime)

	// window := app.NewWindow()
	// window.SetTitle("Nova janela")
	// window.Show()

	// app.Run()

	select {}
}
