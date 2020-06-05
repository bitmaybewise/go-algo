package algo

func InsertionSort(collection []int) {
	size := len(collection)
	for i := 1; i < size; i++ {
		current := collection[i]
		previousIdx := i - 1

		for previousIdx >= 0 && collection[previousIdx] > current {
			collection[previousIdx+1] = collection[previousIdx]
			previousIdx--
		}

		collection[previousIdx+1] = current
	}
}
