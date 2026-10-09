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
		Background: ui.Color{R: 0, G: 0, B: 0, A: 255},
		Child: &ui.Element{
			Children: []ui.Widget{
				&ui.Element{
					Background: ui.Color{R: 255, G: 0, B: 0, A: 255},
					Display:    ui.Row,
					Height:     50,
					Width:      50,
					Children: []ui.Widget{
						&ui.Element{
							Background: ui.Color{R: 0, G: 0, B: 255, A: 255},
							Height:     50,
							Margin:     ui.Edge{B: 5, L: 5, R: 5, T: 5},
							Width:      50,
						},
						&ui.Element{
							Background: ui.Color{R: 0, G: 0, B: 255, A: 255},
							Height:     50,
							Margin:     ui.Edge{B: 5, L: 5, R: 5, T: 20},
							Width:      50,
						},
						&ui.Element{
							Background: ui.Color{R: 0, G: 0, B: 255, A: 255},
							Height:     50,
							Margin:     ui.Edge{B: 5, L: 5, R: 5, T: 5},
							Width:      50,
						},
					},
				},
			},
		},
	})

	window.Show()

	select {}
}
