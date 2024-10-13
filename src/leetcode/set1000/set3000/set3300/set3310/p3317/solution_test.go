package p3317

import "testing"

func runSample(t *testing.T, n int, x int, y int, expect int) {
	res := numberOfWays(n, x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x, y := 1, 2, 3
	expect := 6
	runSample(t, n, x, y, expect)
}

func TestSample2(t *testing.T) {
	n, x, y := 5, 2, 1
	expect := 32
	runSample(t, n, x, y, expect)
}

func TestSample3(t *testing.T) {
	n, x, y := 3, 3, 4
	expect := 684
	runSample(t, n, x, y, expect)
}
