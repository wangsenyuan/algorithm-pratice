package e

import "testing"

func runSample(t *testing.T, edges [][]int, startA int, startB int, expect int) {
	res := chaseGame(edges, startA, startB)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 1}, {2, 5}, {5, 6},
	}
	startA := 3
	startB := 5
	expect := 3
	runSample(t, edges, startA, startB, expect)
}

func TestSample3(t *testing.T) {
	edges := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 1}, {2, 5}, {5, 6},
	}
	startA := 4
	startB := 5
	expect := 4
	runSample(t, edges, startA, startB, expect)
}

func TestSample4(t *testing.T) {
	edges := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 1}, {2, 5}, {5, 6}, {6, 7},
	}
	startA := 3
	startB := 6
	expect := 4
	runSample(t, edges, startA, startB, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 1},
	}
	startA := 1
	startB := 3
	expect := -1
	runSample(t, edges, startA, startB, expect)
}
