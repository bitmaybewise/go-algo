package algo

func Factorial(n int) int {
	var factorial func(n, acc int) int
	factorial = func(n, acc int) int {
		if n <= 1 {
			return acc
		}
		return factorial(n-1, n*acc)
	}
	return factorial(n, 1)
}
