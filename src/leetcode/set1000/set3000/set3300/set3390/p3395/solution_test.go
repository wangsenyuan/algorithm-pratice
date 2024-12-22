package p3395

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := subsequencesWithMiddleMode(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1}
	expect := 6
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 2, 3, 3, 4}
	expect := 4
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 1, 0, 1, -1}
	expect := 0
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, -1, 1, 0, 0}
	expect := 0
	runSample(t, a, expect)
}
