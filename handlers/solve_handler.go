package handlers

import (
	"errors"
	"strconv"

	"simon.alvarez/water_jug_riddle/solver"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	errZeroVal      = errors.New("You should use values different from 0!")
	errInvalidInput = errors.New("You have to use valid inputs!")
	errNotSolvable  = errors.New("Not Solvable")
)

type SolveHandler interface {
	Handle() ([][]int, int, int, int, error)
}

type solveHandler struct {
	App    fyne.App
	Window fyne.Window
	Inputs []*widget.Entry
}

func NewSolveHandler(app fyne.App, window fyne.Window, inputs []*widget.Entry) SolveHandler {
	return &solveHandler{
		App:    app,
		Window: window,
		Inputs: inputs,
	}
}

func (s *solveHandler) Handle() ([][]int, int, int, int, error) {
	jugX, errX := strconv.Atoi(s.Inputs[0].Text)
	jugY, errY := strconv.Atoi(s.Inputs[1].Text)
	measure, errM := strconv.Atoi(s.Inputs[2].Text)

	if errX != nil || errY != nil || errM != nil {
		dialog.ShowError(errInvalidInput, s.Window)
		return [][]int{}, 0, 0, 0, errInvalidInput
	}

	if jugX == 0 || jugY == 0 || measure == 0 {
		dialog.ShowError(errZeroVal, s.Window)
		return [][]int{}, 0, 0, 0, errZeroVal
	}

	if !solver.IsSolvable(jugX, jugY, measure) {
		dialog.ShowError(errNotSolvable, s.Window)
		return [][]int{}, 0, 0, 0, errNotSolvable
	}

	solution, first, second := solver.GetBestSolution(jugX, jugY, measure)

	return solution, first, second, measure, nil
}
