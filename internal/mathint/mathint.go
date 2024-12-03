package mathint

func AbsInt(a, b int) int {
	return MaxInt(a, b) - MinInt(a, b)
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
