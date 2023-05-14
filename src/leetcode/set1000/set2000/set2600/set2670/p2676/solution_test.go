package p2676

import "testing"

func runSample(t *testing.T, derived []int, expect bool) {
	res := doesValidArrayExist(derived)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	derived := []int{1, 1, 0}
	expect := true
	runSample(t, derived, expect)
}

func TestSample2(t *testing.T) {
	derived := []int{1, 1}
	expect := true
	runSample(t, derived, expect)
}

func TestSample3(t *testing.T) {
	derived := []int{1, 0}
	expect := false
	runSample(t, derived, expect)
}
