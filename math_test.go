package algogo

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	in := []int{0, 1, 2, 3, 4, 5}
	ans := []int{1, 1, 2, 6, 24, 120}
	for i, v := range in {
		if Factorial(v) != ans[i] {
			t.Errorf("Factorial(%d) == %d, expected %d", v, Factorial(v), ans[i])
		}
	}
}

func TestPermutations(t *testing.T) {
	// Permutation(n, n) == 1 for all n
	// Permutation(n, 1) == n for all n
	for i := 1; i < 10; i++ {
		if Permutations(i, i) != Factorial(i) {
			t.Errorf("Permutations(%d, %d) == %d, expected %d", i, i, Permutations(i, i), Factorial(1))
		}
		if Permutations(i, 1) != i {
			t.Errorf("Permutations(%d, %d) == %d, expected %d", i, 1, Permutations(i, 1), i)
		}
	}

	if Permutations(6, 3) != 120 ||
		 Permutations(6, 4) != 360 ||
		 Permutations(4, 2) != 12 {
		 t.Errorf("Permutations(...) yielded an unexpected result")
	}
}

