package sorting

func BinarySearch(collection []int, value int) int {
	return binary_search(collection, value, 0, len(collection)-1)
}

func binary_search(collection []int, value int, firstIdx int, lastIdx int) int {
	if firstIdx <= lastIdx {
		middle := (lastIdx + firstIdx) / 2
		if collection[middle] == value {
			return middle
		}
		if collection[middle] > value {
			return binary_search(collection, value, 0, middle-1)
		}
		return binary_search(collection, value, middle+1, lastIdx)
	}
	return -1
}
