package p1000

import "testing"

func runSample(t *testing.T, stones []int, K int, expect int) {
	res := mergeStones(stones, K)

	if res != expect {
		t.Errorf("Sample %v, %d, expect %d, but got %d", stones, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	stones := []int{3, 2, 4, 1}
	K := 2
	expect := 20
	runSample(t, stones, K, expect)
}

func TestSample2(t *testing.T) {
	stones := []int{3, 2, 4, 1}
	K := 3
	expect := -1
	runSample(t, stones, K, expect)
}

func TestSample3(t *testing.T) {
	stones := []int{3, 5, 1, 2, 6}
	K := 3
	expect := 25
	runSample(t, stones, K, expect)
}
