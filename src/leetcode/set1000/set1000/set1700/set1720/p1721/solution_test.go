package p1721

import "testing"

func runSample(t *testing.T, source []int, target []int, swaps [][]int, expect int) {
	res := minimumHammingDistance(source, target, swaps)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{2, 1, 4, 5}
	swaps := [][]int{{0, 1}, {2, 3}}
	expect := 1
	runSample(t, a, b, swaps, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 3, 2, 4}
	swaps := [][]int{}
	expect := 2
	runSample(t, a, b, swaps, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 1, 2, 4, 3}
	b := []int{1, 5, 4, 2, 3}
	swaps := [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}}
	expect := 0
	runSample(t, a, b, swaps, expect)
}
