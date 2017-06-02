package levenshtein

func distanceDivide(a, b []rune, i, j int) int {
	if j == 0 {
		return i
	}
	if i == 0 {
		return j
	}
	if a[i-1] == b[j-1] {
		return distanceDivide(a, b, i-1, j-1)
	}
	return Min(distanceDivide(a, b, i-1, j)+1,
		distanceDivide(a, b, i, j-1)+1,
		distanceDivide(a, b, i-1, j-1)+1,
	)
}

func DistanceDivide(a, b string) int {
	a1, a2 := []rune(a), []rune(b)
	return distanceDivide(a1, a2, len(a1), len(a2))
}
