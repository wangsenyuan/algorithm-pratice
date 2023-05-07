package p2673

import "testing"

func runSample(t *testing.T, n int, cost []int, expect int) {
	res := minIncrements(n, cost)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	cost := []int{1, 5, 2, 2, 3, 3, 1}
	expect := 6
	runSample(t, n, cost, expect)
}
