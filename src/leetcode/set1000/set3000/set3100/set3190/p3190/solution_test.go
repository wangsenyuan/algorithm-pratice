package p3190

import "testing"

func runSample(t *testing.T, n int, requirements [][]int, expect int) {
	res := numberOfPermutations(n, requirements)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	requirements := [][]int{
		{2, 2},
		{0, 0},
	}
	expect := 2
	runSample(t, n, requirements, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	requirements := [][]int{
		{2, 2},
		{1, 1},
		{0, 0},
	}
	expect := 1
	runSample(t, n, requirements, expect)
}
