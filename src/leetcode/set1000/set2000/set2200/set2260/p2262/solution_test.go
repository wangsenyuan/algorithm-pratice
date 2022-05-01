package p2262

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := int(appealSum(s))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abbca"
	expect := 28
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "code"
	expect := 20
	runSample(t, s, expect)
}
