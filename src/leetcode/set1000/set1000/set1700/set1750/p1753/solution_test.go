package p1753

import "testing"

func runSample(t *testing.T, a, b, c int, expect int) {
	res := maximumScore(a, b, c)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := 2
	b := 4
	c := 6
	expect := 6
	runSample(t, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	a := 4
	b := 4
	c := 6
	expect := 7
	runSample(t, a, b, c, expect)
}

func TestSample3(t *testing.T) {
	a := 1
	b := 8
	c := 8
	expect := 8
	runSample(t, a, b, c, expect)
}
