package wide

import (
	"testing"
)

func TestStringUint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected string
	}{
		{Uint128{hi: 0x0, lo: 0x0}, "0x0"},
		{Uint128{hi: maxUint64, lo: maxUint64}, "0xffffffffffffffffffffffffffffffff"},
		{Uint128{hi: 0xdeadbeef, lo: 0xbaadf00d}, "0xdeadbeef00000000baadf00d"},
	}
	for _, test := range tests {
		result := test.inp.String()
		if result != test.expected {
			t.Errorf("Expected %+v.String() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestAddUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 2}, Uint128{hi: 0, lo: 3}},
		{Uint128{hi: 0, lo: 2}, Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 3}},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 2, lo: 0}, Uint128{hi: 3, lo: 0}},
		{Uint128{hi: 2, lo: 0}, Uint128{hi: 1, lo: 0}, Uint128{hi: 3, lo: 0}},
		{Uint128{hi: 0, lo: maxUint64}, Uint128{hi: 0, lo: 1}, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: maxUint64}, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: maxUint64, lo: 0}, Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: maxUint64, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Add(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Add(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestAndUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}},
	}
	for _, test := range tests {
		result := test.op1.And(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.And(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestAndNotUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.AndNot(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.AndNot(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestCmpUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected int
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, 0},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}, +1},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 1}, -1},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1}, 0},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 0}, +1},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 1, lo: 0}, -1},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 1, lo: 0}, 0},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: maxUint64}, +1},
		{Uint128{hi: 0, lo: maxUint64}, Uint128{hi: 1, lo: 0}, -1},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64 - 1}, +1},
		{Uint128{hi: maxUint64, lo: maxUint64 - 1}, Uint128{hi: maxUint64, lo: maxUint64}, -1},
	}
	for _, test := range tests {
		result := test.op1.Cmp(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Cmp(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestDivModUint128(t *testing.T) {
	tests := []struct {
		op1       Uint128
		op2       Uint128
		expected1 Uint128
		expected2 Uint128
	}{
		// Edge cases
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 1}, Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 3}, Uint128{hi: 0, lo: 2}, Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: 3, lo: 0}, Uint128{hi: 2, lo: 0}, Uint128{hi: 0, lo: 1}, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 65535}, Uint128{hi: 0, lo: 0x1000100010001}, Uint128{hi: 0, lo: 1}},
		// Randomly generated tests, so that we aren't exclusively looking at edge cases
		{Uint128{hi: 18078839827447545686, lo: 14690264446208930433}, Uint128{hi: 8998388556054393762, lo: 1806737508591697479}, Uint128{hi: 0, lo: 2}, Uint128{hi: 82062715338758162, lo: 11076789429025535475}},
		{Uint128{hi: 16644964954028498953, lo: 6044532115116883114}, Uint128{hi: 201749283954463444, lo: 14360937540127383586}, Uint128{hi: 0, lo: 82}, Uint128{hi: 101523669762496481, lo: 9039274542082732486}},
		{Uint128{hi: 2093133179559761320, lo: 16602062711200797661}, Uint128{hi: 16864581198376902610, lo: 9459801703427125641}, Uint128{hi: 0, lo: 0}, Uint128{hi: 2093133179559761320, lo: 16602062711200797661}},
		{Uint128{hi: 15749230925143681054, lo: 5141327176156114442}, Uint128{hi: 7370181155092257276, lo: 6109534453883242647}, Uint128{hi: 0, lo: 2}, Uint128{hi: 1008868614959166501, lo: 11369002342099180764}},
		{Uint128{hi: 9688725833947941869, lo: 5829725676053634587}, Uint128{hi: 10771428793017394942, lo: 11481679973018363035}, Uint128{hi: 0, lo: 0}, Uint128{hi: 9688725833947941869, lo: 5829725676053634587}},
		{Uint128{hi: 13115430490891893466, lo: 8603977268700753405}, Uint128{hi: 9528846791087986823, lo: 10534857072914627768}, Uint128{hi: 0, lo: 1}, Uint128{hi: 3586583699803906642, lo: 16515864269495677253}},
		{Uint128{hi: 4464682195163803491, lo: 7489173075960054176}, Uint128{hi: 16981508686037311992, lo: 10249255444110602208}, Uint128{hi: 0, lo: 0}, Uint128{hi: 4464682195163803491, lo: 7489173075960054176}},
		{Uint128{hi: 14969323870889724866, lo: 2190849366522262425}, Uint128{hi: 15744153244977445381, lo: 6216957372856731963}, Uint128{hi: 0, lo: 0}, Uint128{hi: 14969323870889724866, lo: 2190849366522262425}},
		{Uint128{hi: 4242347467606309934, lo: 617379611895620583}, Uint128{hi: 10940397127107242540, lo: 2302117645027716276}, Uint128{hi: 0, lo: 0}, Uint128{hi: 4242347467606309934, lo: 617379611895620583}},
		{Uint128{hi: 13878862832891726409, lo: 12247025815474997093}, Uint128{hi: 15933048822954309307, lo: 10203778256787916680}, Uint128{hi: 0, lo: 0}, Uint128{hi: 13878862832891726409, lo: 12247025815474997093}},
		{Uint128{hi: 838346584253750531, lo: 8728622130014530564}, Uint128{hi: 7623255793094361036, lo: 8997606623242799596}, Uint128{hi: 0, lo: 0}, Uint128{hi: 838346584253750531, lo: 8728622130014530564}},
		{Uint128{hi: 16329672898340907698, lo: 1812476799316563631}, Uint128{hi: 15241696152186921136, lo: 805856988913712022}, Uint128{hi: 0, lo: 1}, Uint128{hi: 1087976746153986562, lo: 1006619810402851609}},
		{Uint128{hi: 6779770950567198933, lo: 5234941809987140509}, Uint128{hi: 12958508852484540729, lo: 6724476107841768838}, Uint128{hi: 0, lo: 0}, Uint128{hi: 6779770950567198933, lo: 5234941809987140509}},
		{Uint128{hi: 14498257309946160406, lo: 18084320459228438784}, Uint128{hi: 13094649420129182662, lo: 13044860498710942198}, Uint128{hi: 0, lo: 1}, Uint128{hi: 1403607889816977744, lo: 5039459960517496586}},
		{Uint128{hi: 15511746663743037887, lo: 9414436614834301257}, Uint128{hi: 241543290869542064, lo: 4264400748188406476}, Uint128{hi: 0, lo: 64}, Uint128{hi: 52976048092345776, lo: 13193949836419561033}},
		{Uint128{hi: 8084744342537816804, lo: 1423207416017422773}, Uint128{hi: 15656244447055551121, lo: 11594742766098937249}, Uint128{hi: 0, lo: 0}, Uint128{hi: 8084744342537816804, lo: 1423207416017422773}},
		{Uint128{hi: 4319800846968343225, lo: 14534725286450089249}, Uint128{hi: 13287988548133458958, lo: 5713573001805029812}, Uint128{hi: 0, lo: 0}, Uint128{hi: 4319800846968343225, lo: 14534725286450089249}},
		{Uint128{hi: 13508470227635185554, lo: 12357508788529651644}, Uint128{hi: 10774097062273294804, lo: 7568828230680467196}, Uint128{hi: 0, lo: 1}, Uint128{hi: 2734373165361890750, lo: 4788680557849184448}},
		{Uint128{hi: 14855998319421094055, lo: 4486542090348589603}, Uint128{hi: 6678463923086748523, lo: 1254263323666630076}, Uint128{hi: 0, lo: 2}, Uint128{hi: 1499070473247597009, lo: 1978015443015329451}},
		{Uint128{hi: 10551012307773401926, lo: 7564802811090251473}, Uint128{hi: 6463539642689907276, lo: 18401146206923932121}, Uint128{hi: 0, lo: 1}, Uint128{hi: 4087472665083494649, lo: 7610400677875870968}},
	}
	for _, test := range tests {
		result1, result2 := test.op1.DivMod(test.op2)
		if result1.lo != test.expected1.lo || result1.hi != test.expected1.hi || result2.lo != test.expected2.lo || result2.hi != test.expected2.hi {
			t.Errorf("Expected %s.DivMod(%s) == %s, %s got: %s, %s", test.op1, test.op2, test.expected1, test.expected2, result1, result2)
		}
	}
}

