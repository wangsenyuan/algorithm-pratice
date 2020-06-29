package p1494

import "testing"

func runSample(t *testing.T, n int, deps [][]int, k int, expect int) {
	res := minNumberOfSemesters(n, deps, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d",  expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	deps := [][]int{
		{2,1}, {3,1}, {1,4},
	}
	k := 2
	expect := 3
	runSample(t, n, deps, k, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	deps := [][]int{
		{2,1}, {3,1}, {4,1}, {1,5},
	}
	k := 2
	expect := 4
	runSample(t, n, deps, k, expect)
}

func TestSample3(t *testing.T) {
	n := 11
	deps := [][]int{}
	k := 2
	expect := 6
	runSample(t, n, deps, k, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	deps := [][]int{{2, 1}, {2, 4}}
	k := 2
	expect := 2
	runSample(t, n, deps, k, expect)
}

func TestSample5(t *testing.T) {
	n := 9
	deps := [][]int{{4,8}, {3,6}, {6,8}, {7,6}, {4,2}, {4,1}, {4,7}, {3,7}, {5,2}, {5,9}, {3,4}, {6,9}, {5,7}}
	k := 2
	expect := 5
	runSample(t, n, deps, k, expect)
}