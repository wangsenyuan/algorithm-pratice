package p2957

import "testing"

func runSample(t *testing.T, n int, maxDist int, roads [][]int, expect int) {
	res := numberOfSets(n, maxDist, roads)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	maxDist := 5
	roads := [][]int{
		{0, 1, 20}, {0, 1, 10}, {1, 2, 2}, {0, 2, 2},
	}
	expect := 7
	runSample(t, n, maxDist, roads, expect)
}
