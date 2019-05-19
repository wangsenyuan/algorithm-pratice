package p1049

import "testing"

func runSample(t *testing.T, stones []int, expect int) {
	res := lastStoneWeightII(stones)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", stones, expect, res)
	}
}

func TestSample1(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	runSample(t, stones, 1)
}

func TestSample2(t *testing.T) {
	stones := []int{76, 1, 28, 17, 20, 48, 51, 28}
	runSample(t, stones, 3)
}
