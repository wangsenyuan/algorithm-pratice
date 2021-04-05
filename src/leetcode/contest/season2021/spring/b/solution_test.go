package b

import "testing"

func runSample(t *testing.T, n, x, y int, expect int) {
	res := orchestraLayout(n, x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	x := 0
	y := 2
	expect := 3
	runSample(t, n, x, y, expect)
}
func TestSample2(t *testing.T) {
	n := 3
	x := 1
	y := 1
	expect := 9
	runSample(t, n, x, y, expect)
}
