package p3372

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := getLargestOutlier(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3, 5, 10}
	expect := 10
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-2, -1, -3, -6, 4}
	expect := 4
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 5, 5}
	expect := 5
	runSample(t, a, expect)
}
