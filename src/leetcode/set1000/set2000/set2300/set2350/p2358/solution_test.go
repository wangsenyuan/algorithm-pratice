package p2358

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := maximumGroups(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 6, 12, 7, 3, 5}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{8, 8}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 2, 5, 39, 4, 28, 9, 41, 1, 4, 23, 45, 46, 37, 4, 17, 28, 6, 29, 39}
	expect := 5
	runSample(t, A, expect)
}
