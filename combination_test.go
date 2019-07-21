package algogo

import (
	"fmt"
	"testing"
)

// TODO Check test results...

func TestCombination(t *testing.T) {
	in := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 2
	fmt.Println("Next K")
	fmt.Println(in[:k], in[k:])
	n := 1
	for NextCombination(in, k) {
		fmt.Println(in[:k], in[k:])
		n++
	}
	fmt.Println("Found ", n)

	PrevCombination(in, k)
	fmt.Println("Prev K")
	fmt.Println(in[:k], in[k:])
	n = 1
	for PrevCombination(in, k) {
		fmt.Println(in[:k], in[k:])
		n++
	}
	fmt.Println("Found ", n)
}
