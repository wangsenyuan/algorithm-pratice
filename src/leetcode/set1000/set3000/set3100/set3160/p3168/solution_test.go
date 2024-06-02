package p3168

import "testing"

func runSample(t *testing.T, k int, a []int, expect int) {
	res := minimumDifference(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	a := []int{1, 2, 4, 5}
	expect := 1
	runSample(t, k, a, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	a := []int{1, 2, 1, 2}
	expect := 0
	runSample(t, k, a, expect)
}
