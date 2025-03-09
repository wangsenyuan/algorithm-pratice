package p3480

import "testing"

func runSample(t *testing.T, n int, conflictingPairs [][]int, expect int) {
	res := maxSubarrays(n, conflictingPairs)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	pairs := [][]int{{2, 3}, {1, 4}}
	expect := 9
	runSample(t, n, pairs, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	pairs := [][]int{{1, 2}, {2, 5}, {3, 5}}
	expect := 12
	runSample(t, n, pairs, expect)
}
