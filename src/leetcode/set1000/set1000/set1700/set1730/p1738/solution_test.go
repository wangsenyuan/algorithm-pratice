package p1738

import "testing"

func runSample(t *testing.T, matrix [][]int, k int, expect int) {
	res := kthLargestValue(matrix, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	matrix := [][]int{{5, 2}, {1, 6}}
	k := 1
	expect := 7
	runSample(t, matrix, k, expect)
}

func TestSample2(t *testing.T) {
	matrix := [][]int{{5, 2}, {1, 6}}
	k := 2
	expect := 5
	runSample(t, matrix, k, expect)
}
