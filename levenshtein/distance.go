package levenshtein

func Distance(str1, str2 string) int {
	var cost, lastdiag, olddiag int
	s1 := []rune(str1)
	s2 := []rune(str2)

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
			column[y] = Min(
				column[y]+1,
				column[y-1]+1,
				lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return column[lenS1]
}

func Min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
