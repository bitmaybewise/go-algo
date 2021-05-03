package algo

type permutation struct {
	CombinationLength    int
	Results              [][]int
	temporaryCombination []int
}

func (p *permutation) IsSolution(_ []int, currentIndex, combinationLength int) bool {
	return currentIndex == combinationLength
}

func (p *permutation) appendCombination(values ...int) {
	newCombination := make([]int, 0)
	newCombination = append(newCombination, values...)
	p.Results = append(p.Results, newCombination)
}

func (p *permutation) ProcessSolution(aCombination []int, _ int) {
	p.appendCombination(aCombination[1:]...)
}

func (p *permutation) ConstructCandidate(temporaryCombination []int, currentIndex, combinationLength int) []int {
	in_perm := make(map[int]bool)
	for i := 1; i < currentIndex; i++ {
		in_perm[temporaryCombination[i]] = true
	}

	candidates := make([]int, 0)
	for i := 1; i <= combinationLength; i++ {
		if _, ok := in_perm[i]; !ok {
			candidates = append(candidates, i)
		}
	}

	return candidates
}

func newPermutation(n int) permutation {
	return permutation{
		CombinationLength:    n,
		Results:              make([][]int, 0),
		temporaryCombination: make([]int, n+1),
	}
}

func GeneratePermutations(n int) [][]int {
	p := newPermutation(n)
	Backtrack(&p, p.temporaryCombination, 0, p.CombinationLength)
	return p.Results
}
