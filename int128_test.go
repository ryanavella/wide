package wide

import (
	"testing"
)

func TestStringInt128(t *testing.T) {
	tests := []struct {
		inp      Int128
		expected string
	}{
		{Int128{hi: 0x0, lo: 0x0}, "0x0"},
		{Int128{hi: maxInt64, lo: maxUint64}, "0x7fffffffffffffffffffffffffffffff"},
		{Int128{hi: minInt64, lo: 0}, "-0x80000000000000000000000000000000"},
		{Int128{hi: 0xdeadbeef, lo: 0xbaadf00d}, "0xdeadbeef00000000baadf00d"},
	}
	for _, test := range tests {
		result := test.inp.String()
		if result != test.expected {
			t.Errorf("Expected %+v.String() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestAddInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 2}, Int128{hi: 0, lo: 3}},
		{Int128{hi: 0, lo: 2}, Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 3}},
		{Int128{hi: 1, lo: 0}, Int128{hi: 2, lo: 0}, Int128{hi: 3, lo: 0}},
		{Int128{hi: 2, lo: 0}, Int128{hi: 1, lo: 0}, Int128{hi: 3, lo: 0}},
		{Int128{hi: 0, lo: maxUint64}, Int128{hi: 0, lo: 1}, Int128{hi: 1, lo: 0}},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: maxUint64}, Int128{hi: 1, lo: 0}},
		{Int128{hi: maxInt64, lo: 0}, Int128{hi: 1, lo: 0}, Int128{hi: minInt64, lo: 0}},
		{Int128{hi: 1, lo: 0}, Int128{hi: maxInt64, lo: 0}, Int128{hi: minInt64, lo: 0}},
		{Int128{hi: maxInt64, lo: maxUint64}, Int128{hi: 0, lo: 1}, Int128{hi: minInt64, lo: 0}},
		{Int128{hi: 0, lo: 1}, Int128{hi: maxInt64, lo: maxUint64}, Int128{hi: minInt64, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Add(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Add(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestAndInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}},
	}
	for _, test := range tests {
		result := test.op1.And(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.And(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestAndNotInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.AndNot(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.AndNot(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestCmpInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected int
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, 0},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 0}, +1},
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 1}, -1},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}, 0},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 0}, +1},
		{Int128{hi: 0, lo: 0}, Int128{hi: 1, lo: 0}, -1},
		{Int128{hi: 1, lo: 0}, Int128{hi: 1, lo: 0}, 0},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: maxUint64}, +1},
		{Int128{hi: 0, lo: maxUint64}, Int128{hi: 1, lo: 0}, -1},
		{Int128{hi: maxInt64, lo: maxUint64}, Int128{hi: maxInt64, lo: maxUint64 - 1}, +1},
		{Int128{hi: maxInt64, lo: maxUint64 - 1}, Int128{hi: maxInt64, lo: maxUint64}, -1},
	}
	for _, test := range tests {
		result := test.op1.Cmp(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Cmp(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestDivInt128(t *testing.T) {
	tests := []struct {
		expected Int128
		op2      Int128
		op1      Int128
	}{
		// Some basic division tests with all-positive arguments
		{Int128{hi: 0, lo: 3}, Int128{hi: 0, lo: 5}, Int128{hi: 0, lo: 15}},
		{Int128{hi: 0, lo: 5}, Int128{hi: 0, lo: 3}, Int128{hi: 0, lo: 15}},
		{Int128{hi: 3, lo: 0}, Int128{hi: 0, lo: 5}, Int128{hi: 15, lo: 0}},
		{Int128{hi: 5, lo: 0}, Int128{hi: 0, lo: 3}, Int128{hi: 15, lo: 0}},
		{Int128{hi: 0, lo: 1 << 63}, Int128{hi: 0, lo: 2}, Int128{hi: 1, lo: 0}},
		{Int128{hi: 0, lo: 2}, Int128{hi: 0, lo: 1 << 63}, Int128{hi: 1, lo: 0}},
		// Testing the resulting sign for: 1/1, 1/-1, -1/1, -1/-1
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 1}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 1}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: 0, lo: 1}, Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}},
	}
	for _, test := range tests {
		result := test.op1.Div(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Div(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestDivModInt128(t *testing.T) {
	tests := []struct {
		op1       Int128
		op2       Int128
		expected1 Int128
		expected2 Int128
	}{
		// Edge cases
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 1}, Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 3}, Int128{hi: 0, lo: 2}, Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}},
		{Int128{hi: 3, lo: 0}, Int128{hi: 2, lo: 0}, Int128{hi: 0, lo: 1}, Int128{hi: 1, lo: 0}},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 65535}, Int128{hi: 0, lo: 0x1000100010001}, Int128{hi: 0, lo: 1}},
		// TODO: More tests for negative dividends and divisors
	}
	for _, test := range tests {
		result1, result2 := test.op1.DivMod(test.op2)
		if result1.lo != test.expected1.lo || result1.hi != test.expected1.hi || result2.lo != test.expected2.lo || result2.hi != test.expected2.hi {
			t.Errorf("Expected %s.DivMod(%s) == %s, %s got: %s, %s", test.op1, test.op2, test.expected1, test.expected2, result1, result2)
		}
	}
}

