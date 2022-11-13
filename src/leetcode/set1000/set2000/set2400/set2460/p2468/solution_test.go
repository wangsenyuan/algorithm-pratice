package p2468

import "testing"

func runSample(t *testing.T, edges [][]int, bob int, amount []int, expect int) {
	res := mostProfitablePath(edges, bob, amount)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{0, 1}, {1, 2}, {1, 3}, {3, 4},
	}
	bob := 3
	amount := []int{-2, 4, 2, -4, 6}
	expect := 6
	runSample(t, edges, bob, amount, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{0, 1},
	}
	bob := 1
	amount := []int{-7280, 2350}
	expect := -7280
	runSample(t, edges, bob, amount, expect)
}

func TestSample3(t *testing.T) {
	edges := [][]int{
		{0, 2}, {1, 8}, {1, 6}, {2, 3}, {2, 4}, {3, 7}, {4, 5}, {4, 9}, {6, 7},
	}
	bob := 3
	amount := []int{-378, -3744, -638, 9938, 3818, -336, 2722, 154, -1750, -1672}
	expect := 2785
	runSample(t, edges, bob, amount, expect)
}
