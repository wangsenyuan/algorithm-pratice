package p1808

import "testing"

func runSample(t *testing.T, primeFactors int, expect int) {
	res := maxNiceDivisors(primeFactors)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pf := 73
	expect := 572712676
	runSample(t, pf, expect)
}
