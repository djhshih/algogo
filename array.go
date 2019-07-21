package algogo

// Reverse reverses the elements of array a
func Reverse(a []int) {
	last := len(a)
	mid := last / 2
	last--
	for i := 0; i < mid; i++ {
		j := last - i
		a[i], a[j] = a[j], a[i]
	}
}

// Reversed returns a reversed copy of array a
func Reversed(a []int) []int {
	b := Copied(a)
	Reverse(b)
	return b
}

// Copied returns a copy of array a
func Copied(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

// Different returns whether arrays a and b are different
func Different(a, b []int) (different bool) {
	for i, v := range a {
		if v != b[i] {
			different = true
			break
		}
	}
	return
}

// Same returns whether arrays a and b store the same values
func Same(a, b []int) (same bool) {
	return !Different(a, b)
}
