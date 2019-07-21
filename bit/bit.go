package bit

import (
	"math/big"
)

// A bit array.
type Array struct {
	*big.Int
}

// NewArray allocates and returns a new Bits set to x.
func NewArray(x int64) *Array {
	return &Array{big.NewInt(x)}
}

// Count counts the number of set bits in b.
func (b *Array) Count() (c int) {
	for _, v := range b.Bits() {
		c += nSetBits(uintptr(v))
	}
	return
}

// Len returns the length of b in bits. The bit length of 0 is 0.
func (b *Array) Len() int {
	return b.Int.BitLen()
}

// Set sets b to x and returns b.
func (b *Array) Set(x *Array) *Array {
	b.Int.Set(x.Int)
	return b
}

// SetBit sets b to x with x's i'th bit set to a (0 or 1).
func (b *Array) SetBit(x *Array, i int, a uint) *Array {
	b.Int.SetBit(x.Int, i, a)
	return b
}

// SetString sets b to the value of s, interpreted in the given base, and
// returns b and a boolean indicating success. If SetString fails, the value
// of b is undefined but he returned value is nil.
func (b *Array) SetString(s string, base int) (*Array, bool) {
	_, succ := b.Int.SetString(s, base)
	return b, succ
}

// Not sets b = ^x and returns b.
func (b *Array) Not(x *Array) *Array {
	b.Int.Not(x.Int)
	return b
}

// Or sets b = x | y and returns b.
func (b *Array) Or(x, y *Array) *Array {
	b.Int.Or(x.Int, y.Int)
	return b
}

// And sets b = x & y and returns b.
func (b *Array) And(x, y *Array) *Array {
	b.Int.And(x.Int, y.Int)
	return b
}

// AndNot sets b = x & ^y and returns b.
func (b *Array) AndNot(x, y *Array) *Array {
	b.Int.AndNot(x.Int, y.Int)
	return b
}

// AndNot sets b = x ^ y and returns b.
func (b *Array) Xor(x, y *Array) *Array {
	b.Int.Xor(x.Int, y.Int)
	return b
}

// Count set bits using the Wegner algorithm.
// Adapted from Wegner, CACM 3 (1960), p. 322.
func nSetBits(x uintptr) (n int) {
	// n accumulates the total bits set in x, counting only set bits
	for ; x > 0; n++ {
		// clear the least significant bit set
		x &= x - 1
	}
	return
}

func countBits(a *Array) int {
	return a.Count()
}

func countBools(a []bool) (c int) {
	for _, v := range a {
		if v {
			c++
		}
	}
	return
}

func xorBitsInPlace(a, b *Array) *Array {
	return a.Xor(a, b)
}

func xorBits(a, b *Array) (c *Array) {
	c = NewArray(0)
	c.Xor(a, b)
	return
}

func xorBoolsInPlace(a, b []bool) []bool {
	for i, v := range a {
		if (v || b[i]) && !(v && b[i]) {
			a[i] = true
		}
	}
	return a
}

func xorBools(a, b []bool) (d []bool) {
	d = make([]bool, len(a))
	for i, v := range a {
		if (v || b[i]) && !(v && b[i]) {
			d[i] = true
		}
	}
	return
}
