package search

func BinarySearch(array []int, key int) int {
	low, high := 0, len(array)-1
	for low <= high {
		mid := (low + high) / 2
		pivot := array[mid]
		switch {
		case key < pivot:
			high = mid - 1
		case key == pivot:
			return mid
		default:
			low = mid + 1
		}
	}
	return -1
}
