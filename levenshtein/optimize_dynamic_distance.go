package levenshtein

func optimizeDynamicDistance(s1, s2 []rune) int {
	var cost, lastdiag, olddiag int

	lenS1 := len(s1)
	lenS2 := len(s2)

	column := make([]int, lenS1+1)

	for y := 1; y <= lenS1; y++ {
		column[y] = y
	}

	for x := 1; x <= lenS2; x++ {
		column[0] = x
		lastdiag = x - 1
		for y := 1; y <= lenS1; y++ {
			olddiag = column[y]
			cost = 0
			if s1[y-1] != s2[x-1] {
				cost = 1
			}
			column[y] = MinOfThree(
				column[y]+1,
				column[y-1]+1,
				lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return column[lenS1]
}

func OptimizeDynamicDistance(str1, str2 string) int {
	return optimizeDynamicDistance([]rune(str1), []rune(str2))
}
