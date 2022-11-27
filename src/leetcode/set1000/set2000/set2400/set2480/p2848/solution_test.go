package p2848

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := countSubarrays(A, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 1, 4, 5}
	k := 4
	expect := 3
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 3, 1}
	k := 3
	expect := 1
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 5, 1, 4, 3, 6}
	k := 1
	expect := 3
	runSample(t, A, k, expect)
}
