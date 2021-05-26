package backtrack

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

func GenerateSubgroups(n int) [][]int {
	s := newSubgroups(n)
	Backtrack(&s, s.temporaryCombination, 0, s.CombinationLength)
	return s.Results
}
