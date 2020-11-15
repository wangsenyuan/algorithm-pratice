package p5552

import "testing"

func runSample(t *testing.T, forbidden []int, a int, b int, x int, expect int) {
	res := minimumJumps(forbidden, a, b, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{14, 4, 18, 1, 15}
	a, b := 3, 15
	x := 9
	expect := 3
	runSample(t, arr, a, b, x, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{8, 3, 16, 6, 12, 20}
	a, b := 15, 13
	x := 11
	expect := -1
	runSample(t, arr, a, b, x, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 6, 2, 14, 5, 17, 4}
	a, b := 16, 9
	x := 7
	expect := 2
	runSample(t, arr, a, b, x, expect)
}
