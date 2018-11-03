// Package wide provides implementations of Int128 and Uint128 for Go. See readme.md for more information.
package wide

import (
	"fmt"
	"math/big"
)

// Int128 is a representation of a signed 128-bit integer
type Int128 struct {
	hi int64
	lo uint64
}

// String returns a hexadecimal (string) representation of a Int128
func (x Int128) String() string {
	switch {
	case x.hi < 0:
		x = x.Neg()
		if x.hi == 0 {
			return fmt.Sprintf("-%#x", x.lo) // ignore leading 0's
		}
		return fmt.Sprintf("-%#x%016x", uint64(x.hi), x.lo)
	case x.hi > 0:
		return fmt.Sprintf("%#x%016x", uint64(x.hi), x.lo)
	default:
		return fmt.Sprintf("%#x", x.lo) // ignore leading 0's
	}
}

// Abs returns the absolute value of an Int128's
func (x Int128) Abs() Int128 {
	if x.IsNeg() {
		return x.Neg()
	}
	return x
}

// Add returns the sum of two Int128's
func (x Int128) Add(y Int128) (z Int128) {
	z.hi = x.hi + y.hi
	z.lo = x.lo + y.lo
	if z.lo < x.lo {
		z.hi++
	}
	return z
}

// And returns the bitwise AND of two Int128's
func (x Int128) And(y Int128) (z Int128) {
	z.hi = int64(uint64(x.hi) & uint64(y.hi))
	z.lo = x.lo & y.lo
	return z
}

