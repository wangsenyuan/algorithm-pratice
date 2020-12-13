package p5627

import "testing"

func runSample(t *testing.T, stones []int, expect int) {
	res := stoneGameVII(stones)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", stones, expect, res)
	}
}

func TestSample1(t *testing.T) {
	stones := []int{5, 3, 1, 4, 2}
	expect := 6
	runSample(t, stones, expect)
}

func TestSample2(t *testing.T) {
	stones := []int{7, 90, 5, 1, 100, 10, 10, 2}
	expect := 122
	runSample(t, stones, expect)
}
