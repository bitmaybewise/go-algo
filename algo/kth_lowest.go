package algo

func KthLowest(collection []int, kth int) int {
	return quickSelect(collection, kth, 0, len(collection)-1)
}

func quickSelect(collection []int, kthLowestValue, leftIdx, rightIdx int) int {
	if (rightIdx - leftIdx) <= 0 {
		return collection[leftIdx]
	}
	pivotIdx := quickselectPartition(collection, leftIdx, rightIdx)
	if kthLowestValue < pivotIdx {
		return quickSelect(collection, kthLowestValue, leftIdx, pivotIdx-1)
	}
	if kthLowestValue > pivotIdx {
		return quickSelect(collection, kthLowestValue, pivotIdx+1, rightIdx)
	}
	return collection[pivotIdx]
}

func quickselectPartition(collection []int, leftIdx, rightIdx int) int {
	pivotIdx := rightIdx
	pivot := collection[pivotIdx]
	rightIdx--
	for {
		for collection[leftIdx] < pivot {
			leftIdx++
		}
		for rightIdx >= 0 && collection[rightIdx] > pivot {
			rightIdx--
		}
		if leftIdx >= rightIdx {
			break
		}
		collection[leftIdx], collection[rightIdx] = collection[rightIdx], collection[leftIdx]
		leftIdx++
	}
	collection[leftIdx], collection[pivotIdx] = collection[pivotIdx], collection[leftIdx]
	return leftIdx
}
