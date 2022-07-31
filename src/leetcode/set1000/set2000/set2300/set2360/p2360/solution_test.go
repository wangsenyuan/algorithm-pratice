package p2360

import "testing"

func runSample(t *testing.T, edges []int, expect int) {
	res := longestCycle(edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := []int{3, 3, 4, 2, 3}
	expect := 3
	runSample(t, edges, expect)
}

func TestSample2(t *testing.T) {
	edges := []int{-1, 4, -1, 2, 0, 4}
	expect := -1
	runSample(t, edges, expect)
}
