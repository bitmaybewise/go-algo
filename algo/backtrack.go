package algo

type Backtrackable interface {
	ConstructCandidate(temporaryCombination []int, currentIndex, combinationLength int) []int
	IsSolution(temporaryCombination []int, currentIndex, combinationLength int) bool
	ProcessSolution(aCombination []int, currentIndex int)
}

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

type subgroups struct {
	CombinationLength    int
	Results              [][]int
	temporaryCombination []int
}

func (p *subgroups) IsSolution(_ []int, currentIndex, combinationLength int) bool {
	return currentIndex == combinationLength
}

func (p *subgroups) ProcessSolution(aCombination []int, currentIndex int) {
	newCombination := make([]int, 0)
	for i := 1; i <= currentIndex; i++ {
		if aCombination[i] == 1 {
			newCombination = append(newCombination, i)
		}
	}
	p.Results = append(p.Results, newCombination)
}

func (p *subgroups) ConstructCandidate(temporaryCombination []int, currentIndex, combinationLength int) []int {
	return []int{1, 0}
}

func newSubgroups(n int) subgroups {
	return subgroups{
		CombinationLength:    n,
		Results:              make([][]int, 0),
		temporaryCombination: make([]int, n+1),
	}
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

func GeneratePermutations(n int) [][]int {
	p := newPermutation(n)
	Backtrack(&p, p.temporaryCombination, 0, p.CombinationLength)
	return p.Results
}

func GenerateSubsets(n int) [][]int {
	s := newSubgroups(n)
	Backtrack(&s, s.temporaryCombination, 0, s.CombinationLength)
	return s.Results
}
