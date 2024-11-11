package p3351

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := sumOfGoodSubsequences(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 1}
	expect := 14
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 4, 5}
	expect := 40
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 1}
	expect := 2
	runSample(t, a, expect)
}
