package p1376

import "testing"

func runSample(t *testing.T, n int, headID int, manager []int, informTime []int, expect int) {
	res := numOfMinutes(n, headID, manager, informTime)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	headID := 0
	manager := []int{-1}
	informTime := []int{0}
	expect := 0
	runSample(t, n, headID, manager, informTime, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	headID := 2
	manager := []int{2, 2, -1, 2, 2, 2}
	informTime := []int{0, 0, 1, 0, 0, 0}
	expect := 1
	runSample(t, n, headID, manager, informTime, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	headID := 6
	manager := []int{1, 2, 3, 4, 5, 6, -1}
	informTime := []int{0, 6, 5, 4, 3, 2, 1}
	expect := 21
	runSample(t, n, headID, manager, informTime, expect)
}

func TestSample4(t *testing.T) {
	n := 15
	headID := 0
	manager := []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}
	informTime := []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	expect := 3
	runSample(t, n, headID, manager, informTime, expect)
}
