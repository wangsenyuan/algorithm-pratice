package p2049

import "testing"

func runSample(t *testing.T, parents []int, expect int) {
	res := countHighestScoreNodes(parents)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	parents := []int{-1, 2, 0}
	expect := 2
	runSample(t, parents, expect)
}
