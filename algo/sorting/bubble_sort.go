package sorting

func BubbleSort(collection []int) {
	size := len(collection)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1; j++ {
			if collection[j] > collection[j+1] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
			}
		}
	}
}
