package main

import (
	"simon.alvarez/water_jug_riddle/render"

	fyneApp "fyne.io/fyne/v2/app"
)

func main() {
	app := fyneApp.New()

	config := render.MainWindowConfig{
		WindowTitle:  "Simon's Jug Riddle Solver",
		Title:        "Welcome to Simon's Jug Riddle Solver",
		Labels:       []string{"Jug X", "Jug Y", "Gallons to measure"},
		Placeholders: []string{"Gallons of X", "Gallons of Y", "Gallons to measure"},
	}

	render.NewMainWindow(config, app)
}