// AndNot returns the bitwise AndNot of two Int128's
func (x Int128) AndNot(y Int128) (z Int128) {
	z.hi = int64(uint64(x.hi) &^ uint64(y.hi))
	z.lo = x.lo &^ y.lo
	return z
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
func (x Int128) Cmp(y Int128) int {
	switch {
	case x.hi > y.hi:
		return 1
	case x.hi < y.hi:
		return -1
	case x.lo > y.lo:
		return 1
	case x.lo < y.lo:
		return -1
	}
	return 0
}

// CmpAbs compares |x| and |y| and returns:
//
//   -1 if |x| <  |y|
//    0 if |x| == |y|
//   +1 if |x| >  |y|
func (x Int128) CmpAbs(y Int128) int {
	x = x.Abs()
	y = y.Abs()
	return x.Cmp(y)
}

// Div returns the quotient corresponding to the provided dividend and divisor
//
// Div panics on division by 0. It checks some common/faster cases before fully committing to long division. This can probably be further optimized by
// implementing a successive approximation algorithm, with an initial seed value determined by a 64-bit division of the most significant bits.
func (x Int128) Div(d Int128) (q Int128) {
	q, _ = x.DivMod(d)
	return q
}

// DivMod returns the quotient and remainder corresponding to the provided dividend and divisor
//
// DivMod panics on division by 0. It checks some common/faster cases before fully committing to long division. This can probably be further optimized by
// implementing a successive approximation algorithm, with an initial seed value determined by a 64-bit division of the most significant bits.
func (x Int128) DivMod(d Int128) (q, r Int128) {
	var zero Int128
	qSign, rSign := +1, +1
	if x.Lt(zero) {
		qSign, rSign = -1, -1
		x = x.Neg()
	}
	if d.Lt(zero) {
		qSign = -qSign
		d = d.Neg()
	}
	qAbs, rAbs := x.Uint128().DivMod(d.Uint128())
	q, r = qAbs.Int128(), rAbs.Int128()
	if qSign < 0 {
		q = q.Neg()
	}
	if rSign < 0 {
		r = r.Neg()
	}
	return q, r
}

// Eq returns whether x is equal to y
func (x Int128) Eq(y Int128) bool {
	return x.hi == y.hi && x.lo == y.lo
}

// Gt returns whether x is greater than y
func (x Int128) Gt(y Int128) bool {
	switch {
	case x.hi > y.hi:
		return true
	case x.hi < y.hi:
		return false
	case x.lo > y.lo:
		return true
	default:
		return false
	}
}

// Gte returns whether x is greater than or equal to y
func (x Int128) Gte(y Int128) bool {
	switch {
	case x.hi > y.hi:
		return true
	case x.hi < y.hi:
		return false
	case x.lo >= y.lo:
		return true
	default:
		return false
	}
}

// IsInt64 checks if the Int128 can be represented as an int64
func (x Int128) IsInt64() bool {
	switch x.hi {
	case 0:
		return x.lo <= maxInt64
	case -1:
		return x.lo > maxInt64
	default:
		return false
	}
}

// IsNeg returns whether or not the Int128 is negative
func (x Int128) IsNeg() bool {
	if x.hi < 0 {
		return true
	}
	return false
}

// IsPos returns whether or not the Int128 is positive
func (x Int128) IsPos() bool {
	switch {
	case x.hi < 0:
		return false
	case x.lo > 0:
		return true
	default: // x is zero
		return false
	}
}

// IsUint64 checks if the Int128 can be represented as a uint64 without wrapping
func (x Int128) IsUint64() bool {
	return x.hi == 0
}

// Int128FromBigInt returns a Int128 from a big.Int
func Int128FromBigInt(a *big.Int) (z Int128) {
	var y Uint128
	neg := false
	if a.Sign() == -1 {
		a = new(big.Int).Neg(a)
		neg = true
	}
	y.lo = a.Uint64()
	b := new(big.Int).Rsh(a, int64Size)
	y.hi = b.Uint64()
	z = y.Int128()
	if neg {
		return z.Neg()
	}
	return z
}

// Int128FromHiLo returns a Int128 from the high and low 64 bits
func Int128FromHiLo(hi int64, lo uint64) Int128 {
	return Int128{hi: hi, lo: lo}
}

// Int128FromInt64 returns a Int128 from an int64
func Int128FromInt64(x int64) Int128 {
	if x >= 0 {
		return Int128{hi: 0, lo: uint64(x)}
	}
	return Int128{hi: -1, lo: uint64(x)}
}

// Int64 returns a representation of the Int128 as the builtin int64
//
// This function overflows silently
func (x Int128) Int64() int64 {
	return int64(x.lo)
}

// LShift returns a Int128 left-shifted by 1
func (x Int128) LShift() (z Int128) {
	z.hi = int64(uint64(x.hi)<<1 | x.lo>>(int64Size-1))
	z.lo = x.lo << 1
	return z
}

// LShiftN returns a Int128 left-shifted by a uint (i.e. x << n)
func (x Int128) LShiftN(n uint) (z Int128) {
	switch {
	case n >= int128Size:
		z.hi = 0
		z.lo = 0
	case n >= int64Size:
		z.hi = int64(x.lo << (n - int64Size))
		z.lo = 0
	default:
		z.hi = int64(uint64(x.hi)<<n | x.lo>>(int64Size-n))
		z.lo = x.lo << n
	}
	return z
}

// lShiftNActual returns a Int128 left-shifted by a uint (i.e. x << n)
//
// Unlike LShiftN, it operates on the actual 2's complement representation
func (x Int128) lShiftNActual(n uint) (z Int128) {
	switch {
	case n >= int128Size:
		z.hi = 0
		z.lo = 0
	case n >= int64Size:
		z.hi = int64(x.lo << (n - int64Size))
		z.lo = 0
	default:
		z.hi = int64(uint64(x.hi)<<n | x.lo>>(int64Size-n))
		z.lo = x.lo << n
	}
	return z
}

// Lt returns whether x is less than y
func (x Int128) Lt(y Int128) bool {
	switch {
	case x.hi < y.hi:
		return true
	case x.hi > y.hi:
		return false
	case x.lo < y.lo:
		return true
	default:
		return false
	}
}

// Lte returns whether x is less than or equal to y
func (x Int128) Lte(y Int128) bool {
	switch {
	case x.hi < y.hi:
		return true
	case x.hi > y.hi:
		return false
	case x.lo <= y.lo:
		return true
	default:
		return false
	}
}

// Mod returns the remainder corresponding to the provided dividend and divisor
//
// Mod panics on division by 0. It checks some common/faster cases before fully committing to long division. This can probably be further optimized by
// implementing a successive approximation algorithm, with an initial seed value determined by a 64-bit division of the most significant bits.
func (x Int128) Mod(d Int128) (r Int128) {
	_, r = x.DivMod(d)
	return r
}

// Mul returns the product of two Int128's
func (x Int128) Mul(y Int128) (z Int128) {
	var i uint
	yhi := uint64(y.hi)
	for i = 0; i < int64Size; i++ {
		if y.lo&(1<<i) != 0 {
			z = z.Add(x.lShiftNActual(i))
		}
	}
	for i = 0; i < int64Size; i++ {
		if yhi&(1<<i) != 0 {
			z = z.Add(x.lShiftNActual(i + int64Size))
		}
	}
	return z
}

// Nand returns the bitwise NAND of two Int128's
func (x Int128) Nand(y Int128) (z Int128) {
	z.hi = int64(^(uint64(x.hi) & uint64(y.hi)))
	z.lo = ^(x.lo & y.lo)
	return z
}

// Neg returns the additive inverse of an Int128
func (x Int128) Neg() (z Int128) {
	/* TODO: Something is wrong about the below
	z.hi = -x.hi
	z.lo = -x.lo
	if z.lo > x.lo {
		z.hi--
	}
	return z
	*/
	var zero Int128
	return zero.Sub(x)
}

// Nor returns the bitwise NOR of two Int128's
func (x Int128) Nor(y Int128) (z Int128) {
	z.hi = int64(^(uint64(x.hi) | uint64(y.hi)))
	z.lo = ^(x.lo | y.lo)
	return z
}

// Not returns the bitwise Not of an Int128
func (x Int128) Not() (z Int128) {
	z.hi = int64(^uint64(x.hi))
	z.lo = ^x.lo
	return z
}

// Or returns the bitwise OR of two Int128's
func (x Int128) Or(y Int128) (z Int128) {
	z.hi = int64(uint64(x.hi) | uint64(y.hi))
	z.lo = x.lo | y.lo
	return z
}

// RShift returns a Int128 right-shifted by 1
func (x Int128) RShift() (z Int128) {
	xhi := uint64(x.hi)
	z.hi = int64(xhi >> 1)
	z.lo = xhi<<(int64Size-1) | x.lo>>1
	return z
}

// RShiftN returns a Int128 right-shifted by a uint (i.e. x >> n)
func (x Int128) RShiftN(n uint) (z Int128) {
	switch {
	case n >= int128Size:
		z.hi = x.hi >> (int64Size - 1) // sign extend
		z.lo = uint64(z.hi)
	case n >= int64Size:
		z.hi = x.hi >> (int64Size - 1) // sign extend
		z.lo = uint64(x.hi) >> (n - int64Size)
	default:
		z.hi = x.hi >> n
		z.lo = uint64(x.hi)<<(int64Size-n) | x.lo>>n
	}
	return z
}

// RShift28 returns a Int128 right-shifted by a Uint128 (i.e. x >> y)
func (x Int128) RShift28(y Uint128) (z Int128) {
	if y.hi != 0 || y.lo >= int128Size {
		return x.RShiftN(int128Size)
	}
	return x.RShiftN(uint(y.lo))
}

// Sign returns the sign of an Int128
func (x Int128) Sign() int {
	switch {
	case x.hi > 0:
		return 1
	case x.hi < 0:
		return -1
	case x.lo > 0:
		return 1
	}
	return 0
}

// Sub returns the difference of two Int128's
func (x Int128) Sub(y Int128) (z Int128) {
	/* TODO: Something is wrong about the below
	z.hi = x.hi - y.hi
	z.lo = x.lo - y.lo
	if z.lo > x.lo {
		z.hi--
	}
	return z
	*/
	return x.Uint128().Sub(y.Uint128()).Int128()
}

// Uint128 returns a Uint128 representation of a Int128
//
// This function overflows silently
func (x Int128) Uint128() (z Uint128) {
	z.hi = uint64(x.hi)
	z.lo = x.lo
	return z
}

// Uint64 returns a representation of the Int128 as the builtin uint64
//
// This function overflows silently
func (x Int128) Uint64() uint64 {
	return x.lo
}

// Xor returns the bitwise XOR of two Int128's
func (x Int128) Xor(y Int128) (z Int128) {
	z.hi = int64(uint64(x.hi) ^ uint64(y.hi))
	z.lo = x.lo ^ y.lo
	return z
}

// Xnor returns the bitwise XNOR of two Int128's
func (x Int128) Xnor(y Int128) (z Int128) {
	z.hi = int64(^(uint64(x.hi) ^ uint64(y.hi)))
	z.lo = ^(x.lo ^ y.lo)
	return z
}
