package render

import (
	"fmt"
	"math"
	"time"

	"simon.alvarez/water_jug_riddle/solver"

	"fyne.io/fyne/v2/widget"
)

type Animation interface {
	Animate()
	SetButton(button *widget.Button)
}

type animator struct {
	Button         *widget.Button
	SolutionLabel  *widget.Label
	ProgressBarX   *widget.ProgressBar
	ProgressBarY   *widget.ProgressBar
	ProgressLevelX *widget.Label
	ProgressLevelY *widget.Label
	Solution       [][]int
	First          int
	Second         int
}

func NewAnimation(progressBarX, progressBarY *widget.ProgressBar, solutionLabel, progressLevelX, progressLevelY *widget.Label, solution [][]int, first, second int) Animation {
	return &animator{
		SolutionLabel:  solutionLabel,
		ProgressBarX:   progressBarX,
		ProgressBarY:   progressBarY,
		ProgressLevelX: progressLevelX,
		ProgressLevelY: progressLevelY,
		Solution:       solution,
		First:          first,
		Second:         second,
	}
}

func (a *animator) Animate() {
	solution := solver.ConvertSolution(a.Solution, a.First, a.Second)

	a.ProgressBarX.SetValue(0.0)
	a.ProgressBarY.SetValue(0.0)

	for i, step := range solution {

		a.ProgressLevelX.Text = fmt.Sprintf("%v gal Jug:  %v", a.First, a.Solution[i][0])
		a.ProgressLevelY.Text = fmt.Sprintf("%v gal Jug:  %v", a.Second, a.Solution[i][1])
		a.SolutionLabel.Text = solver.TranslateProcess(a.Solution[i][2])

		valA := a.ProgressBarX.Value
		valB := a.ProgressBarY.Value
		incA := 0.1
		incB := 0.1
		if IsLessThan(step[0], valA) {
			incA = -0.1
		}

		if IsLessThan(step[1], valB) {
			incB = -0.1
		}

		for !(IsEqual(step[0], valA) && IsEqual(step[1], valB)) {
			time.Sleep(time.Millisecond * 250)

			if !IsEqual(step[0], valA) {
				valA += incA
				a.ProgressBarX.SetValue(valA)
			}
			if !IsEqual(step[1], valB) {
				valB += incB
				a.ProgressBarY.SetValue(valB)
			}
		}
	}

	a.Button.Enable()
}

func (a *animator) SetButton(button *widget.Button) {
	a.Button = button
}

func IsEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.1
}

func IsLessThan(a, b float64) bool {
	return a < b
}
