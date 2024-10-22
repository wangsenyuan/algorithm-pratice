package p1749

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := maxAbsoluteSum(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

}

func TestSample1(t *testing.T) {
	a := []int{1, -3, 2, 3, -4}
	expect := 5
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, -5, 1, -4, 3, -2}
	expect := 8
	runSample(t, a, expect)
}
