package p3132

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := minimumAddedInteger(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 20, 16, 12, 8}
	b := []int{14, 18, 10}
	expect := -2
	runSample(t, a, b, expect)
}