func TestDivByZeroUint128(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Division by 0 did not panic")
		}
	}()
	Uint128{hi: 1, lo: 1}.DivMod(Uint128{hi: 0, lo: 0})
}

func TestEqUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected bool
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, true},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1}, true},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 1, lo: 0}, true},
		{Uint128{hi: 1, lo: 1}, Uint128{hi: 1, lo: 1}, true},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 1}, false},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 0}, false},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 1}, false},
	}
	for _, test := range tests {
		result := test.op1.Eq(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Eq(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestGtUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected bool
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, false},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: 1, lo: 1}, Uint128{hi: 1, lo: 1}, false},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 1}, false},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}, true},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 0}, true},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 1}, true},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: maxUint64}, true},
	}
	for _, test := range tests {
		result := test.op1.Gt(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Gt(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestGteUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected bool
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, true},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1}, true},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 1, lo: 0}, true},
		{Uint128{hi: 1, lo: 1}, Uint128{hi: 1, lo: 1}, true},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 1}, false},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}, true},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 0}, true},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 1}, true},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: maxUint64}, true},
	}
	for _, test := range tests {
		result := test.op1.Gte(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Gte(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestIsInt64Uint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected bool
	}{
		{Uint128{hi: 0, lo: 0}, true},
		{Uint128{hi: 0, lo: maxInt64}, true},
		{Uint128{hi: 0, lo: maxInt64 + 1}, false},
		{Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: maxUint64, lo: maxUint64}, false},
	}
	for _, test := range tests {
		result := test.inp.IsInt64()
		if test.expected != result {
			t.Errorf("Expected %s.IsInt64() == %v, got: %v", test.inp, test.expected, result)
		}
	}
}

