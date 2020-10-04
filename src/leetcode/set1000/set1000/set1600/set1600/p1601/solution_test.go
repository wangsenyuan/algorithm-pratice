package p1601

import "testing"

func runSample(t *testing.T, n int, requests [][]int, expect int) {
	res := maximumRequests(n, requests)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	reqs := [][]int{
		{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4},
	}
	expect := 5
	runSample(t, n, reqs, expect)
}
