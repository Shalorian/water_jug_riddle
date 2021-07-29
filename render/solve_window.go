package render

import (
	"encoding/json"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type solveLayout struct {
}

func (d *solveLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (d *solveLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)

	prog := objects[1]
	size := objects[0].Size()
	prog.Resize(fyne.NewSize(300, size.Height))
	pos = pos.Add(fyne.NewPos(size.Width, 3))
	prog.Move(pos)

}

func NewSolveWindow(solution [][]int, first, second, measure int, app fyne.App) {
	window := app.NewWindow("Solution")
	window.Resize(fyne.Size{Width: 400, Height: 200})

	titleContainer := container.NewCenter(
		widget.NewLabel(fmt.Sprintf("Measuring %v gallons of water", measure)),
	)

	solutionString, _ := json.Marshal(solution)
	solutionLabel := widget.NewLabel(string(solutionString))
	solutionContainer := container.NewCenter(solutionLabel)

	labelX := widget.NewLabel(fmt.Sprintf("%v gal Jug:  %v", first, 0))
	labelY := widget.NewLabel(fmt.Sprintf("%v gal Jug:  %v", second, 0))

	lableContainer := container.NewGridWithRows(2,
		labelX,
		labelY,
	)

	progressX := widget.NewProgressBar()
	progressY := widget.NewProgressBar()

	animationContainer := container.NewGridWithRows(2, progressX, progressY)

	display := container.New(&solveLayout{}, lableContainer, animationContainer)

	animation := NewAnimation(
		progressX,
		progressY,
		solutionLabel,
		labelX,
		labelY,
		solution,
		first,
		second,
	)

	animateButton := widget.NewButton("Animate", func() {
		animation.Animate()
	})
	animateButton.Disable()
	animation.SetButton(animateButton)

	quitButton := widget.NewButton("Quit", func() {
		window.Close()
	})

	window.SetContent(container.NewVBox(titleContainer, solutionContainer, display, animateButton, quitButton))
	window.Show()

	animation.Animate()
}
