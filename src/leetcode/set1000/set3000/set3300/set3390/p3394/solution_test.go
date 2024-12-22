package p3394

import "testing"

func runSample(t *testing.T, n int, rects [][]int, expect bool) {
	res := checkValidCuts(n, rects)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	rects := [][]int{
		{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5},
	}
	expect := true
	runSample(t, n, rects, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	rects := [][]int{
		{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4},
	}
	expect := false
	runSample(t, n, rects, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	rects := [][]int{
		{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3},
	}
	expect := true
	runSample(t, n, rects, expect)
}
