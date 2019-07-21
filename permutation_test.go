package algogo

import (
	"testing"
)

type permutationFunc func(a []int) bool

func testPermutation(t *testing.T, permFunc permutationFunc, name string, in []int) {
	out := Copied(in)
	outs := [][]int{Copied(in)}

	nperm := 1
	for permFunc(out) {
		outs = append(outs, Copied(out))
		nperm++
	}

	// check number of permutations generated
	nperm_e := Factorial(len(in))
	if nperm != nperm_e {
		t.Errorf("%s(%v) generated %d permutations, expected %d", name, in, nperm, nperm_e)
	}

	// check whether the final state of the modified array is consistent
	if Different(out, in) {
		t.Errorf("%s(%v) generated %v on the final iteration, expected %v",
			name, in, out, in )
	}

	// check whether all permutations are unique
	for i := 0; i < len(outs); i++ {
		for j := 0; j < len(outs); j++ {
			if i != j && Same(outs[i], outs[j]) {
				t.Errorf("%s(%v) generated duplicate permutations: %v, %v", name, in, outs[i], outs[j])
			}
		}
	}
}

func TestNextPermutation(t *testing.T) {
	in := []int{1, 2, 3, 4}
	testPermutation(t, NextPermutation, "NextPermutation", in)
}

func TestPrevPermutation(t *testing.T) {
	in := []int{4, 3, 2, 1}
	testPermutation(t, PrevPermutation, "PrevPermutation", in)
}

func TestNextPrevPermutation(t *testing.T) {
	// NextPermutation and PrevPermutation must be inverse functions
	in := []int{4, 2, 9, 3}

	nperm := Factorial(len(in))
	for i := 0; i < nperm; i++ {
		out := Copied(in)
		NextPermutation(out)
		PrevPermutation(out)
		if Different(in, out) {
			t.Errorf("PrevPermutation(%v); NextPermutation(%v) => %v, expected %v", in, in, out, in)
		}
		NextPermutation(in)
	}

	for i := 0; i < nperm; i++ {
		out := Copied(in)
		PrevPermutation(out)
		NextPermutation(out)
		if Different(in, out) {
			t.Errorf("NextPermutation(%v); PrevPermutation(%v) => %v, expected %v", in, in, out, in)
		}
		PrevPermutation(in)
	}

}

