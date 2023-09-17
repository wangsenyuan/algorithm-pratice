package p2862

import "testing"

func runSample(t *testing.T, n int, k int, budget int, comps [][]int, stock []int, cost []int, expect int) {
	res := maxNumberOfAlloys(n, k, budget, comps, stock, cost)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	k := 2
	budget := 15
	composition := [][]int{{1, 1, 1}, {1, 1, 10}}
	stock := []int{0, 0, 0}
	cost := []int{1, 2, 3}
	expect := 2
	runSample(t, n, k, budget, composition, stock, cost, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	k := 2
	budget := 15
	composition := [][]int{{1, 1, 1}, {1, 1, 10}}
	stock := []int{0, 0, 100}
	cost := []int{1, 2, 3}
	expect := 5
	runSample(t, n, k, budget, composition, stock, cost, expect)
}
