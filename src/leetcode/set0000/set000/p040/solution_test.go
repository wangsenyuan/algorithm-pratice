package p040

import "testing"

func runSample(t *testing.T, candidates []int, k int, expect int) {
	res := combinationSum2(candidates, k)

	if len(res) != expect {
		t.Errorf("Sample %v, expect %d, but got %d", candidates, expect, res)
	}
}

func TestSample1(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	k := 8
	expect := 4
	runSample(t, candidates, k, expect)
}
