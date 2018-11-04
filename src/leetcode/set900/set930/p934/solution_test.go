package p934

import "testing"

func runSample(t *testing.T, A [][]int, expect int) {
	res := shortestBridge(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{0, 1}, {1, 0},
	}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{0, 1, 0}, {0, 0, 0}, {0, 0, 1},
	}
	expect := 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := [][]int{
		{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1},
	}
	expect := 1
	runSample(t, A, expect)
}
