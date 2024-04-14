package p3115

import "testing"

func runSample(t *testing.T, coins []int, k int, expect int) {
	res := findKthSmallest(coins, k)

	if expect != int(res) {
		t.Errorf("Sample %v %d, expect %d, but got %d", coins, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	coins := []int{3, 6, 9}
	k := 3
	expect := 9
	runSample(t, coins, k, expect)
}

func TestSample2(t *testing.T) {
	coins := []int{5, 2}
	k := 7
	expect := 12
	runSample(t, coins, k, expect)
}
