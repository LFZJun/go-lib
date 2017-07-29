package search

func IndexOf(src, substring string) []int {
	s, ss := []rune(src), []rune(substring)
	ls, lss := len(s), len(ss)
	if ls|lss == 0 {
		return []int{0, 0}
	}
	end := ls - lss + 1
	for i := 0; i < end; i++ {
		var t int
		for ii, rr := range ss {
			if s[i+ii] != rr {
				break
			}
			t++
		}
		if t == lss {
			return []int{i, i + lss}
		}
	}
	return nil
}
