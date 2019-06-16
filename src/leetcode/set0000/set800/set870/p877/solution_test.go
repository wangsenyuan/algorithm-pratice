package p877

import "testing"

func runSample(t *testing.T, piles []int, expect bool) {
	res := stoneGame(piles)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", piles, expect, res)
	}
}

func TestSample1(t *testing.T) {
	piles := []int{5, 3, 4, 5}
	expect := true

	runSample(t, piles, expect)
}

func TestSample2(t *testing.T) {
	piles := []int{3, 1, 6, 5}
	expect := true

	runSample(t, piles, expect)
}
