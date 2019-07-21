package bit

import (
	"testing"
)

var countTests = []struct {
	in  string
	out int
}{
	{"0", 0},
	{"1", 1},
	{"2", 1},
	{"3", 2},
	{"4", 1},
	{"1025", 2},
	{"1026", 2},
	{"0xabc", 7},
	{"-0x8", 1},
	{"0x80", 1},
	{"0x800", 1},
	{"0x811", 3},
	{"0xfff", 12},
}

func TestCount(t *testing.T) {
	for i, test := range countTests {
		x := NewArray(0)
		_, ok := x.SetString(test.in, 0)
		if !ok {
			t.Errorf("#%d test input invalid: %s", i, test.in)
			continue
		}

		if n := x.Count(); n != test.out {
			t.Errorf("#%d got %d want %d", i, n, test.out)
		}
	}
}

func BenchmarkCountBits(b *testing.B) {
	a := NewArray(1)
	a.Lsh(a.Int, 1000)
	for i := 0; i < b.N; i++ {
		countBits(a)
	}
}

func BenchmarkXorBits(b *testing.B) {
	a := NewArray(1)
	a.Lsh(a.Int, 1000)
	z := NewArray(0).Set(a)
	z.Rsh(z.Int, 1)
	for i := 0; i < b.N; i++ {
		xorBits(a, z)
	}
}

func BenchmarkXorBitsInPlace(b *testing.B) {
	a := NewArray(1)
	a.Lsh(a.Int, 1000)
	z := NewArray(0).Set(a)
	z.Rsh(z.Int, 1)
	for i := 0; i < b.N; i++ {
		xorBitsInPlace(a, z)
	}
}

func BenchmarkCountBools(b *testing.B) {
	a := make([]bool, 1001)
	a[1000] = true
	for i := 0; i < b.N; i++ {
		countBools(a)
	}
}

func BenchmarkXorBools(b *testing.B) {
	a := make([]bool, 1001)
	a[1000] = true
	z := make([]bool, 1001)
	z[999] = true
	for i := 0; i < b.N; i++ {
		xorBools(a, z)
	}
}

func BenchmarkXorBoolsInPlace(b *testing.B) {
	a := make([]bool, 1001)
	a[1000] = true
	z := make([]bool, 1001)
	z[999] = true
	for i := 0; i < b.N; i++ {
		xorBoolsInPlace(a, z)
	}
}