func TestIsUint64Uint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected bool
	}{
		{Uint128{hi: 0, lo: 0}, true},
		{Uint128{hi: 0, lo: maxUint64}, true},
		{Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: maxUint64, lo: 0}, false},
		{Uint128{hi: maxUint64, lo: maxUint64}, false},
	}
	for _, test := range tests {
		result := test.inp.IsUint64()
		if test.expected != result {
			t.Errorf("Expected %s.IsUint64() == %v, got: %v", test.inp, test.expected, result)
		}
	}
}

func TestInt64Uint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected int64
	}{
		{Uint128{hi: 0, lo: 0}, 0},
		{Uint128{hi: 0, lo: maxInt64}, maxInt64},
		{Uint128{hi: 0, lo: maxInt64 + 1}, minInt64},
		{Uint128{hi: 0, lo: maxUint64}, -1},
		{Uint128{hi: 1, lo: 0}, 0},
		{Uint128{hi: maxUint64, lo: 0}, 0},
		{Uint128{hi: maxUint64, lo: maxInt64}, maxInt64},
	}
	for _, test := range tests {
		result := test.inp.Int64()
		if test.expected != result {
			t.Errorf("Expected %s.Int64() == %v, got: %v", test.inp, test.expected, result)
		}
	}
}

func TestLenUint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected uint
	}{
		{Uint128{hi: 0, lo: 0}, 0},
		{Uint128{hi: 0, lo: 1}, 1},
		{Uint128{hi: 0, lo: 2}, 2},
		{Uint128{hi: 0, lo: 3}, 2},
		{Uint128{hi: 0, lo: 4}, 3},
		{Uint128{hi: 0, lo: maxUint64}, 64},
		{Uint128{hi: 1, lo: 0}, 65},
		{Uint128{hi: 2, lo: 0}, 66},
		{Uint128{hi: 3, lo: 0}, 66},
		{Uint128{hi: 4, lo: 0}, 67},
		{Uint128{hi: maxUint64, lo: 0}, 128},
	}
	for _, test := range tests {
		result := test.inp.Len()
		if test.expected != result {
			t.Errorf("Expected %s.Len() == %v, got: %v", test.inp, test.expected, result)
		}
	}
}

