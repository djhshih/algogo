package algogo

// Factorial returns n!.
func Factorial(n int) (f int) {
	if n <= 1 {
		return 1
	}
	f = 1
	for i := n; i > 1; i-- {
		f *= i
	}
	return
}

// Permutations returns the number of k-permutations of n.
func Permutations(n, k int) (p int) {
	if k > n {
		return 0
	}
	p = 1
	m := n - k
	for i := n; i > m; i-- {
		p *= i
	}
	return
}

// Combinations returns the number k-combinations of n.
func Combinations(n, k int) (c int) {
	return Permutations(n, k) / Factorial(k)
}
