package mysort

func quickSortV1(l []int) []int {
	if len(l) < 2 {
		return l
	}
	pivot := l[0]
	var low, equal, high []int
	for _, v := range l {
		if v < pivot {
			low = append(low, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			high = append(high, v)
		}
	}
	return append(append(quickSortV1(low), equal...), high...)
}

func quickSortV2(l []int, start, end int) []int {
	if end <= start {
		return l
	}
	pivot := l[end]
	t := start
	for i := start; i < end; i++ {
		if l[i] < pivot {
			l[i], l[t] = l[t], l[i]
			t++
		}
	}
	l[t], l[end] = l[end], l[t]
	quickSortV2(l, start, t-1)
	quickSortV2(l, t+1, end)
	return l
}

func quickSort(l []int) {
	if len(l) < 2 {
		return
	}
	length := len(l)
	index := 0
	end := length - 1
	pivot := l[end]
	for i := 0; i < end; i++ {
		if l[i] < pivot {
			l[i], l[index] = l[index], l[i]
			index++
		}
	}
	l[index], l[end] = l[end], l[index]
	quickSort(l[:index])
	quickSort(l[index+1:])
}
