package algo

type Backtrackable interface {
	ConstructCandidate(temporaryCombination []int, currentIndex, combinationLength int) []int
	IsSolution(temporaryCombination []int, currentIndex, combinationLength int) bool
	ProcessSolution(aCombination []int, currentIndex int)
}

func Backtrack(b Backtrackable, temporaryCombination []int, currentIndex, length int) {
	if b.IsSolution(temporaryCombination, currentIndex, length) {
		b.ProcessSolution(temporaryCombination, currentIndex)
		return
	}

	currentIndex++
	candidates := b.ConstructCandidate(temporaryCombination, currentIndex, length)
	for i := 0; i < len(candidates); i++ {
		temporaryCombination[currentIndex] = candidates[i]
		Backtrack(b, temporaryCombination, currentIndex, length)
	}
}
