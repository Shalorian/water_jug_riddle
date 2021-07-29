package solver

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func IsSolvable(a, b, c int) bool {
	if c > a && c > b {
		return false
	}
	return c%GCD(a, b) == 0
}

func HasReachedTarget(a, b, target int) bool {
	return a == target || b == target
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetBestSolution(a, b, target int) ([][]int, int, int) {
	x := Solve(a, b, target)
	y := Solve(b, a, target)

	if len(x) > len(y) {
		return y, b, a
	}

	return x, a, b
}
