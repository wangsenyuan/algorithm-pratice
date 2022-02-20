package p2179

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := int(goodTriplets(A, B))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 0, 1, 3}
	B := []int{0, 1, 2, 3}
	expect := 1
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 0, 1, 3, 2}
	B := []int{4, 1, 0, 2, 3}
	expect := 4
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{13, 14, 10, 2, 12, 3, 9, 11, 15, 8, 4, 7, 0, 6, 5, 1}
	B := []int{8, 7, 9, 5, 6, 14, 15, 10, 2, 11, 4, 13, 3, 12, 1, 0}
	expect := 77
	runSample(t, A, B, expect)
}
