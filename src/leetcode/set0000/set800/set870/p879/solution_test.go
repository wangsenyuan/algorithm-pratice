package p879

import "testing"

func runSample(t *testing.T, G int, P int, group []int, profit []int, expect int) {
	res := profitableSchemes(G, P, group, profit)
	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", G, P, group, profit, expect, res)
	}
}

func TestSample1(t *testing.T) {
	G, P := 2, 2
	gang := []int{2, 2}
	profit := []int{2, 3}
	runSample(t, G, P, gang, profit, 2)
}

func TestSample2(t *testing.T) {
	G, P := 10, 5
	gang := []int{2, 3, 5}
	profit := []int{6, 7, 8}
	runSample(t, G, P, gang, profit, 7)
}
