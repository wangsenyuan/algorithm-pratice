package lcp62

import "testing"

func runSample(t *testing.T, path [][]int, expect int) {
	res := transportationHub(path)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	path := [][]int{
		{0, 1}, {0, 3}, {1, 3}, {2, 0}, {2, 3},
	}
	expect := 3
	runSample(t, path, expect)
}
