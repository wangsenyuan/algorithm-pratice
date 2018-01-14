package p762

import "testing"

func runSample(t *testing.T, L int, R int, expect int) {
	res := countPrimeSetBits(L, R)
	if res != expect {
		t.Errorf("sample %d %d, should give %d, but got %d", L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	L, R := 6, 10
	expect := 4
	runSample(t, L, R, expect)
}

func TestSample2(t *testing.T) {
	L, R := 10, 15
	expect := 5
	runSample(t, L, R, expect)
}
