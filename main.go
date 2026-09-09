package main

import (
	"swing-go/application"
	"swing-go/backend/wayland"
	"swing-go/ui"
)

func main() {
	runtime := wayland.NewRuntime()

	app := application.NewApp(runtime)
	window := app.NewWindow(800, 600)

	window.SetTitle("new_window")

	window.Draw(&ui.Root{
		Backgroud: ui.Color{R: 0, G: 0, B: 0, A: 255},
		Childrens: []ui.Widget{
			&ui.Div{
				Backgroud: ui.Color{R: 0, G: 0, B: 255, A: 255},
				Width:     200,
				Height:    100,
			},
		},
	})

	window.Show()

	select {}
}
