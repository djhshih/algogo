package algogo

// NextPermutation modifies array a in place to the next permutation that is
// lexicographically greater.
// It returns false when the permutation is at the lexicographical minimum
// (i.e. sorted in nondecreasing order).
// It returns false no permutation is possible.
// It returns true otherwise.
// To iterate through all permutations of elements of a through repeated calls,
// a must be initialized at the lexicographical minimum;
// i.e. a should be sorted in non-decreasing order.
// Translated from C++ std <algorithm> library
func NextPermutation(a []int) bool {
	i := len(a)
	if i <= 1 {
		// len is 0 or 1: no permutation possible
		return false
	}

	// enter loop with i at the last index
	i--
	for {
		// loop from the end until first decreasing element in reverse order
		// effectively, the array is split into head and tail, where the tail is non-increasing
		ii := i
		i--
		if a[i] < a[ii] {
			// now i indexes the last element of the head
			// reverse-find the min element of the tail that is greater than the a[i]
			j := len(a) - 1
			for ; !(a[i] < a[j]); j-- {
			}
			// swap the two elements, s.t. the tail can begin another cycle to achieve the lexicographical maximum
			a[i], a[j] = a[j], a[i]
			// reverse the tail to begin another cycle
			Reverse(a[ii:])
			return true
		}
		if i == 0 {
			// non-increasing tail spans entire array; a is at lexicographical maximum:
			// reverse a to return to the lexicographical minimum
			Reverse(a)
			return false
		}
	}
	return false
}

// NextKPermutation modifies array a in place to the next k-permutation;
// a[0:k] will store the next k-permutation of the n elements in a.
// Adapted from C++ std proposal WG21/N2639, Herve Bronnimann
func NextKPermutation(a []int, k int) bool {
	Reverse(a[k:])
	return NextPermutation(a)
}

// NextPermutation modifies array a in place to the previous permutation that
// is lexicographically smaller.
// To iterate through all permutations of elements of a by repeated calls,
// a must be initialized at the lexicographical minimum;
// i.e. a should be sorted in non-increasing order.
// Translated from C++ std <algorithm> library
func PrevPermutation(a []int) bool {
	i := len(a)
	if i <= 1 {
		// len is 0 or 1: no permutation possible
		return false
	}

	// enter loop with i at the last index
	i--
	for {
		// loop from the end until first increasing element in reverse order
		// effectively, the array is split into head and tail, where the tail is non-decreasing
		ii := i
		i--
		if a[i] > a[ii] {
			// now i indexes the last element of the head
			// reverse-find the max element of the tail that is less than the a[i]
			j := len(a) - 1
			for ; !(a[i] > a[j]); j-- {
			}
			// swap the two elements, s.t. the tail can begin another cycle to achieve the lexicographical minimum
			a[i], a[j] = a[j], a[i]
			// reverse the tail to begin another cycle
			Reverse(a[ii:])
			return true
		}
		if i == 0 {
			// non-decreasing tail spans entire array; a is at lexicographical minimum:
			// reverse a to return to lexicographical maximum
			Reverse(a)
			return false
		}
	}
	return false
}

// PrevKPermutation modifies array a in place to the previous k-permutation;
// a[0:k] will store the previous k-permutation of the n elements in a.
// Adapted from C++ std proposal WG21/N2639, Herve Bronnimann
func PrevKPermutation(a []int, k int) bool {
	result := PrevPermutation(a)
	Reverse(a[k:])
	return result
}
