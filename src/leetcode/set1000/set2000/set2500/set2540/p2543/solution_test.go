package p2543

import "testing"

func runSample(t *testing.T, x, y int, expect bool) {
	res := isReachable(x, y)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 6
	y := 9
	expect := false
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x := 4
	y := 7
	expect := true
	runSample(t, x, y, expect)
}
