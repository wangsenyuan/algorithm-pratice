package p3397

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := maxDistinctElements(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{9, 5, 5}
	k := 0
	expect := 2
	runSample(t, a, k, expect)
}
