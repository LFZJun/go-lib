package levenshtein

func dynamicDistance(a, b []rune) int {
	lenY, lenX := len(a)+1, len(b)+1
	d := make([][]int, lenY)
	for t := 0; t < lenY; t++ {
		d[t] = make([]int, lenX)
	}
	var y, x int
	for y = 0; y < lenY; y++ {
		d[y][0] = y
	}
	for x = 0; x < lenX; x++ {
		d[0][x] = x
	}

	for y = 1; y < lenY; y++ {
		for x = 1; x < lenX; x++ {
			if a[y-1] == b[x-1] {
				d[y][x] = d[y-1][x-1]
			} else {
				d[y][x] = MinOfThree(d[y-1][x]+1, d[y][x-1]+1, d[y-1][x-1]+1)
			}
		}
	}
	return d[len(a)][len(b)]
}

func DynamicDistance(a1, a2 string) int {
	return dynamicDistance([]rune(a1), []rune(a2))
}
