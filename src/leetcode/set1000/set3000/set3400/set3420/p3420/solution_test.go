package p3420

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := countNonDecreasingSubarrays(a, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{6, 3, 1, 2, 4, 4}
	k := 7
	expect := 17
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 3, 1, 3, 6}
	k := 4
	expect := 12
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{9, 22, 9, 20, 23, 8}
	k := 15
	expect := 19
	runSample(t, a, k, expect)
}
