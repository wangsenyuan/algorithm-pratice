package p2001

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := maxProduct(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "leetcodecom"
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "bb"
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "accbcaxxcxx"
	expect := 25
	runSample(t, s, expect)
}