package algo

func Permutations(n int) [][]int {
	results := make([][]int, 0)
	appendPermutation := func(a []int, k int) {
		p := make([]int, 0)
		for i := 1; i <= k; i++ {
			p = append(p, a[i])
		}
		results = append(results, p)
	}

	a := make([]int, n+1)
	backtrack(a, 0, n, appendPermutation)
	return results
}

type processSolution func([]int, int)

func backtrack(a []int, k, n int, solveOne processSolution) {
	if k == n {
		solveOne(a, k)
		return
	}

	k++
	candidates := constructCandidates(a, k, n)
	for i := 0; i < len(candidates); i++ {
		a[k] = candidates[i]
		backtrack(a, k, n, solveOne)
	}
}

func constructCandidates(a []int, k, n int) []int {
	in_perm := make(map[int]bool)
	for i := 1; i < k; i++ {
		in_perm[a[i]] = true
	}

	candidates := make([]int, 0)
	for i := 1; i <= n; i++ {
		if _, ok := in_perm[i]; !ok {
			candidates = append(candidates, i)
		}
	}

	return candidates
}
