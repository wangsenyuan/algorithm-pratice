package p3356

import "testing"

func runSample(t *testing.T, a []int, queries [][]int, expect int) {
	res := minZeroArray(a, queries)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", a, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 0, 2}
	queries := [][]int{
		{0, 2, 1}, {0, 2, 1}, {1, 1, 3},
	}
	expect := 2
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 3, 2, 1}
	queries := [][]int{
		{1, 3, 2},
		{0, 2, 1},
	}
	expect := -1
	runSample(t, a, queries, expect)
}
