package algogo

import (
	"testing"
)

func TestReverse(t *testing.T) {
	in := []int{1, 2, 3, 4}
	out := make([]int, len(in))
	copy(out, in)
	ans := []int{4, 3, 2, 1}
	Reverse(out)
	different := false
	for i, _ := range ans {
		if out[i] != ans[i] {
			different = true
			break
		}
	}
	if different {
		t.Errorf("Reverse(%v) == %v, expected %v", in, out, ans)
	}
}

