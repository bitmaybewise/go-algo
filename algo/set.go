package algo

type IntSet struct {
	values map[int]bool
}

func NewSet() *IntSet {
	return &IntSet{values: make(map[int]bool)}
}

func SetFromValues(values ...int) *IntSet {
	set := NewSet()
	for _, v := range values {
		set.Add(v)
	}
	return set
}

func (s *IntSet) Add(n int) {
	s.values[n] = true
}

func (s *IntSet) Size() int {
	return len(s.values)
}

func (s *IntSet) Values() []int {
	v := make([]int, len(s.values))
	i := 0
	for k := range s.values {
		v[i] = k
		i++
	}
	return v
}

func (s *IntSet) Contains(numbers ...int) bool {
	for k := range numbers {
		if s.values[k] {
			continue
		}
		return false
	}
	return true
}
