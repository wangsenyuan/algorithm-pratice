package p3279

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := int(countSubarrays(a, k))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1}
	k := 1
	expect := 6
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 2}
	k := 1
	expect := 3
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3}
	k := 2
	expect := 2
	runSample(t, a, k, expect)
}
