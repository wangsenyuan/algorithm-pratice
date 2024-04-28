package p3130

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := medianOfUniquenessArray(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 4, 3, 4, 5}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 3, 5, 4}
	expect := 2
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1}
	expect := 1
	runSample(t, a, expect)
}
