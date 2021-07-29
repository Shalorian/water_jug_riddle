package solver

const (
	Fill = iota
	Empty
	Transfer
)

func Solve(fromCap, toCap, target int) (steps [][]int) {
	from := fromCap
	to := 0

	steps = append(steps, []int{from, to, Fill})

	for !HasReachedTarget(from, to, target) {
		pour := Min(from, toCap-to)

		to = to + pour
		from = from - pour

		steps = append(steps, []int{from, to, Transfer})

		if HasReachedTarget(from, to, target) {
			return steps
		}

		if from == 0 {
			from = fromCap
			steps = append(steps, []int{from, to, Fill})
		}

		if to == toCap {
			to = 0
			steps = append(steps, []int{from, to, Empty})
		}
	}

	return steps
}

func ConvertSolution(solution [][]int, first, second int) (converted [][]float64) {
	for _, x := range solution {
		converted = append(converted, []float64{(float64(x[0]) / float64(first)), (float64(x[1]) / float64(second))})
	}
	return converted
}

func TranslateProcess(c int) string {
	return []string{"Fill", "Empty", "Transfer"}[c]
}
