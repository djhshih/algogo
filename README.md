# Algorithms in Go

Package algogo provides basic algorithms not in the Go standard library.

- Array utilities (e.g. Copy, Reverse, Different, Same) for integer array
- Combinations for integer array
- Permutations for integer array

## Subdirectories

- bit: Package bit implements a bit array (extends math/big.Int). Package bit provides bit operations (implemented by math/big.Int) and a bit count method for counting the number of set bits. Using an integer array decreases run time by an order of magnitude compared to using an boolean array.

- unionfind: Package unionfind implements quick-union and find with path compression for integer set.
Additional algorithms and data structures will be implemented in the future.

Most functions are implemented for integer arrays, which are likely the most useful in a world without generics. Each data structure can contain integer values serving as indices into an array in which the actual data are stored. The interface{} type will not be used, it subverts the type system and requires runtime checks.
