package render

import (
	"simon.alvarez/water_jug_riddle/handlers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainWindowConfig struct {
	WindowTitle  string
	Title        string
	Labels       []string
	Placeholders []string
	Buttons      []fyne.Widget
}

func NewMainWindow(config MainWindowConfig, app fyne.App) {
	window := app.NewWindow(config.WindowTitle)
	window.Resize(fyne.Size{Width: 400, Height: 200})
	// win.CenterOnScreen()

	titleContainer := container.NewCenter(widget.NewLabel(config.Title))

	labelsContainer := container.NewGridWithRows(3)
	inputsContainer := container.NewVBox()

	inputs := []*widget.Entry{}

	for _, label := range config.Labels {
		labelsContainer.Add(widget.NewLabel(label))
	}

	for _, holder := range config.Placeholders {
		input := widget.NewEntry()
		input.SetPlaceHolder(holder)
		inputsContainer.Add(input)
		inputs = append(inputs, input)
	}

	formContainer := container.NewAdaptiveGrid(2, labelsContainer, inputsContainer)

	quitButton := widget.NewButton("Quit", func() {
		app.Quit()
	})

	calculateButton := widget.NewButton("Solve", func() {
		handler := handlers.NewSolveHandler(app, window, inputs)
		solution, first, second, measure, err := handler.Handle()
		if err != nil {
			return
		}

		NewSolveWindow(solution, first, second, measure, app)

	})

	window.SetContent(container.NewVBox(titleContainer, formContainer, calculateButton, quitButton))
	window.ShowAndRun()
}
