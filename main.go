package main

import (
	"swing-go/application"
	"swing-go/backend/wayland"
)

func main() {
	runtime := wayland.NewRuntime()

	app := application.New(runtime)

	app.NewWindow()
	// window.SetTitle("Nova janela")
	// window.Show()

	// app.Run()

	select {}
}
