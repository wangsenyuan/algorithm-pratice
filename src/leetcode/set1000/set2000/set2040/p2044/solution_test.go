package p2044

import "testing"

func runSample(t *testing.T, n int, edges [][]int, time int, change int, expect int) {
	res := secondMinimum(n, edges, time, change)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5},
	}
	time := 3
	change := 5
	expect := 13
	runSample(t, n, edges, time, change, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2},
	}
	time := 3
	change := 2
	expect := 11
	runSample(t, n, edges, time, change, expect)
}
