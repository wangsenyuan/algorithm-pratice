package p3287

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := maxValue(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 6, 7}
	k := 1
	expect := 5
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 2, 5, 6, 7}
	k := 2
	expect := 2
	runSample(t, a, k, expect)
}
