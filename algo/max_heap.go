package algo

type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{data: make([]int, 0)}
}

func (h *MaxHeap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) Insert(value int) {
	h.data = append(h.data, value)
	newNodeIndex := len(h.data) - 1

	for newNodeIndex > 0 && h.data[newNodeIndex] > h.data[h.parentIndex(newNodeIndex)] {
		h.data[h.parentIndex(newNodeIndex)], h.data[newNodeIndex] = h.data[newNodeIndex], h.data[h.parentIndex(newNodeIndex)]
		newNodeIndex = h.parentIndex(newNodeIndex)
	}
}

func (h *MaxHeap) Root() int {
	return h.data[0]
}

func (h *MaxHeap) leftChildIndexOf(index int) int {
	return index*2 + 1
}

func (h *MaxHeap) rightChildIndexOf(index int) int {
	return index*2 + 2
}

func (h *MaxHeap) hasGreaterChild(index int) bool {
	leftIndex := h.leftChildIndexOf(index)
	rightIndex := h.rightChildIndexOf(index)
	limit := len(h.data) - 1
	return (leftIndex >= 0 && leftIndex <= limit && h.data[leftIndex] > h.data[index]) ||
		(rightIndex >= 0 && rightIndex <= limit && h.data[rightIndex] > h.data[index])
}

func (h *MaxHeap) calculateLargerChildIndex(index int) int {
	if h.rightChildIndexOf(index) < 0 {
		return h.leftChildIndexOf(index)
	}
	if h.data[h.rightChildIndexOf(index)] > h.data[h.leftChildIndexOf(index)] {
		return h.rightChildIndexOf(index)
	}
	return h.leftChildIndexOf(index)
}

func (h *MaxHeap) Pop() int {
	value := h.Root()

	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[0 : len(h.data)-1]
	trickleNodeIndex := 0

	for h.hasGreaterChild(trickleNodeIndex) {
		largerChildIndex := h.calculateLargerChildIndex(trickleNodeIndex)
		h.data[trickleNodeIndex], h.data[largerChildIndex] = h.data[largerChildIndex], h.data[trickleNodeIndex]
		trickleNodeIndex = largerChildIndex
	}

	return value
}

func (h *MaxHeap) Last() int {
	return h.data[len(h.data)-1]
}
