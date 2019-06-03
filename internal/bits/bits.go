// Package bits provides an alternative implementation of bits.Len64 from math/bits.
package bits

// Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
//
// Alternative implementation of bits.Len64 from math/bits, optimized for minimal type conversions
func Len64(x uint64) (n uint) {
	if x >= 1<<32 {
		x >>= 32
		n = 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return n + len8tab[x]
}
