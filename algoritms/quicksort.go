package algoritms

func QuickSort(slice []int, start int, end int, comparer func(int, int) bool) {
	if start >= end {
		return
	}
	pivot := partition(slice, start, end, comparer)
	QuickSort(slice, start, pivot-1, comparer)
	QuickSort(slice, pivot+1, end, comparer)
}
func partition(slice []int, start int, end int, comparer func(int, int) bool) int {
	marker := start
	for i := start; i < end; i++ {
		if comparer(slice[i], slice[end]) {
			slice[marker], slice[i] = slice[i], slice[marker]
			marker++
		}
	}
	slice[marker], slice[end] = slice[end], slice[marker]
	return marker
}
