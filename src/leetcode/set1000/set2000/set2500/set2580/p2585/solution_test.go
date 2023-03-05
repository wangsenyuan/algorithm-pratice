package p2585

import "testing"

func runSample(t *testing.T, target int, types [][]int, expect int) {
	res := waysToReachTarget(target, types)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	target := 6
	types := [][]int{
		{6, 1}, {3, 2}, {2, 3},
	}
	expect := 7
	runSample(t, target, types, expect)
}
func TestSample2(t *testing.T) {
	target := 5
	types := [][]int{
		{50, 1}, {50, 2}, {50, 5},
	}
	expect := 4
	runSample(t, target, types, expect)
}

func TestSample3(t *testing.T) {
	target := 18
	types := [][]int{
		{6, 1}, {3, 2}, {2, 3},
	}
	expect := 1
	runSample(t, target, types, expect)
}
