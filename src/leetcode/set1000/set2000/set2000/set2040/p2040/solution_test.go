package p2040

import "testing"

func runSample(t *testing.T, edges [][]int, patience []int, expect int) {
	res := networkBecomesIdle(edges, patience)

	if res != expect {
		t.Errorf("Sample expect %d but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{0, 1}, {1, 2},
	}
	patience := []int{0, 2, 1}
	expect := 8
	runSample(t, edges, patience, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{0, 1}, {0, 2}, {1, 2},
	}
	patience := []int{0, 10, 10}
	expect := 3
	runSample(t, edges, patience, expect)
}
