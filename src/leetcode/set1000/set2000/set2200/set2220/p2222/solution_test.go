package p2222

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := int(numberOfWays(s))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "001101"
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "11100"
	expect := 0
	runSample(t, s, expect)
}
