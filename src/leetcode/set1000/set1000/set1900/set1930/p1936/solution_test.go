package p1936

import "testing"

func runSample(t *testing.T, rungs []int, dist int, expect int) {
	res := addRungs(rungs, dist)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	rungs := []int{1, 3, 5, 10}
	dist := 2
	expect := 2
	runSample(t, rungs, dist, expect)
}

func TestSample2(t *testing.T) {
	rungs := []int{3, 6, 8, 10}
	dist := 3
	expect := 0
	runSample(t, rungs, dist, expect)
}

func TestSample3(t *testing.T) {
	rungs := []int{3, 4, 6, 7}
	dist := 2
	expect := 1
	runSample(t, rungs, dist, expect)
}
