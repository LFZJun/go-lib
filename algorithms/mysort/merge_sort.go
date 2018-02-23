package mysort

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(append(result, left[l:]...), right[r:]...)
	return
}

func mergeSort(l []int) []int {
	if len(l) <= 1 {
		return l
	}
	mid := len(l) / 2
	return merge(mergeSort(l[:mid]), mergeSort(l[mid:]))
}
