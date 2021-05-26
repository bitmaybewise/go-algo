package sorting

func SelectionSort(collection []int) {
	size := len(collection)
	for i := 0; i < size-1; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if collection[j] < collection[min] {
				min = j
			}
		}
		collection[i], collection[min] = collection[min], collection[i]
	}
}