func TestLShiftUint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1 << 1}},
		{Uint128{hi: 0, lo: maxUint64}, Uint128{hi: 1, lo: maxUint64 - 1}},
		{Uint128{hi: maxUint64, lo: 0}, Uint128{hi: maxUint64 - 1, lo: 0}},
	}
	for _, test := range tests {
		result := test.inp.LShift()
		if test.expected != result {
			t.Errorf("Expected %s.LShift() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestLShiftNUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      uint
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, 0, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 0}, 1, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 0}, 2, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 1}, 0, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: 0, lo: 1}, 1, Uint128{hi: 0, lo: 2}},
		{Uint128{hi: 0, lo: 1}, 2, Uint128{hi: 0, lo: 4}},
		{Uint128{hi: 0, lo: 1}, 63, Uint128{hi: 0, lo: 1 << 63}},
		{Uint128{hi: 0, lo: 1}, 64, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: 0, lo: 1}, 127, Uint128{hi: 1 << 63, lo: 0}},
		{Uint128{hi: 1, lo: 0}, 0, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: 1, lo: 0}, 1, Uint128{hi: 2, lo: 0}},
		{Uint128{hi: 1, lo: 0}, 2, Uint128{hi: 4, lo: 0}},
		{Uint128{hi: 1, lo: 0}, 63, Uint128{hi: 1 << 63, lo: 0}},
		{Uint128{hi: 1, lo: 0}, 64, Uint128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.LShiftN(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.LShiftN(%v) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestLtUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected bool
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, false},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 1, lo: 0}, false},
		{Uint128{hi: 1, lo: 1}, Uint128{hi: 1, lo: 1}, false},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 1}, true},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 0}, false},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 1, lo: 0}, true},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 1, lo: 0}, true},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 1}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: maxUint64}, false},
		{Uint128{hi: 0, lo: maxUint64}, Uint128{hi: 1, lo: 0}, true},
	}
	for _, test := range tests {
		result := test.op1.Lt(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Lt(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestLteUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected bool
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, true},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 1}, true},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 1, lo: 0}, true},
		{Uint128{hi: 1, lo: 1}, Uint128{hi: 1, lo: 1}, true},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 1}, true},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 0}, false},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 1, lo: 0}, true},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 1, lo: 0}, true},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 1}, false},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: maxUint64}, false},
		{Uint128{hi: 0, lo: maxUint64}, Uint128{hi: 1, lo: 0}, true},
	}
	for _, test := range tests {
		result := test.op1.Lte(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.Lte(%s) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestMulUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 3}, Uint128{hi: 0, lo: 5}, Uint128{hi: 0, lo: 15}},
		{Uint128{hi: 0, lo: 5}, Uint128{hi: 0, lo: 3}, Uint128{hi: 0, lo: 15}},
		{Uint128{hi: 3, lo: 0}, Uint128{hi: 0, lo: 5}, Uint128{hi: 15, lo: 0}},
		{Uint128{hi: 5, lo: 0}, Uint128{hi: 0, lo: 3}, Uint128{hi: 15, lo: 0}},
		{Uint128{hi: 0, lo: 1 << 63}, Uint128{hi: 0, lo: 2}, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: 0, lo: 2}, Uint128{hi: 0, lo: 1 << 63}, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: 0, lo: 0xFFFFFFFFFFFFFFFF}, Uint128{hi: 0, lo: 0xFFFFFFFFFFFFFFFF}, Uint128{hi: 0xFFFFFFFFFFFFFFFE, lo: 1}},
	}
	for _, test := range tests {
		result := test.op1.Mul(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Mul(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestNandUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Nand(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Nand(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestNegUint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: 1 << 63, lo: 1}, Uint128{hi: 1<<63 - 1, lo: 1<<64 - 1}},
		{Uint128{hi: 1<<63 - 1, lo: 1<<64 - 1}, Uint128{hi: 1 << 63, lo: 1}},
		{Uint128{hi: 1 << 63, lo: 0}, Uint128{hi: 1 << 63, lo: 0}}, // most negative number has no positive counterpart
	}
	for _, test := range tests {
		result := test.inp.Neg()
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Neg() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestNorUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Nor(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Nor(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestNotUint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.inp.Not()
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Not() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestOrUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}},
	}
	for _, test := range tests {
		result := test.op1.Or(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Or(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestRShiftUint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 2}, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: 0, lo: 4}, Uint128{hi: 0, lo: 2}},
		{Uint128{hi: 1, lo: maxUint64 - 1}, Uint128{hi: 0, lo: maxUint64}},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 1 << 63}},
	}
	for _, test := range tests {
		result := test.inp.RShift()
		if test.expected != result {
			t.Errorf("Expected %s.RShift() == %s, got: %s", test.inp, test.expected, result)
		}
	}
}

func TestRShiftNUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      uint
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, 0, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 0}, 1, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 1}, 0, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: 0, lo: 1}, 1, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 2}, 1, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: 0, lo: 4}, 2, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: 0, lo: 1 << 63}, 63, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: 1, lo: 0}, 1, Uint128{hi: 0, lo: 1 << 63}},
		{Uint128{hi: 2, lo: 0}, 1, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: 4, lo: 0}, 1, Uint128{hi: 2, lo: 0}},
		{Uint128{hi: 4, lo: 0}, 2, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: 1, lo: 0}, 64, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: 1 << 63, lo: 0}, 127, Uint128{hi: 0, lo: 1}},
		{Uint128{hi: maxUint64, lo: maxUint64}, 127, Uint128{hi: 0, lo: 1}},
	}
	for _, test := range tests {
		result := test.op1.RShiftN(test.op2)
		if test.expected != result {
			t.Errorf("Expected %s.RShiftN(%v) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestSubUint128(t *testing.T) {
	tests := []struct {
		expected Uint128
		op2      Uint128
		op1      Uint128
	}{
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 2}, Uint128{hi: 0, lo: 3}},
		{Uint128{hi: 0, lo: 2}, Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 3}},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: 2, lo: 0}, Uint128{hi: 3, lo: 0}},
		{Uint128{hi: 2, lo: 0}, Uint128{hi: 1, lo: 0}, Uint128{hi: 3, lo: 0}},
		{Uint128{hi: 0, lo: maxUint64}, Uint128{hi: 0, lo: 1}, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: maxUint64}, Uint128{hi: 1, lo: 0}},
		{Uint128{hi: maxUint64, lo: 0}, Uint128{hi: 1, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 1, lo: 0}, Uint128{hi: maxUint64, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 1}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 1}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Sub(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Sub(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestUint64Uint128(t *testing.T) {
	tests := []struct {
		inp      Uint128
		expected uint64
	}{
		{Uint128{hi: 0, lo: 0}, 0},
		{Uint128{hi: 0, lo: maxInt64}, maxInt64},
		{Uint128{hi: 0, lo: maxInt64 + 1}, maxInt64 + 1},
		{Uint128{hi: 0, lo: maxUint64}, maxUint64},
		{Uint128{hi: 1, lo: 0}, 0},
		{Uint128{hi: 1, lo: maxUint64}, maxUint64},
		{Uint128{hi: maxUint64, lo: 0}, 0},
		{Uint128{hi: maxUint64, lo: maxUint64}, maxUint64},
	}
	for _, test := range tests {
		result := test.inp.Uint64()
		if test.expected != result {
			t.Errorf("Expected %s.Uint64() == %v, got: %v", test.inp, test.expected, result)
		}
	}
}

func TestXorUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
	}
	for _, test := range tests {
		result := test.op1.Xor(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Xor(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}

func TestXnorUint128(t *testing.T) {
	tests := []struct {
		op1      Uint128
		op2      Uint128
		expected Uint128
	}{
		{Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}},
		{Uint128{hi: 0, lo: 0}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: 0, lo: 0}, Uint128{hi: 0, lo: 0}},
		{Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}, Uint128{hi: maxUint64, lo: maxUint64}},
	}
	for _, test := range tests {
		result := test.op1.Xnor(test.op2)
		if result.lo != test.expected.lo || result.hi != test.expected.hi {
			t.Errorf("Expected %s.Xnor(%s) == %s, got: %s", test.op1, test.op2, test.expected, result)
		}
	}
}
