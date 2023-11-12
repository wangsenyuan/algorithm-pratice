package p2932

import "testing"

func runSample(t *testing.T, a, b []int, expect int) {
	res := minOperations(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 7}
	b := []int{4, 5, 3}
	expect := 1
	runSample(t, a, b, expect)
}
