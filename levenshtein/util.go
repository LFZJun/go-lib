package levenshtein

func MinOfThree(a, b, c int) int {
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

func Min(args ...int) int {
	length := len(args)
	for i := length - 1; i >= 0; i-- {
		if args[i-1] > args[i] {
			args[i-1], args[i] = args[i], args[i-1]
		}
	}
	return args[0]
}
