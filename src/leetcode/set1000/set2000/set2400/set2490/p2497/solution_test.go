package p2497

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := int(minimumTotalCost(A, B))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	B := []int{1, 2, 3, 4, 5}
	expect := 10
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 2, 1, 3}
	B := []int{1, 2, 2, 3, 3}
	expect := 10
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 2}
	B := []int{1, 2, 2}
	expect := -1
	runSample(t, A, B, expect)
}
