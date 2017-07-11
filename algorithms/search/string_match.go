package search

func IndexOf(src, substring string) []int {
	s, ss := []rune(src), []rune(substring)
	ls, lss := len(s), len(ss)
	for i := range s {
		var t int
		end := i + lss
		if end > ls {
			return nil
		}
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
