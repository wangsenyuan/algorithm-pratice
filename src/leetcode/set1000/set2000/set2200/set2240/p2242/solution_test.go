package p2242

import "testing"

func runSample(t *testing.T, scores []int, edges [][]int, expect int) {
	res := maximumScore(scores, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	scores := []int{5, 2, 9, 8, 4}
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4},
	}
	expect := 24

	runSample(t, scores, edges, expect)
}
