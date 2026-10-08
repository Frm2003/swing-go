package main

import (
	"swing-go/application"
	"swing-go/backend/linux"
	"swing-go/backend/wayland"
	"swing-go/fonts"
	"swing-go/ui"
)

func main() {
	runtime := wayland.NewRuntime()

	app := application.NewApp(runtime)
	window := app.NewWindow(800, 600)

	manager := fonts.NewManager(linux.NewFontConfigResolver())

	window.SetTitle("new_window")

	window.Draw(&ui.Root{
		Background: ui.Color{R: 0, G: 0, B: 0, A: 255},
		Child: &ui.Element{
			Margin:  ui.Edge{B: 5, L: 5, R: 5, T: 5},
			Padding: ui.Edge{B: 5, L: 5, R: 5, T: 5},
			Children: []ui.Widget{
				&ui.Element{
					Background: ui.Color{R: 255, G: 0, B: 0, A: 255},
					Padding:    ui.Edge{B: 5, L: 5, R: 5, T: 5},
					Children: []ui.Widget{
						&ui.Text{
							Content:     "tigrinho",
							FontManager: manager,
							TextStyle: fonts.Query{
								Family: "DejaVu Sans",
								Slant:  fonts.SlantItalic,
								Weight: fonts.WeightBold,
							},
						},
					},
				},
			},
		},
	})

	window.Show()

	select {}
}
