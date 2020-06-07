package algo

func QuickSort(collection []int) {
	if len(collection) > 1 {
		pivotIdx := partition(collection)
		QuickSort(collection[0:pivotIdx])
		QuickSort(collection[pivotIdx+1:])
	}
}

func partition(col []int) int {
	pivotIdx := len(col) - 1
	i := 0
	for j := range col[0:pivotIdx] {
		if col[j] <= col[pivotIdx] {
			col[j], col[i] = col[i], col[j]
			i++
		}
	}
	col[i], col[pivotIdx] = col[pivotIdx], col[i]
	return i
}