func TestDivByZeroInt128(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Division by 0 did not panic")
		}
	}()
	Int128{hi: 1, lo: 1}.DivMod(Int128{hi: 0, lo: 0})
}

func TestEqInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected bool
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, true},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}, true},
		{Int128{hi: 1, lo: 0}, Int128{hi: 1, lo: 0}, true},
		{Int128{hi: 1, lo: 1}, Int128{hi: 1, lo: 1}, true},
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 1}, false},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 0}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 0}, false},
		{Int128{hi: 0, lo: 0}, Int128{hi: 1, lo: 0}, false},
		{Int128{hi: 0, lo: 1}, Int128{hi: 1, lo: 0}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 1}, false},
	}
	for _, test := range tests {
		result := test.op1.Eq(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Eq(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestGtInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected bool
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, false},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 1, lo: 0}, false},
		{Int128{hi: 1, lo: 1}, Int128{hi: 1, lo: 1}, false},
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 1}, false},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 0}, true},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 0}, true},
		{Int128{hi: 0, lo: 0}, Int128{hi: 1, lo: 0}, false},
		{Int128{hi: 0, lo: 1}, Int128{hi: 1, lo: 0}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 1}, true},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: maxUint64}, true},
	}
	for _, test := range tests {
		result := test.op1.Gt(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Gt(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestGteInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected bool
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, true},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}, true},
		{Int128{hi: 1, lo: 0}, Int128{hi: 1, lo: 0}, true},
		{Int128{hi: 1, lo: 1}, Int128{hi: 1, lo: 1}, true},
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 1}, false},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 0}, true},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 0}, true},
		{Int128{hi: 0, lo: 0}, Int128{hi: 1, lo: 0}, false},
		{Int128{hi: 0, lo: 1}, Int128{hi: 1, lo: 0}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 1}, true},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: maxUint64}, true},
	}
	for _, test := range tests {
		result := test.op1.Gte(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Gte(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestIsInt64Int128(t *testing.T) {
	tests := []struct {
		inp      Int128
		expected bool
	}{
		{Int128{hi: 0, lo: 0}, true},
		{Int128{hi: 0, lo: maxInt64}, true},
		{Int128{hi: 0, lo: maxInt64 + 1}, false},
		{Int128{hi: 1, lo: 0}, false},
		{Int128{hi: maxInt64, lo: maxUint64}, false},
		{Int128{hi: -1, lo: maxUint64}, true},
		{Int128{hi: -1, lo: maxInt64}, false},
		{Int128{hi: -1, lo: maxInt64 + 1}, true},
	}
	for _, test := range tests {
		result := test.inp.IsInt64()
		if test.expected != result {
			t.Errorf("Expected %s.IsInt64() == %v, got: %v", test.inp, test.expected, result)
		}
	}
}

func TestIsUint64Int128(t *testing.T) {
	tests := []struct {
		inp      Int128
		expected bool
	}{
		{Int128{hi: 0, lo: 0}, true},
		{Int128{hi: 0, lo: maxUint64}, true},
		{Int128{hi: 1, lo: 0}, false},
		{Int128{hi: maxInt64, lo: 0}, false},
		{Int128{hi: maxInt64, lo: maxUint64}, false},
	}
	for _, test := range tests {
		result := test.inp.IsUint64()
		if test.expected != result {
			t.Errorf("Expected %s.IsUint64() == %v, got: %v", test.inp, test.expected, result)
		}
	}
}

func TestInt64Int128(t *testing.T) {
	tests := []struct {
		inp      Int128
		expected int64
	}{
		{Int128{hi: 0, lo: 0}, 0},
		{Int128{hi: 0, lo: maxInt64}, maxInt64},
		{Int128{hi: 0, lo: maxInt64 + 1}, minInt64},
		{Int128{hi: 0, lo: maxUint64}, -1},
		{Int128{hi: 1, lo: 0}, 0},
		{Int128{hi: maxInt64, lo: 0}, 0},
		{Int128{hi: maxInt64, lo: maxInt64}, maxInt64},
	}
	for _, test := range tests {
		result := test.inp.Int64()
		if test.expected != result {
			t.Errorf("Expected %s.Int64() == %v, got: %v", test.inp, test.expected, result)
		}
	}
}

func TestLShiftInt128(t *testing.T) {
	tests := []struct {
		inp      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1 << 1}},
		{Int128{hi: 0, lo: maxUint64}, Int128{hi: 1, lo: maxUint64 - 1}},
		{Int128{hi: maxInt64 >> 1, lo: 1 << 63}, Int128{hi: maxInt64, lo: 0}},
	}
	for _, test := range tests {
		result := test.inp.LShift()
		if test.expected != result {
			t.Errorf("Expected %s.LShift() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestLShiftNInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      uint
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, 0, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 0}, 1, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 0}, 2, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 1}, 0, Int128{hi: 0, lo: 1}},
		{Int128{hi: 0, lo: 1}, 1, Int128{hi: 0, lo: 2}},
		{Int128{hi: 0, lo: 1}, 2, Int128{hi: 0, lo: 4}},
		{Int128{hi: 0, lo: 1}, 63, Int128{hi: 0, lo: 1 << 63}},
		{Int128{hi: 0, lo: 1}, 64, Int128{hi: 1, lo: 0}},
		{Int128{hi: 1, lo: 0}, 0, Int128{hi: 1, lo: 0}},
		{Int128{hi: 1, lo: 0}, 1, Int128{hi: 2, lo: 0}},
		{Int128{hi: 1, lo: 0}, 2, Int128{hi: 4, lo: 0}},
		{Int128{hi: 1, lo: 0}, 64, Int128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.LShiftN(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.LShiftN(%v) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestLtInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected bool
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, false},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 1, lo: 0}, false},
		{Int128{hi: 1, lo: 1}, Int128{hi: 1, lo: 1}, false},
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 1}, true},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 0}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 0}, false},
		{Int128{hi: 0, lo: 0}, Int128{hi: 1, lo: 0}, true},
		{Int128{hi: 0, lo: 1}, Int128{hi: 1, lo: 0}, true},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 1}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: maxUint64}, false},
		{Int128{hi: 0, lo: maxUint64}, Int128{hi: 1, lo: 0}, true},
	}
	for _, test := range tests {
		result := test.op1.Lt(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Lt(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestLteInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected bool
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, true},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 1}, true},
		{Int128{hi: 1, lo: 0}, Int128{hi: 1, lo: 0}, true},
		{Int128{hi: 1, lo: 1}, Int128{hi: 1, lo: 1}, true},
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 1}, true},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 0}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 0}, false},
		{Int128{hi: 0, lo: 0}, Int128{hi: 1, lo: 0}, true},
		{Int128{hi: 0, lo: 1}, Int128{hi: 1, lo: 0}, true},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 1}, false},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: maxUint64}, false},
		{Int128{hi: 0, lo: maxUint64}, Int128{hi: 1, lo: 0}, true},
	}
	for _, test := range tests {
		result := test.op1.Lte(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Lte(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestMulInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 3}, Int128{hi: 0, lo: 5}, Int128{hi: 0, lo: 15}},
		{Int128{hi: 0, lo: 5}, Int128{hi: 0, lo: 3}, Int128{hi: 0, lo: 15}},
		{Int128{hi: 3, lo: 0}, Int128{hi: 0, lo: 5}, Int128{hi: 15, lo: 0}},
		{Int128{hi: 5, lo: 0}, Int128{hi: 0, lo: 3}, Int128{hi: 15, lo: 0}},
		{Int128{hi: 0, lo: 1 << 63}, Int128{hi: 0, lo: 2}, Int128{hi: 1, lo: 0}},
		{Int128{hi: 0, lo: 2}, Int128{hi: 0, lo: 1 << 63}, Int128{hi: 1, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Mul(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Mul(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestNegInt128(t *testing.T) {
	tests := []struct {
		inp      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 1}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: 0, lo: 2}, Int128{hi: -1, lo: maxUint64 - 1}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 1}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 1}},
		{Int128{hi: minInt64, lo: 0}, Int128{hi: minInt64, lo: 0}}, // most negative number has no positive counterpart
	}
	for _, test := range tests {
		result := test.inp.Neg()
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Neg() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestNorInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Nor(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Nor(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestNotInt128(t *testing.T) {
	tests := []struct {
		inp      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.inp.Not()
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Not() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestOrInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}},
	}
	for _, test := range tests {
		result := test.op1.Or(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Or(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestRShiftInt128(t *testing.T) {
	tests := []struct {
		inp      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 2}, Int128{hi: 0, lo: 1}},
		{Int128{hi: 0, lo: 4}, Int128{hi: 0, lo: 2}},
		{Int128{hi: 1, lo: maxUint64 - 1}, Int128{hi: 0, lo: maxUint64}},
		{Int128{hi: 1, lo: 0}, Int128{hi: 0, lo: 1 << 63}},
	}
	for _, test := range tests {
		result := test.inp.RShift()
		if test.expected != result {
			t.Errorf("Expected %s.RShift() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestRShiftNInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      uint
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, 0, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 0}, 1, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 1}, 0, Int128{hi: 0, lo: 1}},
		{Int128{hi: 0, lo: 1}, 1, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 2}, 1, Int128{hi: 0, lo: 1}},
		{Int128{hi: 0, lo: 4}, 2, Int128{hi: 0, lo: 1}},
		{Int128{hi: 0, lo: 1 << 63}, 63, Int128{hi: 0, lo: 1}},
		{Int128{hi: 1, lo: 0}, 1, Int128{hi: 0, lo: 1 << 63}},
		{Int128{hi: 2, lo: 0}, 1, Int128{hi: 1, lo: 0}},
		{Int128{hi: 4, lo: 0}, 1, Int128{hi: 2, lo: 0}},
		{Int128{hi: 4, lo: 0}, 2, Int128{hi: 1, lo: 0}},
		{Int128{hi: 1, lo: 0}, 64, Int128{hi: 0, lo: 1}},
		{Int128{hi: maxInt64, lo: maxUint64}, 126, Int128{hi: 0, lo: 1}},
		{Int128{hi: maxInt64, lo: maxUint64}, 127, Int128{hi: 0, lo: 0}},

		{Int128{hi: -1, lo: maxUint64}, 0, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: -1, lo: maxUint64}, 1, Int128{hi: 0, lo: 0}},
		{Int128{hi: -1, lo: maxUint64 - 1}, 1, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: -1, lo: maxUint64 - 3}, 2, Int128{hi: -1, lo: maxUint64}},

		{Int128{hi: minInt64, lo: 0}, 127, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: minInt64, lo: 0}, 128, Int128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.RShiftN(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.RShiftN(%v) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestSubInt128(t *testing.T) {
	tests := []struct {
		expected Int128
		op2      Int128
		op1      Int128
	}{
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 2}, Int128{hi: 0, lo: 3}},
		{Int128{hi: 0, lo: 2}, Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: 3}},
		{Int128{hi: 1, lo: 0}, Int128{hi: 2, lo: 0}, Int128{hi: 3, lo: 0}},
		{Int128{hi: 2, lo: 0}, Int128{hi: 1, lo: 0}, Int128{hi: 3, lo: 0}},
		{Int128{hi: 0, lo: maxUint64}, Int128{hi: 0, lo: 1}, Int128{hi: 1, lo: 0}},
		{Int128{hi: 0, lo: 1}, Int128{hi: 0, lo: maxUint64}, Int128{hi: 1, lo: 0}},
		{Int128{hi: maxInt64, lo: 0}, Int128{hi: 1, lo: 0}, Int128{hi: minInt64, lo: 0}},
		{Int128{hi: 1, lo: 0}, Int128{hi: maxInt64, lo: 0}, Int128{hi: minInt64, lo: 0}},
		{Int128{hi: maxInt64, lo: maxUint64}, Int128{hi: 0, lo: 1}, Int128{hi: minInt64, lo: 0}},
		{Int128{hi: 0, lo: 1}, Int128{hi: maxInt64, lo: maxUint64}, Int128{hi: minInt64, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Sub(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Sub(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestUint64Int128(t *testing.T) {
	tests := []struct {
		inp      Int128
		expected uint64
	}{
		{Int128{hi: 0, lo: 0}, 0},
		{Int128{hi: 0, lo: maxInt64}, maxInt64},
		{Int128{hi: 0, lo: maxInt64 + 1}, maxInt64 + 1},
		{Int128{hi: 0, lo: maxUint64}, maxUint64},
		{Int128{hi: 1, lo: 0}, 0},
		{Int128{hi: 1, lo: maxUint64}, maxUint64},
		{Int128{hi: maxInt64, lo: 0}, 0},
		{Int128{hi: maxInt64, lo: maxUint64}, maxUint64},
		{Int128{hi: -1, lo: maxUint64}, maxUint64},
	}
	for _, test := range tests {
		result := test.inp.Uint64()
		if test.expected != result {
			t.Errorf("Expected %s.Uint64() == %v, got: %v", test.inp, test.expected, result)
		}
	}
}

func TestXorInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Xor(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Xor(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestXnorInt128(t *testing.T) {
	tests := []struct {
		op1      Int128
		op2      Int128
		expected Int128
	}{
		{Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}},
		{Int128{hi: 0, lo: 0}, Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: 0, lo: 0}, Int128{hi: 0, lo: 0}},
		{Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}, Int128{hi: -1, lo: maxUint64}},
	}
	for _, test := range tests {
		result := test.op1.Xnor(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Xnor(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}
