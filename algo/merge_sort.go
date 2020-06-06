package algo

func MergeSort(collection []int) {
	size := len(collection)
	if size > 1 {
		middle := size / 2
		MergeSort(collection[0:middle])
		MergeSort(collection[middle:])
		merge(collection, middle)
	}
}

func merge(col []int, middle int) {
	head, tail := col[0:middle], col[middle:]

	left, right := make([]int, len(head)), make([]int, len(tail))
	copy(left, head)
	copy(right, tail)
	topLeft, topRight := 0, 0

	for i := range col {
		switch {
		case topLeft >= len(left):
			col[i] = right[topRight]
			topRight++
		case topRight >= len(right):
			col[i] = left[topLeft]
			topLeft++
		case left[topLeft] <= right[topRight]:
			col[i] = left[topLeft]
			topLeft++
		case left[topLeft] >= right[topRight]:
			col[i] = right[topRight]
			topRight++
		}
	}
}
